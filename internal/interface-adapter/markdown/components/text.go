package components

import "strings"

type textOption struct {
	format *string
}
type TextOptionApplier func(*textOption)

func Text(label string, value string, options ...TextOptionApplier) MarkdownBlock {
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

func WithFormat(format string) TextOptionApplier {
	return func(lo *textOption) {
		lo.format = &format
	}
}
