package command

import (
	"fmt"

	"github.com/liushuochen/gotable"
	"github.com/urfave/cli/v2"
)

func ShowInfoCommand() *cli.Command {
	return &cli.Command{
		Name:    "show",
		Aliases: []string{"s"},
		Usage:   "show vdns information",
		Subcommands: []*cli.Command{
			{
				Name:    "provider",
				Aliases: []string{"p"},
				Usage:   "support providers",
				Action: func(_ *cli.Context) error {
					table, err := gotable.Create("provider", "api document")
					if err != nil {
						fmt.Println("Create table failed: ", err.Error())
						return err
					}
					table.AddRow([]string{"AliDNS", "https://help.aliyun.com/document_detail/39863.html"})
					table.AddRow([]string{"DNSPod", "https://cloud.tencent.com/document/product/1427"})
					table.AddRow([]string{"Cloudflare", "https://api.cloudflare.com/#dns-records-for-a-zone-properties"})
					table.AddRow([]string{"HuaweiDNS", "https://support.huaweicloud.com/function-dns/index.html"})
					fmt.Printf("%v\n", table)
					return nil
				},
			},
			{
				Name:    "record",
				Aliases: []string{"r"},
				Usage:   "supports record types",
				Action: func(_ *cli.Context) error {
					fmt.Println("supports record types: A|AAAA|NS|MX|CNAME|TXT|SRV|CA|REDIRECT_URL|FORWARD_URL")
					table, err := gotable.Create("type", "value", "description")
					if err != nil {
						return err
					}
					table.AddRow([]string{"A", "A", "将域名指向一个IPV4地址"})
					table.AddRow([]string{"AAAA", "AAAA", "将域名指向一个IPV6地址"})
					table.AddRow([]string{"NS", "NS", "将子域名指定其他DNS服务器解析"})
					table.AddRow([]string{"MX", "MX", "将域名指向邮件服务器地址"})
					table.AddRow([]string{"CNAME", "CNAME", "将域名指向另外一个域名"})
					table.AddRow([]string{"TXT", "TXT", "文本长度限制512，通常做SPF记录（反垃圾邮件）"})
					table.AddRow([]string{"SRV", "SRV", "记录提供特定的服务的服务器"})
					table.AddRow([]string{"CA", "CA", "CA证书颁发机构授权校验"})
					table.AddRow([]string{"REDIRECT_URL", "REDIRECT_URL", "将域名重定向到另外一个地址"})
					table.AddRow([]string{"FORWARD_URL", "FORWARD_URL", "显性URL类似，但是会隐藏真实目标地址"})
					fmt.Printf("%v\n", table)
					fmt.Println("reference: https://help.aliyun.com/document_detail/29805.html?spm=a2c4g.11186623.0.0.30e73067AXxwak")
					return nil
				},
			},
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
