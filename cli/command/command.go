package command

import (
	"fmt"
	"vdns/cli/config"
	"vdns/lib/api"
	"vdns/lib/api/models"
	"vdns/lib/standard/record"
	"vdns/lib/util/convert"
	"vdns/lib/util/strs"

	"github.com/liushuochen/gotable"
	"github.com/urfave/cli/v2"
)

//goland:noinspection SpellCheckingInspection
func ShowInfoCommand() *cli.Command {
	return &cli.Command{
		Name:    "show",
		Aliases: []string{"s"},
		Usage:   "Show vdns information.",
		Subcommands: []*cli.Command{
			{
				Name:    "provider",
				Aliases: []string{"p"},
				Usage:   "support providers.",
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
				Usage:   "supports record types.",
				Action: func(_ *cli.Context) error {
					fmt.Println("supports record types: A、AAAA、NS、MX、CNAME、TXT、SRV、CA、REDIRECT_URL、FORWARD_URL")
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
		Usage:   "Configure dns service provider access key pair.",
		Subcommands: []*cli.Command{
			configCommand("alidns", config.ALIDNS_PROVIDER),
			configCommand("dnspod", config.DNSPOD_PROVIDER),
			configCommand("huaweidns", config.HUAWERI_DNS_PROVIDER),
			configCommand("cloudflare", config.CLOUDFLARE_PROVIDER),
			{
				Name:  "cat",
				Usage: "Print all dns configuration.",
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
		Usage: "Configure " + commandName + " access key pair.",
		Subcommands: []*cli.Command{
			{
				Name:  "cat",
				Usage: "Print dns provider configuration.",
				Action: func(_ *cli.Context) error {
					readConfig, err := config.ReadConfig()
					if err != nil {
						return err
					}
					dnsConfig := readConfig.ConfigsMap.Get(providerKey)
					table, err := gotable.Create("provider", "ak", "sk", "token")
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
				Usage:       "api accessKey.",
				Destination: &ak,
			},
			&cli.StringFlag{
				Name:        "sk",
				Usage:       "api secretKey.",
				Destination: &sk,
			},
			&cli.StringFlag{
				Name:        "token",
				Usage:       "api token.",
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

//goland:noinspection SpellCheckingInspection
func ResolveDNSRecord() *cli.Command {
	alidnsCommandName := "alidns"
	dnspodCommandName := "dnspod"
	huaweidnsCommandName := "huaweidns"
	cloudflareCommandName := "cloudflare"
	return &cli.Command{
		Name:    "resolve",
		Usage:   "Resolving DNS records.",
		Aliases: []string{"r"},
		Subcommands: []*cli.Command{
			{
				Name:  alidnsCommandName,
				Usage: "resolve " + config.ALIDNS_PROVIDER + " DNS records.",
				Subcommands: []*cli.Command{
					DescribeDNSRecord(config.ALIDNS_PROVIDER),
					CreateDNSRecord(config.ALIDNS_PROVIDER),
					UpdateDNSRecord(config.ALIDNS_PROVIDER),
					DeleteDNSRecord(config.ALIDNS_PROVIDER),
				},
			},
			{
				Name:  dnspodCommandName,
				Usage: "resolve " + config.DNSPOD_PROVIDER + " DNS records.",
				Subcommands: []*cli.Command{
					DescribeDNSRecord(config.DNSPOD_PROVIDER),
					CreateDNSRecord(config.DNSPOD_PROVIDER),
					UpdateDNSRecord(config.DNSPOD_PROVIDER),
					DeleteDNSRecord(config.DNSPOD_PROVIDER),
				},
			},
			{
				Name:  huaweidnsCommandName,
				Usage: "resolve " + config.HUAWERI_DNS_PROVIDER + " DNS records.",
				Subcommands: []*cli.Command{
					DescribeDNSRecord(config.HUAWERI_DNS_PROVIDER),
					CreateDNSRecord(config.HUAWERI_DNS_PROVIDER),
					UpdateDNSRecord(config.HUAWERI_DNS_PROVIDER),
					DeleteDNSRecord(config.HUAWERI_DNS_PROVIDER),
				},
			},
			{
				Name:  cloudflareCommandName,
				Usage: "resolve " + config.CLOUDFLARE_PROVIDER + " DNS records.",
				Subcommands: []*cli.Command{
					DescribeDNSRecord(config.CLOUDFLARE_PROVIDER),
					CreateDNSRecord(config.CLOUDFLARE_PROVIDER),
					UpdateDNSRecord(config.CLOUDFLARE_PROVIDER),
					DeleteDNSRecord(config.CLOUDFLARE_PROVIDER),
				},
			},
		},
	}
}

func DescribeDNSRecord(providerKey string) *cli.Command {
	var pageSize int64
	var pageNumber int64
	var domain string
	var recordType string
	var rrKeyWork string
	var valueKeyWork string
	return &cli.Command{
		Name:    "search",
		Aliases: []string{"s"},
		Usage:   "describe " + providerKey + " DNS records.",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:        "ps",
				Usage:       "page size.",
				Value:       5,
				Destination: &pageSize,
			},
			&cli.Int64Flag{
				Name:        "pn",
				Usage:       "page number.",
				Value:       1,
				Destination: &pageNumber,
			},
			&cli.StringFlag{
				Name:        "domain",
				Usage:       "record domain.",
				Destination: &domain,
			},
			&cli.StringFlag{
				Name:        "type",
				Usage:       "record type.",
				Destination: &recordType,
			},
			&cli.StringFlag{
				Name:        "rk",
				Usage:       "the keywords recorded by the host, (fuzzy matching before and after) pattern search, are not case-sensitive.",
				Destination: &rrKeyWork,
			},
			&cli.StringFlag{
				Name:        "vk",
				Usage:       "the record value keyword (fuzzy match before and after) pattern search, not case-sensitive.",
				Destination: &valueKeyWork,
			},
		},
		Action: func(_ *cli.Context) error {
			provider, err := getProvider(providerKey)
			if err != nil {
				return err
			}
			request := models.NewDescribeDomainRecordsRequest()
			request.Domain = &domain
			request.PageSize = &pageSize
			request.PageNumber = &pageNumber
			request.ValueKeyWord = &valueKeyWork
			request.RRKeyWord = &rrKeyWork
			//goland:noinspection GoRedundantConversion
			ofType, b := record.OfType(record.Type(recordType))
			if b {
				request.RecordType = ofType
			}
			describeRecords, err := provider.DescribeRecords(request)
			if err != nil {
				return err
			}
			if *describeRecords.ListCount > 0 {
				table, err := gotable.CreateByStruct(describeRecords.Records[0])
				if err != nil {
					return err
				}
				for _, record := range describeRecords.Records {
					_ = table.AddRow([]string{*record.ID, record.RecordType.String(), *record.RR, *record.Domain, *record.Value, *record.Status, convert.AsStringValue(*record.TTL)})
				}
				fmt.Print(table)
				table, err = gotable.Create("Total", "PageSize", "PageNumber")
				if err != nil {
					return err
				}
				table.AddRow([]string{
					convert.AsStringValue(*describeRecords.TotalCount),
					convert.AsStringValue(*describeRecords.PageSize),
					convert.AsStringValue(*describeRecords.PageNumber),
				})
				fmt.Println(table)
			}
			return nil
		},
	}
}

func CreateDNSRecord(providerKey string) *cli.Command {
	var domain string
	var recordType string
	var value string
	return &cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "create " + providerKey + " DNS record.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "domain",
				Usage:       "domain record.",
				Destination: &domain,
			},
			&cli.StringFlag{
				Name:        "type",
				Usage:       "domain record type.",
				Destination: &recordType,
			},
			&cli.StringFlag{
				Name:        "value",
				Usage:       "domain record value.",
				Destination: &value,
			},
		},
		Action: func(_ *cli.Context) error {
			provider, err := getProvider(providerKey)
			if err != nil {
				return err
			}
			request := models.NewCreateDomainRecordRequest()
			request.Domain = &domain
			request.Value = &value
			//goland:noinspection GoRedundantConversion
			ofType, b := record.OfType(record.Type(recordType))
			if b {
				request.RecordType = &ofType
			}
			createRecord, err := provider.CreateRecord(request)
			if err != nil {
				return err
			}
			table, err := gotable.CreateByStruct(createRecord)
			if err != nil {
				return err
			}
			_ = table.AddRow([]string{*createRecord.RequestId, *createRecord.RecordId})
			fmt.Println(table)
			return nil
		},
	}
}

func UpdateDNSRecord(providerKey string) *cli.Command {
	var id string
	var domain string
	var recordType string
	var value string
	return &cli.Command{
		Name:    "update",
		Aliases: []string{"u"},
		Usage:   "update " + providerKey + " DNS record.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "id",
				Usage:       "domain record identifier",
				Destination: &id,
			},
			&cli.StringFlag{
				Name:        "domain",
				Usage:       "domain record.",
				Destination: &domain,
			},
			&cli.StringFlag{
				Name:        "type",
				Usage:       "domain record type.",
				Destination: &recordType,
			},
			&cli.StringFlag{
				Name:        "value",
				Usage:       "domain record value.",
				Destination: &value,
			},
		},
		Action: func(_ *cli.Context) error {
			provider, err := getProvider(providerKey)
			if err != nil {
				return err
			}
			request := models.NewUpdateDomainRecordRequest()
			request.ID = &id
			request.Domain = &domain
			request.Value = &value
			//goland:noinspection GoRedundantConversion
			ofType, b := record.OfType(record.Type(recordType))
			if b {
				request.RecordType = &ofType
			}
			updateRecord, err := provider.UpdateRecord(request)
			if err != nil {
				return err
			}
			table, err := gotable.CreateByStruct(updateRecord)
			if err != nil {
				return err
			}
			_ = table.AddRow([]string{*updateRecord.RequestId, *updateRecord.RecordId})
			fmt.Println(table)
			return nil
		},
	}
}

func DeleteDNSRecord(providerKey string) *cli.Command {
	var id string
	var domain string
	return &cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete " + providerKey + " DNS record.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "id",
				Usage:       "record identifier.",
				Destination: &id,
			},
			&cli.StringFlag{
				Name:        "domain",
				Usage:       "record super domain.",
				Destination: &domain,
			},
		},
		Action: func(_ *cli.Context) error {
			provider, err := getProvider(providerKey)
			if err != nil {
				return err
			}
			request := models.NewDeleteDomainRecordRequest()
			request.Domain = &domain
			request.ID = &id
			deleteRecord, err := provider.DeleteRecord(request)
			if err != nil {
				return err
			}
			table, err := gotable.CreateByStruct(deleteRecord)
			if err != nil {
				return err
			}
			_ = table.AddRow([]string{*deleteRecord.RequestId, *deleteRecord.RecordId})
			fmt.Println(table)
			return nil
		},
	}
}

func getProvider(providerKey string) (api.VdnsProvider, error) {
	credentials, err := config.ReadCredentials(providerKey)
	if err != nil {
		return nil, err
	}
	var provider api.VdnsProvider
	if providerKey == config.ALIDNS_PROVIDER {
		provider = api.NewAliDNSProvider(credentials)
	}
	if providerKey == config.DNSPOD_PROVIDER {
		provider = api.NewDNSPodProvider(credentials)
	}
	if providerKey == config.CLOUDFLARE_PROVIDER {
		provider = api.NewCloudflareProvider(credentials)
	}
	if providerKey == config.HUAWERI_DNS_PROVIDER {
		provider = api.NewHuaweiProvider(credentials)
	}
	return provider, nil
}
