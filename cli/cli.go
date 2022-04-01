package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"time"
	"vdns/cli/command"
	"vdns/lib/api"
	"vdns/lib/vlog"
)

var app = cli.NewApp()

//goland:noinspection SpellCheckingInspection
const (
	CliVersion = api.Version
	CliName    = "vdns"
	Usage      = "A tool that supports multi-DNS service provider resolution operations"
)

//goland:noinspection SpellCheckingInspection
func main() {
	initCLI()
	err := app.Run(os.Args)
	if err != nil {
		vlog.Fatalf("running err: %v", err)
		return
	}
}

//goland:noinspection SpellCheckingInspection
func initCLI() {

	app.Commands = []*cli.Command{
		command.ShowCommand(),
		command.ConfigCommand(),
		command.ResolveRecord(),
		{
			Name:  "start",
			Usage: "Start vdns service",
		},
		{
			Name:  "stop",
			Usage: "Stop vdns service",
		},
		{
			Name:  "restart",
			Usage: "Restart vdns service",
		},
		{
			Name:  "install",
			Usage: "Install vdns service",
		},
		{
			Name:  "uninstall",
			Usage: "Uninstall vdns service",
		},
	}
}

func init() {
	app.Name = CliName
	app.HelpName = CliName
	app.Usage = Usage
	app.Compiled = time.Now()
	app.Version = CliVersion
}
