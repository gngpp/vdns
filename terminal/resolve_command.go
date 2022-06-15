package terminal

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/urfave/cli/v2"
	"vdns/config"
	"vdns/lib/api"
	"vdns/lib/api/model"
	"vdns/lib/standard/record"
	"vdns/lib/util/convert"
)

//goland:noinspection SpellCheckingInspection
func ResolveRecordList() []*cli.Command {
	strings := []string{config.AlidnsProvider, config.DnspodProvider, config.CloudflareProvider, config.HuaweiDnsProvider}
	var commands []*cli.Command
	for _, commandName := range strings {
		command := &cli.Command{
			Name:    commandName,
			Aliases: []string{convert.AsStringValue(string(commandName[0]))},
			Usage:   "Resolve " + config.AlidnsProvider + " records",
			Subcommands: []*cli.Command{
				describeDNSRecord(commandName),
				createDNSRecord(commandName),
				updateDNSRecord(commandName),
				deleteDNSRecord(commandName),
			},
		}
		commands = append(commands, command)
	}
	return commands
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
			request := model.NewDescribeDomainRecordsRequest()
			request.Domain = &domain
			request.PageSize = &pageSize
			request.PageNumber = &pageNumber
			request.ValueKeyWord = &valueKeyWork
			request.RRKeyWord = &rrKeyWork
			request.RecordType = record.Type(recordType)
			go spinner()
			err = provider.Support(record.Type(recordType))
			if err != nil {
				return err
			}
			describeRecords, err := provider.DescribeRecords(request)
			if *describeRecords.ListCount > 0 {
				table, err := gotable.CreateByStruct(describeRecords.Records[0])
				if err != nil {
					return err
				}
				for _, r := range describeRecords.Records {
					_ = table.AddRow([]string{*r.ID, r.RecordType.String(), *r.RR, *r.Domain, *r.Value, *r.Status, convert.AsStringValue(*r.TTL)})
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
			request := model.NewCreateDomainRecordRequest()
			request.Domain = &domain
			request.Value = &value
			request.RecordType = record.Type(recordType)

			err = provider.Support(record.Type(recordType))
			if err != nil {
				return err
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
			request := model.NewUpdateDomainRecordRequest()
			request.ID = &id
			request.Domain = &domain
			request.Value = &value
			request.RecordType = record.Type(recordType)

			err = provider.Support(record.Type(recordType))
			if err != nil {
				return err
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
			request := model.NewDeleteDomainRecordRequest()
			request.Domain = &domain
			request.ID = &id

			if err != nil {
				return err
			}
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
