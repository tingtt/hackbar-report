package open

import (
	"fmt"
	"hackbar-report/internal/usecase/open"
	"io"

	"github.com/spf13/cobra"
)

func NewCmd(out io.Writer, in io.Reader) *cobra.Command {
	return &cobra.Command{
		Use: "open",
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := open.Run(out, in)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(out, *res)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
