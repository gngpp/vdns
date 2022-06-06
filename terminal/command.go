package terminal

import (
	"errors"
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
	"math"
	"strings"
	"vdns/config"
	"vdns/lib/util/vhttp"
	"vdns/lib/vlog"
)

//goland:noinspection SpellCheckingInspection
func Command() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "provider",
			Usage: "Support providers",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "path",
					Usage: "sava table to csv filepath",
				},
			},
			Action: providerAction(),
		},
		{
			Name:  "record",
			Usage: "Support record types",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "path",
					Usage: "sava table to csv filepath",
				},
			},
			Action: recordAction(),
		},
		{
			Name:    "print-ip-api",
			Aliases: []string{"pia"},
			Usage:   "Print search ip request api list",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "path",
					Usage: "sava table to csv filepath",
				},
			},
			Action: printIpApiAction(),
		},
		{
			Name:    "test-ip-api",
			Aliases: []string{"tia"},
			Usage:   "Test the API for requesting query ip",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "type",
					Usage: "value is ipv4 or ipv6",
				},
			},
			Action: testIpApiAction(),
		},
		requestCommand(),
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
			req, err := vhttp.Get(url, nil)
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

func recordAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		fmt.Println("Supports record types: A、AAAA、NS、MX、CNAME、TXT、SRV、CA、REDIRECT_URL、FORWARD_URL")
		table, err := gotable.Create("type", "value", "description")
		if err != nil {
			return err
		}
		_ = table.AddRow([]string{"A", "A", "将域名指向一个IPV4地址"})
		_ = table.AddRow([]string{"AAAA", "AAAA", "将域名指向一个IPV6地址"})
		_ = table.AddRow([]string{"NS", "NS", "将子域名指定其他DNS服务器解析"})
		_ = table.AddRow([]string{"MX", "MX", "将域名指向邮件服务器地址"})
		_ = table.AddRow([]string{"CNAME", "CNAME", "将域名指向另外一个域名"})
		_ = table.AddRow([]string{"TXT", "TXT", "文本长度限制512，通常做SPF记录（反垃圾邮件）"})
		_ = table.AddRow([]string{"SRV", "SRV", "记录提供特定的服务的服务器"})
		_ = table.AddRow([]string{"CA", "CA", "CA证书颁发机构授权校验"})
		_ = table.AddRow([]string{"REDIRECT_URL", "REDIRECT_URL", "将域名重定向到另外一个地址"})
		_ = table.AddRow([]string{"FORWARD_URL", "FORWARD_URL", "显性URL类似，但是会隐藏真实目标地址"})

		err = printTableAndSavaToCSVFile(table, ctx)
		if err != nil {
			return err
		}
		fmt.Println("Reference: https://help.aliyun.com/document_detail/29805.html?spm=a2c4g.11186623.0.0.30e73067AXxwak")
		return nil
	}
}

func providerAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		table, err := gotable.Create("provider", "DNS API Document")
		if err != nil {
			return err
		}
		_ = table.AddRow([]string{config.AlidnsProvider, "https://help.aliyun.com/document_detail/39863.html"})
		_ = table.AddRow([]string{config.DnspodProvider, "https://cloud.tencent.com/document/product/1427"})
		_ = table.AddRow([]string{config.CloudflareProvider, "https://api.cloudflare.com/#dns-records-for-a-zone-properties"})
		_ = table.AddRow([]string{config.HuaweiDnsProvider, "https://support.huaweicloud.com/function-dns/index.html"})

		fmt.Print(table)
		path := ctx.String("path")
		err = toCSVFile(table, path)
		if err != nil {
			return err
		}
		return nil
	}
}

func printIpApiAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		table, err := gotable.Create("Ipv4 Request API", "Ipv6 Request API")
		if err != nil {
			return err
		}
		ipv4ApiList := config.GetIpv4ApiList()
		ipv6ApiList := config.GetIpv6ApiList()
		abs := int(math.Abs(float64(len(ipv4ApiList) - len(ipv6ApiList))))
		max := int(math.Max(float64(len(ipv4ApiList)), float64(len(ipv6ApiList))))
		if len(ipv4ApiList) > len(ipv6ApiList) {
			for i := 0; i < abs; i++ {
				ipv6ApiList = append(ipv6ApiList, "")
			}
		} else {
			for i := 0; i < abs; i++ {
				ipv4ApiList = append(ipv4ApiList, "")
			}
		}
		for i := 0; i < max; i++ {
			_ = table.AddRow([]string{ipv4ApiList[i], ipv6ApiList[i]})
		}

		return printTableAndSavaToCSVFile(table, ctx)
	}
}

func testIpApiAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		ipType := ctx.String("type")
		if (ipType != "ipv4") && (ipType != "ipv6") {
			return errors.New("ip type must be: ipv4 or ipv6.\nfor example: --type=ipv4 or --type=ipv6")
		}
		var t *table.Table
		var err error
		var ipApiList []string
		if ipType == "ipv4" {
			t, err = gotable.Create("Ipv4 Request API", "Status")
			if err != nil {
				return err
			}
			ipApiList = config.GetIpv4ApiList()
		}
		if ipType == "ipv6" {
			t, err = gotable.Create("Ipv6 Request API", "Status")
			if err != nil {
				return err
			}
			ipApiList = config.GetIpv6ApiList()
		}
		for _, api := range ipApiList {
			_ = t.AddRow([]string{api, vhttp.GetIpv4AddrForUrl(api)})
		}
		fmt.Println(t)
		return err
	}
}
