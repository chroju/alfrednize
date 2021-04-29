package main

import (
	"os"

	ui "github.com/chroju/alfrednize/UI"
	"github.com/chroju/alfrednize/command"
)

func main() {
	args := os.Args[1:]
	ui := ui.NewCmdUI()
	os.Exit(command.ExecCommand(ui, args))
}
