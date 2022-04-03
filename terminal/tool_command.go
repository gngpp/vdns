package terminal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
	"net/http"
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
			client := vhttp.CreateClient()
			request, err := http.NewRequest(vhttp.HttpMethodGet.String(), url, nil)
			request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36")
			if err != nil {
				return nil
			}
			response, err := client.Do(request)
			body := response.Body
			defer func(body io.ReadCloser) {
				err := body.Close()
				if err != nil {
					vlog.Fatal(err)
				}
			}(body)
			bytes, err := ioutil.ReadAll(body)
			fmt.Printf("\n%v\n", string(bytes))
			return nil
		},
	}
}
