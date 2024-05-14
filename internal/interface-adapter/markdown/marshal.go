package markdown

import (
	"bytes"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

type MarkdownBlock string

const NO_INPUT = "入力無し"

func Marshal[From comparable](v From) MarkdownBlock {
	mdBlocks := iterateFields(reflect.ValueOf(v),
		func(fieldName string, tag reflect.StructTag, rv reflect.Value) []MarkdownBlock {
			rt := rv.Type()
			fieldCount := rt.NumField()
			mdBlocks := make([]MarkdownBlock, 0, fieldCount+1)
			mdBlocks = append(mdBlocks, heading(2, label(tag, fieldName)))

			body := iterateFields(rv,
				func(fieldName string, tag reflect.StructTag, value reflect.Value) MarkdownBlock {
					if value.Kind() != reflect.String {
						return ""
					}

					if isNone(value.String()) {
						if isSkippable(tag) {
							return ""
						}
						v, _ := lookupDefault(tag, NO_INPUT)
						value = reflect.ValueOf(v)
					}

					label := label(tag, fieldName /* default */)

					if isList, options := isList(tag); isList {
						options = append(options, withChild(value.String()))
						return list(label, options...)
					}

					if isFormatten, options := isFormatten(tag); isFormatten {
						return text(label, value.String(), options...)
					}

					return text(label, value.String())
				},
			)
			return append(mdBlocks, body...)
		},
	)
	// TODO: filter (remove empty string)
	return join(flatten(mdBlocks))
}

func iterateFields[T any](rv reflect.Value, yield func(fieldName string, tag reflect.StructTag, rv reflect.Value) T) []T {
	rt := rv.Type()
	fieldCount := rt.NumField()

	res := make([]T, 0, fieldCount)

	for i := range make([]interface{}, fieldCount) {
		field := rt.Field(i)
		res = append(res, yield(field.Name, field.Tag, rv.FieldByName(field.Name)))
	}

	return res
}

func lookup(tag reflect.StructTag, key string, defaultValue string) string {
	label, ok := tag.Lookup(key)
	if !ok {
		label = defaultValue
	}
	return label
}

func label(tag reflect.StructTag, defaultValue string) string {
	return lookup(tag, "label", defaultValue /* default */)
}

func isSkippable(tag reflect.StructTag) bool {
	return strings.HasSuffix(lookup(tag, "mdblk-type", ""), ",omitempty")
}

func isList(tag reflect.StructTag) (bool, []listOptionApplier) {
	if !strings.HasPrefix(lookup(tag, "mdblk-type", ""), "list") {
		return false, nil
	}

	optionAppliers := make([]listOptionApplier, 0, 1)

	if separator := lookup(tag, "mdblk-list-separate-with", ""); separator != "" {
		optionAppliers = append(optionAppliers, withSeparateBy(separator))
	}

	return true, optionAppliers
}

func isFormatten(tag reflect.StructTag) (bool, []textOptionApplier) {
	if lookup(tag, "mdblk-format", "") == "" {
		return false, nil
	}
	format := lookup(tag, "mdblk-format", "")
	return true, []textOptionApplier{withFormat(format)}
}

func isNone(value string) bool {
	return slices.Contains(
		[]string{"", "-", "no", "none", "off", "false"},
		value,
	)
}

func lookupDefault(tag reflect.StructTag, defaultValue string) (string, bool) {
	_defaultValue := lookup(tag, "mdblk-default", defaultValue)
	return _defaultValue, defaultValue != _defaultValue
}

func heading(size int, value string) MarkdownBlock {
	prefix := string(bytes.Repeat([]byte("#"), size))
	return MarkdownBlock(
		fmt.Sprintf("%s %s\n", prefix, value),
	)
}

type textOption struct {
	format *string
}
type textOptionApplier func(*textOption)

func text(label string, value string, options ...textOptionApplier) MarkdownBlock {
	option := &textOption{}
	for _, apply := range options {
		apply(option)
	}

	if option.format == nil {
		format := "${label} ${value}"
		option.format = &format
	}

	res := strings.ReplaceAll(*option.format, "${label}", label)
	res = strings.ReplaceAll(res, "${value}", value)
	return MarkdownBlock(res)

}

type listOption struct {
	indentSize   int
	separateWith *string
	children     *string
}
type listOptionApplier func(*listOption)

func withFormat(format string) textOptionApplier {
	return func(lo *textOption) {
		lo.format = &format
	}
}

func list(value string, options ...listOptionApplier) (res MarkdownBlock) {
	option := &listOption{}
	for _, apply := range options {
		apply(option)
	}

	if option.indentSize != 0 {
		defer func() {
			indent := string(bytes.Repeat([]byte(" "), option.indentSize))
			res = MarkdownBlock(
				indent + strings.ReplaceAll(string(res), "\n", fmt.Sprintf("\n%s", indent)),
			)
		}()
	}
	if option.children != nil {
		defer func() {
			res = applyChild(res, *option.children, option.separateWith)
		}()
	}

	if option.separateWith != nil {
		separated := strings.Split(value, *option.separateWith)
		separated[0] = string(list(separated[0]))
		return MarkdownBlock(
			strings.Join(separated, "\n- "),
		)
	}

	return MarkdownBlock(
		fmt.Sprintf("- %s", value),
	)
}

func applyChild(res MarkdownBlock, children string, separater *string) MarkdownBlock {
	options := make([]listOptionApplier, 0, 2)
	if separater != nil {
		options = append(options, withSeparateBy(*separater))
	}
	options = append(options, withIndent(2))

	elems := []string{string(res)}
	elems = append(elems, string(list(children, options...)))
	return MarkdownBlock(strings.Join(elems, "\n"))
}

func withSeparateBy(separater string) listOptionApplier {
	return func(lo *listOption) {
		lo.separateWith = &separater
	}
}

func withIndent(size int) listOptionApplier {
	return func(lo *listOption) {
		lo.indentSize = size
	}
}

func withChild(children string) listOptionApplier {
	return func(lo *listOption) {
		lo.children = &children
	}
}

func flatten(lists [][]MarkdownBlock) []MarkdownBlock {
	var res []MarkdownBlock
	for _, list := range lists {
		res = append(res, "")
		res = append(res, list...)
	}
	return res[1:]
}

func join(blocks []MarkdownBlock) MarkdownBlock {
	separator := "\n"

	var n int
	if len(separator) > 0 {
		n += len(separator) * (len(blocks) - 1)
	}
	for _, elem := range blocks {
		n += len(elem)
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(string(blocks[0]))
	for _, s := range blocks[1:] {
		b.WriteString(separator)
		b.WriteString(string(s))
	}
	return MarkdownBlock(b.String())
}
