package terminal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
	"strings"
	"vdns/lib/util/vhttp"
	"vdns/lib/vlog"
)

func ToolCommand() *cli.Command {
	return &cli.Command{
		Name:  "tool",
		Usage: "Tool command",
		Subcommands: []*cli.Command{
			requestCommand(),
		},
	}
}

func requestCommand() *cli.Command {
	return &cli.Command{
		Name:  "request",
		Usage: "Request Api (only support get method)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "url",
				Usage: "request url",
			},
		},
		Action: func(ctx *cli.Context) error {
			url := strings.TrimSpace(ctx.String("url"))
			req, err := vhttp.Get(url)
			if err != nil {
				return err
			}
			body := req.Body
			defer func(body io.ReadCloser) {
				err := body.Close()
				if err != nil {
					vlog.Fatal(err)
				}
			}(body)
			bytes, err := ioutil.ReadAll(body)
			fmt.Printf("body: %v", string(bytes))
			return nil
		},
	}
}
