package components

import (
	"bytes"
	"fmt"
	"strings"
)

type listOption struct {
	indentSize   int
	separateWith []string
	children     *string
}
type ListOptionApplier func(*listOption)

func List(value string, options ...ListOptionApplier) (res MarkdownBlock) {
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
			res = ApplyChild(res, *option.children, option.separateWith)
		}()
	}

	if len(option.separateWith) != 0 {
		separated := SplitAny(value, strings.Join(option.separateWith, ""))
		separated[0] = string(List(separated[0]))
		return MarkdownBlock(
			strings.Join(separated, "\n- "),
		)
	}

	return MarkdownBlock(
		fmt.Sprintf("- %s", value),
	)
}

func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func ApplyChild(res MarkdownBlock, children string, separators []string) MarkdownBlock {
	options := make([]ListOptionApplier, 0, 2)
	if len(separators) != 0 {
		options = append(options, WithSeparateBy(separators))
	}
	options = append(options, WithIndent(2))

	elems := []string{string(res)}
	elems = append(elems, string(List(children, options...)))
	return MarkdownBlock(strings.Join(elems, "\n"))
}

func WithSeparateBy(separators []string) ListOptionApplier {
	return func(lo *listOption) {
		lo.separateWith = separators
	}
}

func WithIndent(size int) ListOptionApplier {
	return func(lo *listOption) {
		lo.indentSize = size
	}
}

func WithChild(children string) ListOptionApplier {
	return func(lo *listOption) {
		lo.children = &children
	}
}
