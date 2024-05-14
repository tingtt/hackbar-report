package report

import (
	"fmt"
	"hackbar-report/internal/interface-adapter/markdown"
	"hackbar-report/internal/interface-adapter/markdown/components"
	promptgroup "hackbar-report/internal/usecase/prompt-group"
	"io"

	"github.com/spf13/cobra"
)

func command[T comparable](out io.Writer, in io.Reader, prompt T) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		err := promptgroup.Run(out, in, &prompt)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(out, components.Separator(32))
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(out, markdown.Marshal(prompt))
		if err != nil {
			return err
		}

		return nil
	}
}
