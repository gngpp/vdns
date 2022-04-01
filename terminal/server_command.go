package terminal

import (
	"github.com/kardianos/service"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"vdns/lib/util/strs"
	"vdns/lib/vlog"
	"vdns/server"
)

var vdns = server.Vdns{}

func ServerCommand() *cli.Command {
	var controlAction = [6]string{"exec", "start", "stop", "restart", "install", "uninstall"}
	var subCommand = make([]*cli.Command, 6)
	for index, c := range controlAction {
		subCommand[index] = &cli.Command{
			Name:  c,
			Usage: strFirstToUpper(c) + " vdns service",
			Action: func(_ *cli.Context) error {
				handleServer()
				return nil
			},
		}
	}
	return &cli.Command{
		Name:        "server",
		Usage:       "Use vdns server (support DDNS)",
		Subcommands: subCommand,
	}
}

func handleServer() {
	cfg := &service.Config{
		Name:        "vdns",
		DisplayName: "vdns server",
		Description: "This is an vdns Go service.",
		Arguments:   []string{"server", "exec"},
	}
	s, err := service.New(&vdns, cfg)
	if err != nil {
		vlog.Fatal(err)
	}
	logger, err := s.Logger(nil)
	if err != nil {
		vlog.Fatal(err)
	}
	if len(os.Args) == 3 && os.Args[2] != "exec" {
		err = service.Control(s, os.Args[2])
		if err != nil {
			_ = logger.Error(err)
		}
	} else {
		err = s.Run()
		if err != nil {
			_ = logger.Error(err)
		}
	}
	if err != nil {
		_ = logger.Error(err)
	}
}

func strFirstToUpper(str string) string {
	upper := strings.ToUpper(string(str[0]))
	return strs.Concat(upper, str[1:])
}
