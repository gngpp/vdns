package command

import (
	"fmt"
	"vdns/cli/config"
	"vdns/lib/util/strs"

	"github.com/liushuochen/gotable"
	"github.com/urfave/cli/v2"
)

//goland:noinspection SpellCheckingInspection
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
					err = table.AddRow([]string{"AliDNS", "https://help.aliyun.com/document_detail/39863.html"})
					err = table.AddRow([]string{"DNSPod", "https://cloud.tencent.com/document/product/1427"})
					err = table.AddRow([]string{"Cloudflare", "https://api.cloudflare.com/#dns-records-for-a-zone-properties"})
					err = table.AddRow([]string{"HuaweiDNS", "https://support.huaweicloud.com/function-dns/index.html"})
					if err != nil {
						return err
					}
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
					err = table.AddRow([]string{"A", "A", "将域名指向一个IPV4地址"})
					err = table.AddRow([]string{"AAAA", "AAAA", "将域名指向一个IPV6地址"})
					err = table.AddRow([]string{"NS", "NS", "将子域名指定其他DNS服务器解析"})
					err = table.AddRow([]string{"MX", "MX", "将域名指向邮件服务器地址"})
					err = table.AddRow([]string{"CNAME", "CNAME", "将域名指向另外一个域名"})
					err = table.AddRow([]string{"TXT", "TXT", "文本长度限制512，通常做SPF记录（反垃圾邮件）"})
					err = table.AddRow([]string{"SRV", "SRV", "记录提供特定的服务的服务器"})
					err = table.AddRow([]string{"CA", "CA", "CA证书颁发机构授权校验"})
					err = table.AddRow([]string{"REDIRECT_URL", "REDIRECT_URL", "将域名重定向到另外一个地址"})
					err = table.AddRow([]string{"FORWARD_URL", "FORWARD_URL", "显性URL类似，但是会隐藏真实目标地址"})
					if err != nil {
						return err
					}
					fmt.Printf("%v\n", table)
					fmt.Println("reference: https://help.aliyun.com/document_detail/29805.html?spm=a2c4g.11186623.0.0.30e73067AXxwak")
					return nil
				},
			},
		},
	}
}

//goland:noinspection SpellCheckingInspection
func DNSConfigCommand() *cli.Command {
	return &cli.Command{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "Configure dns service provider access key pair",
		Subcommands: []*cli.Command{
			configCommand("alidns", config.ALIDNS_PROVIDER),
			configCommand("dnspod", config.DNSPOD_PROVIDER),
			configCommand("huaweidns", config.HUAWERI_DNS_PROVIDER),
			configCommand("cloudflare", config.CLOUDFLARE_PROVIDER),
			{
				Name:  "cat",
				Usage: "Print all dns configuration",
				Action: func(_ *cli.Context) error {
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
					fmt.Println(table)
					return nil
				},
			},
		},
	}
}

//goland:noinspection SpellCheckingInspection
func configCommand(commandName string, providerKey string) *cli.Command {
	var ak string
	var sk string
	var token string
	return &cli.Command{
		Name:  commandName,
		Usage: "Configure " + commandName + " access key pair",
		Subcommands: []*cli.Command{
			{
				Name:  "cat",
				Usage: "Print dns provider configuration",
				Action: func(_ *cli.Context) error {
					readConfig, err := config.ReadConfig()
					if err != nil {
						return err
					}
					dnsConfig := readConfig.ConfigsMap.Get(providerKey)
					table, err := gotable.Create("provider", "ak", "sk", "token")
					err = table.AddRow([]string{*dnsConfig.Provider, *dnsConfig.Ak, *dnsConfig.Sk, *dnsConfig.Token})
					if err != nil {
						return err
					}
					fmt.Println(table)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "ak",
				Usage:       "api accessKey",
				Destination: &ak,
			},
			&cli.StringFlag{
				Name:        "sk",
				Usage:       "api secretKey",
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
			dnsConfig := readConfig.ConfigsMap.Get(providerKey)
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
