package server

import (
	"errors"
	"fmt"
	"vdns/config"
	"vdns/lib/api"
	"vdns/lib/api/model"
	"vdns/lib/standard/record"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
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
	for _, pc := range conf.GetProviderConfigList() {
		provider, err := pc.ToVdnsProvider()
		if err != nil {
			vlog.Error(err)
		}
		switch provider.(type) {
		case *api.HuaweiProvider:
			{
			}
		case *api.AliDNSProvider, *api.DNSPodProvider, *api.CloudflareProvider:
			{
				go d.handler(pc.V4, provider, pc.Provider)
				go d.handler(pc.V6, provider, pc.Provider)
			}

		}

	}
	return nil
}

func (d DDNS) handler(ipv config.IP, provider api.VdnsProvider, providerName string) {
	vlog.Infof("[%v][%v] start processing...", providerName, ipv.Type)
	err := d.resolveDnsRecordHandler(ipv, provider, providerName)
	if err != nil {
		vlog.Errorf("[%v][%v] parsing dns record error: %v", providerName, ipv.Type, err)
	}
	vlog.Infof("[%v][%v] processing ends...", providerName, ipv.Type)
}

func (d *DDNS) resolveDnsRecordHandler(ipv config.IP, provider api.VdnsProvider, providerName string) error {
	vlog.Debugf("ip config: %v", ipv)
	if ipv.Enabled {
		var ip string
		var r record.Type
		if ipv.Ipv4() {
			ip = d.getPubIpv4Addr(ipv)
			r = record.A
		} else if ipv.Ipv6() {
			ip = d.getPubIpv6Addr(ipv)
			r = record.AAAA
		} else {
			return errors.New(fmt.Sprintf("[%v] unknown ip type", providerName))
		}
		if strs.NotEmpty(ip) {
			err := provider.Support(r)
			if err != nil {
				vlog.Debugf("[%v] provider support error: %v", providerName, err)
				return err
			}
			for _, domain := range ipv.DomainList {
				request := model.NewDescribeDomainRecordsRequest().
					SetDomain(domain).
					SetRecordType(r)
				records, err := provider.DescribeRecords(request)
				if err != nil {
					vlog.Debugf("[%v] desribe record error: %v", providerName, err)
					return err
				}
				// If the domain name record exists, update the record
				if records != nil && records.Records != nil && len(records.Records) > 0 {
					for _, res := range records.Records {
						fullDomain := res.FullDomain()
						err := vhttp.CheckDomain(fullDomain)
						if err != nil {
							vlog.Errorf("[%v] query record error exception: %v", providerName, err)
							return err
						}
						if res.RecordType == r && fullDomain == domain && ip != strs.StringValue(res.Value) {
							request := model.NewUpdateDomainRecordRequest().
								SetID(strs.StringValue(res.ID)).
								SetRecordType(r).
								SetDomain(domain).
								SetValue(ip)
							response, err := provider.UpdateRecord(request)
							if err != nil {
								vlog.Errorf("[%v] failed to update record: %v", providerName, err)
								return err
							}
							vlog.Infof("[%v] update domain record: %v --- %v, record id: %v", providerName, domain, ip, strs.StringValue(response.RecordId))
						} else {
							vlog.Infof("[%v] domain name records have not changed: %v --- %v", providerName, domain, ip)
						}
					}
				} else {
					// Domain record does not exist, create record
					request := model.NewCreateDomainRecordRequest().
						SetDomain(domain).
						SetRecordType(r).
						SetValue(ip)
					response, err := provider.CreateRecord(request)
					if err != nil {
						vlog.Errorf("[%v] failed to create record: %v", providerName, err)
						return err
					}
					vlog.Infof("[%v] create domain record: %v --- %v, record id: %v", providerName, domain, ip, strs.StringValue(response.RecordId))
				}
			}
		}
	}
	return nil
}

// The network card may have multiple public network IPs, it is best to bind a public network IP to each network card
func (d *DDNS) getPubIpv4Addr(v4 config.IP) string {
	var ipArr []string
	if v4.OnCard && len(v4.DomainList) > 0 {
		ipArr = vnet.GetPubIpv4AddrForName(v4.Card)
	} else {
		var ip string
		if strs.IsEmpty(v4.Api) {
			ip = vnet.GetPubIpv4Addr()
		} else {
			ip = vnet.GetPubIpv4AddrForUrl(v4.Api)
		}
		ipArr = []string{ip}
	}
	for _, addr := range ipArr {
		if !vnet.IsPrivateAddr(addr) {
			return addr
		}
	}
	return ""
}

// The network card may have multiple public network IPs, it is best to bind a public network IP to each network card
func (d *DDNS) getPubIpv6Addr(v6 config.IP) string {
	var ipArr []string
	if v6.OnCard && len(v6.DomainList) > 0 {
		ipArr = vnet.GetPubIpv6AddrForName(v6.Card)
	} else {
		var ip string
		if strs.IsEmpty(v6.Api) {
			ip = vnet.GetPubIpv6Addr()
		} else {
			ip = vnet.GetPubIpv6AddrForUrl(v6.Api)
		}
		ipArr = []string{ip}
	}
	for _, addr := range ipArr {
		if vnet.IsPrivateAddr(addr) {
			return addr
		}
	}
	return ""
}
