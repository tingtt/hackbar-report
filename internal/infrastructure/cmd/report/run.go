package report

import (
	"io"
)

func Run(out io.Writer, in io.Reader) error {
	report := new(out, in)
	return report.Execute()
}
