package cmd

import (
	"hackbar-report/internal/infrastructure/cmd/report/cmd/close"
	"hackbar-report/internal/infrastructure/cmd/report/cmd/open"
	"io"

	"github.com/spf13/cobra"
)

func New(out io.Writer, in io.Reader) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "report",
		SilenceUsage: true,
	}

	cmd.AddCommand(open.NewCmd(out, in))
	cmd.AddCommand(close.NewCmd(out, in))

	return cmd
}
