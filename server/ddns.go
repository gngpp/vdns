package server

import (
	"fmt"
	"vdns/config"
	"vdns/lib/api"
	"vdns/lib/api/model"
	"vdns/lib/standard/record"
	"vdns/lib/util/strs"
	"vdns/lib/util/vjson"
	"vdns/lib/util/vnet"
	"vdns/lib/vlog"
)

type DDNS struct {
}

func (d *DDNS) Resolve() error {
	conf, err := config.ReadVdnsConfig()
	if err != nil {
		return err
	}
	for _, providerConfig := range conf.GetProviderConfigList() {
		provider, err := providerConfig.ToVdnsProvider()
		if err != nil {
			vlog.Error(err)
		}
		switch provider.(type) {
		case *api.HuaweiProvider:
			{
			}
		case *api.AliDNSProvider, *api.DNSPodProvider, *api.CloudflareProvider:
			{
				var err error
				err = d.resolveDnsRecordForIpv4Handler(providerConfig.V4, provider)
				if err != nil {
					return err
				}

				err = d.resolveDnsRecordForIpv6Handler(providerConfig.V6, provider)
				if err != nil {
					return err
				}
			}

		}

	}
	return nil
}

func (d *DDNS) resolveDnsRecordForIpv4Handler(v4 config.IP, provider api.VdnsProvider) error {
	if v4.Enabled {
		var ipArr []string
		if v4.OnCard && len(v4.DomainList) > 0 {
			ipArr = vnet.GetIpv4AddrForName(v4.Card)
		} else {
			var ip string
			if strs.IsEmpty(v4.Api) {
				ip = vnet.GetIpv4Addr()
			} else {
				ip = vnet.GetIpv4AddrForUrl(v4.Api)
			}
			ipArr = []string{ip}
		}

		for _, ip := range ipArr {
			fmt.Println(ip)
		}
		err := provider.Support(record.A)
		if err != nil {
			vlog.Debugf("provider support error: %v", err)
			return err
		}
		for _, domain := range v4.DomainList {
			describeDomainRecordsRequest := model.NewDescribeDomainRecordsRequest().
				SetDomain(domain).
				SetRecordType(record.A).
				SetPageSize(9999)
			records, err := provider.DescribeRecords(describeDomainRecordsRequest)
			if err != nil {
				vlog.Debugf("desribe record error: %v", err)
				return err
			}
			if records != nil && records.Records != nil {
				fmt.Println(vjson.PrettifyString(records))
			}
		}

	}
	return nil
}

func (d *DDNS) resolveDnsRecordForIpv6Handler(v6 config.IP, provider api.VdnsProvider) error {
	if v6.Enabled {
		err := provider.Support(record.A)
		if err != nil {
			vlog.Debugf("provider support error: %v", err)
			return err
		}
		if len(v6.DomainList) > 0 {
			for _, domain := range v6.DomainList {
				describeDomainRecordsRequest := model.NewDescribeDomainRecordsRequest().
					SetDomain(domain).
					SetRecordType(record.AAAA).
					SetPageSize(9999)
				records, err := provider.DescribeRecords(describeDomainRecordsRequest)
				if err != nil {
					vlog.Debugf("desribe record error: %v", err)
					return err
				}
				if records != nil && records.Records != nil {
					fmt.Println(vjson.PrettifyString(records))
				}
			}
		}

	}
	return nil
}
