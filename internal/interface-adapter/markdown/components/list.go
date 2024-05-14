package components

import (
	"bytes"
	"fmt"
	"strings"
)

type listOption struct {
	indentSize   int
	separateWith *string
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

	if option.separateWith != nil {
		separated := strings.Split(value, *option.separateWith)
		separated[0] = string(List(separated[0]))
		return MarkdownBlock(
			strings.Join(separated, "\n- "),
		)
	}

	return MarkdownBlock(
		fmt.Sprintf("- %s", value),
	)
}

func ApplyChild(res MarkdownBlock, children string, separater *string) MarkdownBlock {
	options := make([]ListOptionApplier, 0, 2)
	if separater != nil {
		options = append(options, WithSeparateBy(*separater))
	}
	options = append(options, WithIndent(2))

	elems := []string{string(res)}
	elems = append(elems, string(List(children, options...)))
	return MarkdownBlock(strings.Join(elems, "\n"))
}

func WithSeparateBy(separater string) ListOptionApplier {
	return func(lo *listOption) {
		lo.separateWith = &separater
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
