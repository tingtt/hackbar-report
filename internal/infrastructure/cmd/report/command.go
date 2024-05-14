package report

import (
	"fmt"
	"hackbar-report/internal/infrastructure/clipboard"
	"hackbar-report/internal/interface-adapter/markdown"
	"hackbar-report/internal/interface-adapter/markdown/components"
	promptgroup "hackbar-report/internal/usecase/prompt-group"
	"io"

	"github.com/spf13/cobra"
)

func command[T comparable](out io.Writer, in io.Reader, prompt T, heading string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		err := promptgroup.Run(out, in, &prompt)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(out, components.Separator(32))
		if err != nil {
			return err
		}

		md := markdown.Marshal(prompt)

		_, err = fmt.Fprintln(out, md)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(out, components.Separator(32))
		if err != nil {
			return err
		}

		err = clipboard.Write([]byte(md))
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(out, "クリップボードにコピーされました。")
		if err != nil {
			return err
		}

		return nil
	}
}
