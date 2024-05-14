package components

import (
	"bytes"
	"fmt"
)

func Heading(size int, value string) MarkdownBlock {
	prefix := string(bytes.Repeat([]byte("#"), size))
	return MarkdownBlock(
		fmt.Sprintf("%s %s\n", prefix, value),
	)
}
