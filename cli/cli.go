package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"vdns/cli/command"
	"vdns/lib/api"
	"vdns/lib/vlog"
)

//goland:noinspection SpellCheckingInspection
const (
	CliVersion = api.Version
	Author     = "zf1976"
	CliName    = "vdns"
	Usage      = "vdns is a tool that supports multi-DNS service provider resolution operations."
	Email      = "verticle@foxmail.com"
)

var app = cli.NewApp()

//goland:noinspection SpellCheckingInspection
func main() {
	app.Commands = []*cli.Command{
		command.SupportCommand(),
		command.DNSConfigCommand(),
	}
	err := app.Run(os.Args)
	if err != nil {
		vlog.Fatalf("running err: %v", err)
		return
	}
}

func init() {
	app.Name = CliName
	app.HelpName = CliName
	app.Usage = Usage
	//app.Authors = []*cli.Author{{
	//	Name:  Author,
	//	Email: Email,
	//}}
	//app.Version = CliVersion
}
