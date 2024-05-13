package report

import (
	"hackbar-report/internal/usecase/close"
	"hackbar-report/internal/usecase/open"
	"io"

	"github.com/spf13/cobra"
)

func new(out io.Writer, in io.Reader) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "report",
		SilenceUsage: true,
	}

	cmd.AddCommand(&cobra.Command{
		Use:  "open",
		RunE: command(out, in, open.Prompt{}),
	})
	cmd.AddCommand(&cobra.Command{
		Use:  "close",
		RunE: command(out, in, close.Prompt{}),
	})

	return cmd
}
