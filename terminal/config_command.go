package terminal

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/urfave/cli/v2"
	"vdns/config"
	"vdns/lib/util/convert"
	"vdns/lib/util/strs"
)

//goland:noinspection SpellCheckingInspection
func ConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Configure dns service provider access key pair",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Usage: "sava data table to csv filepath",
			},
		},
		Subcommands: []*cli.Command{
			configCommand(config.AlidnsProvider),
			configCommand(config.DnspodProvider),
			configCommand(config.HuaweiDnsProvider),
			configCommand(config.CloudflareProvider),
			{
				Name:  "print",
				Usage: "Print all dns configuration",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "path",
						Usage: "sava table to csv filepath",
					},
				},
				Action: printConfigAction(),
			},
			{
				Name:   "import",
				Usage:  "Import all dns configuration",
				Action: importConfigAction(),
			},
		},
	}
}

//goland:noinspection SpellCheckingInspection
func configCommand(commandName string) *cli.Command {
	var ak string
	var sk string
	var token string
	return &cli.Command{
		Name:                   commandName,
		Aliases:                []string{convert.AsStringValue(string(commandName[0]))},
		Usage:                  "Configure " + commandName + " access key pair",
		UseShortOptionHandling: true,
		Subcommands: []*cli.Command{
			{
				Name:    "print",
				Aliases: []string{"p"},
				Usage:   "Print dns provider configuration",
				Action: func(_ *cli.Context) error {
					readConfig, err := config.ReadConfig()
					if err != nil {
						return err
					}
					dnsConfig := readConfig.ConfigsMap.Get(commandName)
					table, err := gotable.CreateByStruct(dnsConfig)
					if err != nil {
						return err
					}
					_ = table.AddRow([]string{*dnsConfig.Provider, *dnsConfig.Ak, *dnsConfig.Sk, *dnsConfig.Token})
					fmt.Println(table)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "ak",
				Usage:       "api access key",
				Destination: &ak,
			},
			&cli.StringFlag{
				Name:        "sk",
				Usage:       "api secret key",
				Destination: &sk,
			},
			&cli.StringFlag{
				Name:        "token",
				Usage:       "api token",
				Destination: &token,
			},
		},
		Action: func(ctx *cli.Context) error {
			readConfig, err := config.ReadConfig()
			dnsConfig := readConfig.ConfigsMap.Get(commandName)
			if err != nil {
				return err
			}
			isModify := false
			if !strs.IsEmpty(ak) {
				dnsConfig.Ak = &ak
				isModify = true
			}
			if !strs.IsEmpty(sk) {
				dnsConfig.Sk = &sk
				isModify = true
			}
			if !strs.IsEmpty(token) {
				dnsConfig.Token = &token
				isModify = true
			}

			if isModify {
				err = config.WriteConfig(readConfig)
				if err != nil {
					return err
				}
				table, err := gotable.Create("provider", "ak", "sk", "token")
				if err != nil {
					return err
				}
				err = table.AddRow([]string{*dnsConfig.Provider, *dnsConfig.Ak, *dnsConfig.Sk, *dnsConfig.Token})
				if err != nil {
					return err
				}
				fmt.Println(table)
			}
			return nil
		},
	}
}

func printConfigAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		config, err := config.ReadConfig()
		if err != nil {
			return err
		}
		table, err := gotable.Create("provider", "ak", "sk", "token")
		for key := range config.ConfigsMap {
			dnsConfig := config.ConfigsMap.Get(key)
			if dnsConfig != nil {
				err := table.AddRow([]string{*dnsConfig.Provider, *dnsConfig.Ak, *dnsConfig.Sk, *dnsConfig.Token})
				if err != nil {
					return err
				}
			} else {
				err := table.AddRow([]string{key})
				if err != nil {
					return err
				}
			}
		}
		return printTableAndSavaToJSONFile(table, ctx)
	}
}

func importConfigAction() cli.ActionFunc {
	return func(context *cli.Context) error {
		//readConfig, iotool := config.ReadConfig()
		//if iotool != nil {
		//	return iotool
		//}
		//
		return nil
	}
}
