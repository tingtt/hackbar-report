package close

import (
	"fmt"
	"hackbar-report/internal/usecase/close"
	"io"

	"github.com/spf13/cobra"
)

func NewCmd(out io.Writer, in io.Reader) *cobra.Command {
	return &cobra.Command{
		Use: "close",
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := close.Run(out, in)
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
