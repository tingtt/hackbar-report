package components

import (
	"strconv"
	"strings"
)

type textOption struct {
	format *string
}
type TextOptionApplier func(*textOption)

type TextValue struct {
	Value string
	Total int
}

func Text(label string, value TextValue, options ...TextOptionApplier) MarkdownBlock {
	option := &textOption{}
	for _, apply := range options {
		apply(option)
	}

	if option.format == nil {
		format := "${label} ${value}"
		option.format = &format
	}

	res := strings.ReplaceAll(*option.format, "${label}", label)
	res = strings.ReplaceAll(res, "${value}", value.Value)
	res = strings.ReplaceAll(res, "${total}", strconv.Itoa(value.Total))
	return MarkdownBlock(res)
}

func WithFormat(format string) TextOptionApplier {
	return func(lo *textOption) {
		lo.format = &format
	}
}
