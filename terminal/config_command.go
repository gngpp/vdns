package terminal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"vdns/config"
	"vdns/lib/util/file"
	"vdns/lib/util/strs"
)

//goland:noinspection SpellCheckingInspection
func ConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Configure DNS service provider access key pair",
		Subcommands: []*cli.Command{
			setConfigCommand(),
			setLogConfigCommand(),
			catConfigCommand(),
			resetConfigCommand(),
			importConfigCommand(),
			exportConfigCommand(),
		},
	}
}

//goland:noinspection SpellCheckingInspection
func setConfigCommand() *cli.Command {
	var provider string
	var ak string
	var sk string
	var token string
	return &cli.Command{
		Name:                   "set",
		Usage:                  "Configure DNS provider access key pair",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "provider",
				Usage:       "DNS provider name",
				Destination: &provider,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "ak",
				Usage:       "API Access key",
				Destination: &ak,
			},
			&cli.StringFlag{
				Name:        "sk",
				Usage:       "API Secret key",
				Destination: &sk,
			},
			&cli.StringFlag{
				Name:        "token",
				Usage:       "API Token",
				Destination: &token,
			},
		},
		Action: func(ctx *cli.Context) error {
			vdnsConfig, err := config.LoadVdnsConfig()
			if err != nil {
				return err
			}
			providerConfig := vdnsConfig.ProviderMap.Get(provider)
			isModify := false
			if !strs.IsEmpty(ak) {
				providerConfig.SetAk(&ak)
				isModify = true
			}
			if !strs.IsEmpty(sk) {
				providerConfig.SetSK(&sk)
				isModify = true
			}
			if !strs.IsEmpty(token) {
				providerConfig.SetToken(&token)
				isModify = true
			}

			if isModify {
				err = config.WriteVdnsConfig(vdnsConfig)
				if err != nil {
					return err
				}
				return vdnsConfig.PrintTable()
			}
			return nil
		},
	}
}

func setLogConfigCommand() *cli.Command {
	var dir string
	var reserveDay int
	var filePrefix string
	return &cli.Command{
		Name:  "set-log",
		Usage: "Set log configuration",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "dir",
				Usage:       "Log storage directory",
				Destination: &dir,
			},
			&cli.IntFlag{
				Name:        "reserve-day",
				Usage:       "log retention date",
				Destination: &reserveDay,
			},
			&cli.StringFlag{
				Name:        "file-prefix",
				Usage:       "Log file prefix",
				Destination: &filePrefix,
			},
		},
		Action: func(context *cli.Context) error {
			vdnsConfig, err := config.LoadVdnsConfig()
			if err != nil {
				return err
			}

			isModify := false
			if !strs.IsEmpty(dir) {
				if !file.IsDir(dir) {
					return fmt.Errorf("system does not exist path or not is dir: %v", dir)
				}
				vdnsConfig.LogDir = dir
				isModify = true
			}

			if !strs.IsEmpty(filePrefix) {
				vdnsConfig.LogFilePrefix = filePrefix
				isModify = true
			}

			if reserveDay > 0 {
				vdnsConfig.LogReserveDay = reserveDay
				isModify = true
			}

			if isModify {
				err = config.WriteVdnsConfig(vdnsConfig)
				if err != nil {
					return err
				}
			}

			return vdnsConfig.PrintTable()
		},
	}
}

func catConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "cat",
		Usage: "Print all DNS configuration",
		Action: func(ctx *cli.Context) error {
			loadConfig, err := config.LoadVdnsConfig()
			if err != nil {
				return err
			}
			return loadConfig.PrintTable()
		},
	}
}

func resetConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "reset",
		Usage: "Reset initial configuration",
		Action: func(context *cli.Context) error {
			vdnsConfig := config.NewVdnsConfig()
			err := config.WriteVdnsConfig(vdnsConfig)
			if err != nil {
				return err
			}
			loadVdnsConfig, err := config.LoadVdnsConfig()
			if err != nil {
				return err
			}
			return loadVdnsConfig.PrintTable()
		},
	}
}

func importConfigCommand() *cli.Command {
	var path string
	return &cli.Command{
		Name:  "import",
		Usage: "Import DNS configuration",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "path",
				Usage:       "sava table to csv filepath",
				Destination: &path,
				Required:    true,
			},
		},
		Action: func(context *cli.Context) error {
			return nil
		},
	}
}

func exportConfigCommand() *cli.Command {
	var path string
	return &cli.Command{
		Name:  "export",
		Usage: "Export DNS configuration",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "path",
				Usage:       "sava data table to csv filepath",
				Destination: &path,
				Required:    true,
			},
		},
		Action: func(context *cli.Context) error {
			return nil
		},
	}
}
