package report

import (
	"hackbar-report/internal/infrastructure/cmd/report/cmd"
	"io"
)

func Run(out io.Writer, in io.Reader) error {
	report := cmd.New(out, in)
	return report.Execute()
}
