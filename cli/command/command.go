package command

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func SupportCommand() *cli.Command {
	return &cli.Command{
		Name:    "supports",
		Aliases: []string{"s"},
		Usage:   "Supported DNS service providers",
		Action: func(c *cli.Context) error {
			fmt.Println("Support: AliDNS, DNSPod, Cloudflare, HuaweiDNS")
			return nil
		},
	}
}

func DNSConfigCommand() *cli.Command {
	return &cli.Command{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "Configure DNS Service Provider Access Key Pair",
		Subcommands: []*cli.Command{
			alidnsConfigCommand(),
		},
	}
}

//goland:noinspection SpellCheckingInspection
func alidnsConfigCommand() *cli.Command {
	var ak string = ""
	var sk string = ""
	return &cli.Command{
		Name:  "alidns",
		Usage: "Configure AliDNS access key pair",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "ak",
				Usage:       "alidns accessKey",
				Destination: &ak,
			},
			&cli.StringFlag{
				Name:        "sk",
				Usage:       "alidns secretKey",
				Destination: &sk,
			},
		},
		Action: func(ctx *cli.Context) error {

			return nil
		},
	}
}
