package markdown

import (
	"hackbar-report/internal/interface-adapter/markdown/components"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

const NO_INPUT = "入力無し"

func Marshal[From comparable](v From) components.MarkdownBlock {
	lists := iterateFields(reflect.ValueOf(v),
		func(fieldName string, tag reflect.StructTag, rv reflect.Value) []components.MarkdownBlock {
			rt := rv.Type()
			fieldCount := rt.NumField()
			blocks := make([]components.MarkdownBlock, 0, fieldCount+2)
			blocks = append(blocks, components.Heading(2, label(tag, fieldName)))

			body := iterateFields(rv,
				func(fieldName string, tag reflect.StructTag, value reflect.Value) components.MarkdownBlock {
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
						if !isYes(value.String()) {
							options = append(options, components.WithChild(value.String()))
						}
						return components.List(label, options...)
					}

					if isFormatten, format, options := hasFormat(tag); isFormatten {
						v := components.TextValue{
							Value: value.String(),
							Total: 0,
						}
						if strings.Contains(format, "${total}") {
							skipFieldName := fieldName
							vv := iterateFields(rv,
								func(fieldName string, tag reflect.StructTag, value reflect.Value) int {
									if fieldName == skipFieldName {
										return 0
									}
									if value.Kind() != reflect.String {
										return 0
									}
									valueInt, err := strconv.Atoi(value.String())
									if err != nil {
										return 0
									}
									rate := totalAddUpRate(tag)
									return valueInt * rate
								},
							)
							v.Total = reduce(add, vv)
						}
						return components.Text(label, v, options...)
					}

					return components.Text(label, components.TextValue{Value: value.String()})
				},
			)

			return append(append(blocks, body...), "  ")
		},
	)

	return Join(Filter(Flatten(lists), notEmpty)...)
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

func isYes(value string) bool {
	return slices.Contains(
		[]string{"y", "yes", "on", "true"},
		value,
	)
}

func isNone(value string) bool {
	return slices.Contains(
		[]string{"", "-", "n", "no", "none", "off", "false"},
		value,
	)
}

func Flatten(lists [][]components.MarkdownBlock) []components.MarkdownBlock {
	var res []components.MarkdownBlock
	for _, blocks := range lists {
		res = append(res, "")
		res = append(res, blocks...)
	}
	return res[1:]
}

func Filter(blocks []components.MarkdownBlock, condition func(components.MarkdownBlock) bool) []components.MarkdownBlock {
	new := make([]components.MarkdownBlock, 0, len(blocks))
	for _, block := range blocks {
		if condition(block) {
			new = append(new, block)
		}
	}
	return new
}

func notEmpty(block components.MarkdownBlock) bool {
	return block != ""
}

func Join(blocks ...components.MarkdownBlock) components.MarkdownBlock {
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
	return components.MarkdownBlock(b.String())
}

func add(x, y int) int {
	return x + y
}

func reduce(f func(int, int) int, ary []int) int {
	v := 0
	for _, x := range ary {
		v = f(v, x)
	}
	return v
}
