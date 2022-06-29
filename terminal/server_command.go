package terminal

import (
	"fmt"
	"github.com/kardianos/service"
	"github.com/liushuochen/gotable"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"vdns/lib/util/convert"
	"vdns/lib/util/strs"
	"vdns/lib/vlog"
	"vdns/server"
)

func ServerCommand() *cli.Command {
	var controlAction = [6]string{"exec", "start", "stop", "restart", "install", "uninstall"}
	var subCommandList = make([]*cli.Command, 6)
	for index, c := range controlAction {
		if c == "exec" {
			subCommandList[index] = &cli.Command{
				Name:  c,
				Usage: "Exec vdns server",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "interval",
						Aliases: []string{"i"},
						Usage:   "Interval execution time",
						Value:   5,
					},
				},
				Action: func(ctx *cli.Context) error {
					i := ctx.Int("interval")
					return handleServer(i)
				},
			}
		} else {
			if c == "install" {
				subCommandList[index] = &cli.Command{
					Name:  c,
					Usage: strFirstToUpper(c) + " vdns service",
					Flags: []cli.Flag{
						&cli.IntFlag{
							Name:    "interval",
							Aliases: []string{"i"},
							Usage:   "Interval execution time",
							Value:   5,
						},
					},
					Action: func(ctx *cli.Context) error {
						i := ctx.Int("interval")
						return handleServer(i)
					},
				}
			} else {
				subCommandList[index] = &cli.Command{
					Name:  c,
					Usage: strFirstToUpper(c) + " vdns service",
					Action: func(_ *cli.Context) error {
						// default time interval
						return handleServer(-1)
					},
				}
			}
		}
	}

	table, err := gotable.Create("Usage")
	if err != nil {
		vlog.Fatalf("creating error")
	}
	_ = table.AddRow([]string{"Service will install / un-install, start / stop, and run a program as a service (daemon)."})
	_ = table.AddRow([]string{"Currently supports Windows XP+, Linux/(systemd | Upstart | SysV), and OSX/Launchd."})
	return &cli.Command{
		Name:        "server",
		Usage:       "Use vdns server (support DDNS)",
		Description: table.String(),
		Subcommands: subCommandList,
	}
}

func handleServer(interval int) error {
	cfg := &service.Config{
		Name:        "vdns",
		DisplayName: "vdns server",
		Description: "This is an vdns Go service.",
		Arguments:   []string{"server", "exec", "-i", convert.AsStringValue(interval)},
	}

	vdns := server.NewVdns(interval)
	vdnsService, err := service.New(&vdns, cfg)
	if err != nil {
		return err
	}
	vlog.Debugf("running args: %v", os.Args)
	if len(os.Args) >= 3 && os.Args[2] != "exec" {
		fmt.Println(os.Args[2])
		err = service.Control(vdnsService, os.Args[2])
		if err != nil {
			return err
		}
	} else {
		err = vdnsService.Run()
		if err != nil {
			return err
		}
	}

	return err
}

func strFirstToUpper(str string) string {
	upper := strings.ToUpper(string(str[0]))
	return strs.Concat(upper, str[1:])
}
