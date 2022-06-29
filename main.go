package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"time"
	"vdns/lib/vlog"
	"vdns/terminal"
)

//goland:noinspection SpellCheckingInspection
const (
	CliName = "vdns"
	Usage   = "This is A tool that supports multi-DNS service provider resolution operations."
)

func main() {
	var app = cli.NewApp()
	app.Name = CliName
	app.HelpName = CliName
	app.Usage = Usage
	//app.Version = api.Version
	app.Compiled = time.Now()
	app.EnableBashCompletion = true

	// debug mode
	var debug bool
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"d"},
			Usage:       "Enable debug mode",
			Destination: &debug,
		},
	}

	app.Before = func(ctx *cli.Context) error {
		if debug {
			vlog.SetLevel(vlog.Level.DEBUG)
		}
		return nil
	}

	// provider config and ddns server cli
	app.Commands = []*cli.Command{
		terminal.ConfigCommand(),
		terminal.ServerCommand(),
	}
	// dns record resolve cli
	app.Commands = append(app.Commands, terminal.ResolveRecordCommand())
	// common cli
	app.Commands = append(app.Commands, terminal.Command()...)
	err := app.Run(os.Args)
	if err != nil {
		vlog.Fatalf("running fatal: %v", err)
		return
	}
}
