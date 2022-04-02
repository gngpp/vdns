package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"time"
	"vdns/lib/api"
	"vdns/lib/vlog"
	"vdns/terminal"
)

var app = cli.NewApp()

//goland:noinspection SpellCheckingInspection
const (
	CliVersion = api.Version
	CliName    = "vdns"
	Usage      = "This is A tool that supports multi-DNS service provider resolution operations."
)

func main() {
	initCLI()
	err := app.Run(os.Args)
	if err != nil {
		vlog.Fatalf("running err: %v", err)
		return
	}
}

func initCLI() {
	app.Commands = []*cli.Command{
		terminal.ShowCommand(),
		terminal.ConfigCommand(),
		terminal.ResolveRecord(),
		terminal.ServerCommand(),
	}
}

func init() {
	app.Name = CliName
	app.HelpName = CliName
	app.Usage = Usage
	app.Compiled = time.Now()
}
