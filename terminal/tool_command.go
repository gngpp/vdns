package terminal

import "github.com/urfave/cli/v2"

func ToolCommand() *cli.Command {
	return &cli.Command{
		Name:  "tool",
		Usage: "Tool command",
	}
}
