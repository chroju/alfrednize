package ui

import (
	"bufio"
	"io"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// UI is a command line UI interface
type UI interface {
	In() []string
	Out() io.Writer
	ErrOut() io.Writer
}

type cmdUI struct{}

// NewCmdUI generate new UI for command line
func NewCmdUI() UI {
	ui := &cmdUI{}
	return ui
}

func (ui *cmdUI) In() []string {
	if terminal.IsTerminal(int(syscall.Stdin)) {
		return nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	var items []string

	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	return items
}

func (ui *cmdUI) Out() io.Writer {
	return os.Stdout
}

func (ui *cmdUI) ErrOut() io.Writer {
	return os.Stderr
}
