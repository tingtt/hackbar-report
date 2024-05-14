package clipboard

import "golang.design/x/clipboard"

func Write(buf []byte) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}

	clipboard.Write(clipboard.FmtText, buf)
	return nil
}
