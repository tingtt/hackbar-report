package components

import "bytes"

func Separator(width int) MarkdownBlock {
	return MarkdownBlock(bytes.Repeat([]byte("-"), 32))
}
