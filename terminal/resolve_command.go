package terminal

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/urfave/cli/v2"
	"vdns/config"
	"vdns/lib/api"
	"vdns/lib/api/models"
	"vdns/lib/standard/record"
	"vdns/lib/util/convert"
)

//goland:noinspection SpellCheckingInspection
func ResolveRecord() *cli.Command {
	alidnsCommandName := "alidns"
	dnspodCommandName := "dnspod"
	huaweidnsCommandName := "huaweidns"
	cloudflareCommandName := "cloudflare"
	return &cli.Command{
		Name:  "resolve",
		Usage: "Resolving dns records",
		Subcommands: []*cli.Command{
			{
				Name:    alidnsCommandName,
				Aliases: []string{convert.AsStringValue(string(alidnsCommandName[0]))},
				Usage:   "Resolve " + config.AlidnsProvider + " records",
				Subcommands: []*cli.Command{
					describeDNSRecord(config.AlidnsProvider),
					createDNSRecord(config.AlidnsProvider),
					updateDNSRecord(config.AlidnsProvider),
					deleteDNSRecord(config.AlidnsProvider),
				},
			},
			{
				Name:    dnspodCommandName,
				Aliases: []string{convert.AsStringValue(string(dnspodCommandName[0]))},
				Usage:   "Resolve " + config.DnspodProvider + " records",
				Subcommands: []*cli.Command{
					describeDNSRecord(config.DnspodProvider),
					createDNSRecord(config.DnspodProvider),
					updateDNSRecord(config.DnspodProvider),
					deleteDNSRecord(config.DnspodProvider),
				},
			},
			{
				Name:    huaweidnsCommandName,
				Aliases: []string{convert.AsStringValue(string(huaweidnsCommandName[0]))},
				Usage:   "Resolve " + config.HuaweiDnsProvider + " records",
				Subcommands: []*cli.Command{
					describeDNSRecord(config.HuaweiDnsProvider),
					createDNSRecord(config.HuaweiDnsProvider),
					updateDNSRecord(config.HuaweiDnsProvider),
					deleteDNSRecord(config.HuaweiDnsProvider),
				},
			},
			{
				Name:    cloudflareCommandName,
				Aliases: []string{convert.AsStringValue(string(cloudflareCommandName[0]))},
				Usage:   "Resolve " + config.CloudflareProvider + " records",
				Subcommands: []*cli.Command{
					describeDNSRecord(config.CloudflareProvider),
					createDNSRecord(config.CloudflareProvider),
					updateDNSRecord(config.CloudflareProvider),
					deleteDNSRecord(config.CloudflareProvider),
				},
			},
		},
	}
}

func describeDNSRecord(providerKey string) *cli.Command {
	var pageSize int64
	var pageNumber int64
	var domain string
	var recordType string
	var rrKeyWork string
	var valueKeyWork string
	return &cli.Command{
		Name:  "describe",
		Usage: "Describe " + providerKey + " records",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:        "ps",
				Usage:       "page size",
				Value:       10,
				Destination: &pageSize,
			},
			&cli.Int64Flag{
				Name:        "pn",
				Usage:       "page number",
				Value:       1,
				Destination: &pageNumber,
			},
			&cli.StringFlag{
				Name:        "domain",
				Usage:       "record domain",
				Destination: &domain,
			},
			&cli.StringFlag{
				Name:        "type",
				Usage:       "record type",
				Destination: &recordType,
			},
			&cli.StringFlag{
				Name:        "rk",
				Usage:       "the keywords recorded by the host, (fuzzy matching before and after) pattern search, are not case-sensitive",
				Destination: &rrKeyWork,
			},
			&cli.StringFlag{
				Name:        "vk",
				Usage:       "the record value keyword (fuzzy match before and after) pattern search, not case-sensitive",
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
				_ = table.AddRow([]string{
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

func createDNSRecord(providerKey string) *cli.Command {
	var domain string
	var recordType string
	var value string
	return &cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "Create " + providerKey + " record",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "domain",
				Usage:       "domain record",
				Destination: &domain,
			},
			&cli.StringFlag{
				Name:        "type",
				Usage:       "domain record type",
				Destination: &recordType,
			},
			&cli.StringFlag{
				Name:        "value",
				Usage:       "domain record value",
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

func updateDNSRecord(providerKey string) *cli.Command {
	var id string
	var domain string
	var recordType string
	var value string
	return &cli.Command{
		Name:    "update",
		Aliases: []string{"u"},
		Usage:   "Update " + providerKey + " record",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "id",
				Usage:       "domain record identifier",
				Destination: &id,
			},
			&cli.StringFlag{
				Name:        "domain",
				Usage:       "domain record",
				Destination: &domain,
			},
			&cli.StringFlag{
				Name:        "type",
				Usage:       "domain record type",
				Destination: &recordType,
			},
			&cli.StringFlag{
				Name:        "value",
				Usage:       "domain record value",
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

func deleteDNSRecord(providerKey string) *cli.Command {
	var id string
	var domain string
	return &cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "Delete " + providerKey + " record",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "id",
				Usage:       "record identifier",
				Destination: &id,
			},
			&cli.StringFlag{
				Name:        "domain",
				Usage:       "record super domain",
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
	if providerKey == config.AlidnsProvider {
		provider = api.NewAliDNSProvider(credentials)
	}
	if providerKey == config.DnspodProvider {
		provider = api.NewDNSPodProvider(credentials)
	}
	if providerKey == config.CloudflareProvider {
		provider = api.NewCloudflareProvider(credentials)
	}
	if providerKey == config.HuaweiDnsProvider {
		provider = api.NewHuaweiProvider(credentials)
	}
	return provider, nil
}
