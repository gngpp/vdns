package terminal

import (
	"errors"
	"fmt"
	"github.com/kardianos/service"
	"github.com/liushuochen/gotable"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"vdns/lib/util/convert"
	"vdns/lib/util/strs"
	"vdns/lib/vlog"
	"vdns/server"
)

var shellPathList = []string{os.Getenv("SHELL"), "bash", "sh"}

func findPath(paths []string) (path string, err error) {
	for _, p := range paths {
		path, err = exec.LookPath(p)
		if err == nil {
			break
		}
	}

	return
}

func ServerCommand() *cli.Command {
	var controlAction = []string{"exec", "start", "stop", "restart", "install", "uninstall", "status"}
	var subCommandList = make([]*cli.Command, len(controlAction))
	commandFlags := []cli.Flag{
		&cli.IntFlag{
			Name:    "interval",
			Aliases: []string{"i"},
			Usage:   "interval execution time",
			Value:   5,
		},
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "enable debug mode",
		},
	}

	for index, c := range controlAction {
		if index == len(controlAction)-1 {
			break
		}
		if c == "exec" {
			subCommandList[index] = &cli.Command{
				Name:  c,
				Usage: "Exec vdns server",
				Flags: commandFlags,
				Action: func(ctx *cli.Context) error {
					i := ctx.Int("interval")
					d := ctx.Bool("debug")
					return handleServer(i, d)
				},
			}
		} else {
			if c == "install" {
				subCommandList[index] = &cli.Command{
					Name:  c,
					Usage: strFirstToUpper(c) + " vdns service",
					Flags: commandFlags,
					Action: func(ctx *cli.Context) error {
						i := ctx.Int("interval")
						d := ctx.Bool("debug")
						return handleServer(i, d)
					},
				}
			} else {
				subCommandList[index] = &cli.Command{
					Name:  c,
					Usage: strFirstToUpper(c) + " vdns service",
					Action: func(_ *cli.Context) error {
						// default time interval
						return handleServer(-1, false)
					},
				}
			}
		}
	}

	subCommandList[len(controlAction)-1] = &cli.Command{
		Name:  "status",
		Usage: " Show vdns service status (Windows is not supported)",
		Action: func(_ *cli.Context) error {
			// try to find shell binary
			if shellPath, err := findPath(shellPathList); err == nil {
				cmd := exec.Command(shellPath, "-c", "ps -ef | grep vdns | grep -v grep | awk '{print $2}' | grep -v "+convert.AsStringValue(os.Getpid()))

				stdout, err := cmd.StdoutPipe()
				if err != nil {
					return err
				}
				if err := cmd.Start(); err != nil {
					return err
				}

				bytes, err := ioutil.ReadAll(stdout)
				if err != nil {
					return err
				}

				if stdout.Close() != nil {
					return err
				}

				pid := strings.TrimSpace(string(bytes))

				if strs.IsEmpty(pid) {
					fmt.Println("vdns is stop...")
					return nil
				}
				fmt.Printf("vdns (pid %v) is running...\n", pid)

				if err := cmd.Wait(); err != nil {
					return err
				}

				return nil
			}
			return errors.New("shell not found")
		},
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

func handleServer(interval int, debug bool) error {
	args := []string{"server", "exec", "-i", convert.AsStringValue(interval)}
	if debug {
		args = append(args, "-d")
	}
	cfg := &service.Config{
		Name:        "vdns",
		DisplayName: "vdns server",
		Description: "This is an vdns Go service.",
		Arguments:   args,
	}

	vdns := server.NewVdns(interval, debug)
	vdnsService, err := service.New(&vdns, cfg)
	if err != nil {
		return err
	}
	if len(os.Args) >= 3 && os.Args[2] != "exec" {
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
