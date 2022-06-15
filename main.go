package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"time"
	"vdns/lib/api"
	"vdns/lib/vlog"
	"vdns/terminal"
)

//goland:noinspection SpellCheckingInspection
const (
	CliName = "vdns"
	Usage   = "This is A tool that supports multi-DNS service provider resolution operations."
)

func main() {
	app := initCLI()
	err := app.Run(os.Args)
	if err != nil {
		vlog.Fatalf("running fatal: %v", err)
		return
	}
}

func initCLI() *cli.App {
	var app = cli.NewApp()
	app.Name = CliName
	app.HelpName = CliName
	app.Usage = Usage
	app.Version = api.Version
	app.Compiled = time.Now()
	// provider config and ddns server cli
	app.Commands = []*cli.Command{
		terminal.ConfigCommand(),
		terminal.ServerCommand(),
	}
	// dns record resolve cli
	app.Commands = append(app.Commands, terminal.ResolveRecordCommand())
	// common cli
	app.Commands = append(app.Commands, terminal.Command()...)
	return app
}
