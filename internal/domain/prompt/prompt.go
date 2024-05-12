package prompt

import (
	"bufio"
	"fmt"
	"io"
)

type Prompt interface {
	Run(io.Writer, io.Reader) (answer string, err error)
}

func New(message string) Prompt {
	return &prompt{message}
}

type prompt struct {
	message string
}

func (p *prompt) Run(out io.Writer, in io.Reader) (answer string, err error) {
	err = p.printMessage(out)
	if err != nil {
		return "", err
	}

	return p.scan(in)
}

func (p *prompt) printMessage(out io.Writer) error {
	_, err := fmt.Fprint(out, p.message)
	return err
}

func (p *prompt) scan(in io.Reader) (input string, err error) {
	scanner := bufio.NewScanner(in)
	if !scanner.Scan() {
		return
	}
	return scanner.Text(), nil
}
