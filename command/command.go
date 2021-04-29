package command

import (
	"fmt"

	ui "github.com/chroju/alfrednize/UI"
	"github.com/chroju/alfrednize/alfrednize"
)

const (
	version     = "0.1.0"
	helpMessage = `
alfrednize - So simple command to generate JSON for Alfred

Usage:
  Use with standard input.

Example:
  $ echo -e 'foo\nbar' | alfrednize | jq
  {
    "items": [
      {
        "uid": "foo",
        "title": "foo",
        "subtitle": "",
        "arg": "foo",
        "match": "foo",
        "autocomplete": "foo"
      },
      {
        "uid": "bar",
        "title": "bar",
        "subtitle": "",
        "arg": "bar",
        "match": "bar",
        "autocomplete": "bar"
      }
    ]
  }

`
)

// ExecCommand exec command
func ExecCommand(cmd ui.UI, args []string) int {
	if len(args) > 1 {
		fmt.Fprint(cmd.ErrOut(), helpCommand())
		return 1
	}

	if len(args) == 1 {
		switch args[0] {
		case "-h":
			fmt.Fprint(cmd.Out(), helpCommand())
			return 0
		case "--help":
			fmt.Fprint(cmd.Out(), helpCommand())
			return 0
		case "-v":
			fmt.Fprint(cmd.Out(), versionCommand())
			return 0
		case "--version":
			fmt.Fprint(cmd.Out(), versionCommand())
			return 0
		case "-":
			break
		default:
			fmt.Fprint(cmd.ErrOut(), helpCommand())
			return 1
		}
	}

	items := cmd.In()
	if len(items) == 0 {
		fmt.Fprint(cmd.ErrOut(), helpCommand())
		return 1
	}

	alfrednizeItems, err := alfrednize.Alfrednize(items)
	if err != nil {
		fmt.Fprint(cmd.ErrOut(), err)
		return 1
	}

	fmt.Fprintln(cmd.Out(), string(alfrednizeItems))

	return 0
}

func helpCommand() string {
	return helpMessage
}

func versionCommand() string {
	return "alfrednize v" + version
}
