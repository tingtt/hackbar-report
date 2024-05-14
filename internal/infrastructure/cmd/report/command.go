package report

import (
	"bytes"
	"fmt"
	"hackbar-report/internal/interface-adapter/markdown"
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

		fmt.Fprintln(out, string(bytes.Repeat([]byte("-"), 32)))

		_, err = fmt.Fprintln(out, markdown.Marshal(prompt))
		if err != nil {
			return err
		}

		return nil
	}
}
