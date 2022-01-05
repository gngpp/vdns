package test

import (
	"fmt"
	"testing"
	"vdns/lib/api"
	"vdns/lib/api/models"
	"vdns/lib/auth"
	"vdns/lib/standard/record"
	"vdns/vlog"
)

func Test_AliyunDnsProvider_DescribeDnsRecord(t *testing.T) {
	vlog.SetLevel(vlog.Level.DEBUG)
	credential := auth.NewBasicCredential("LTAI5tSB1faJDcbNo3yXHCec", "24GruqqDNKYICbDMIZfNS9IstVw4IQ")

	aliyunDnsProvider := api.NewAlidnsProvider(credential)
	if aliyunDnsProvider.Support(record.A) {
		request := models.NewDescribeDomainRecordsRequest().
			SetDomain("innas.cn").
			SetPageSize(10).
			SetPageNumber(1)
		resp, err := aliyunDnsProvider.DescribeRecords(request)
		if err != nil {
			vlog.Error(err)
		}
		vlog.Info(resp.String())
	}
}

func TestAliyunDnsProvider_CreateDnsRecord(t *testing.T) {
	credential := auth.NewBasicCredential("LTAI5tSB1faJDcbNo3yXHCec", "24GruqqDNKYICbDMIZfNS9IstVw4IQ")

	aliyunDnsProvider := api.NewAlidnsProvider(credential)
	domain := "hanbi.innas.cn"
	if aliyunDnsProvider.Support(record.A) {
		request := models.NewCreateDomainRecordRequest().
			SetDomain(domain).
			SetValue("192.168.2.7").
			SetRecordType(record.A)
		result, err := aliyunDnsProvider.CreateRecord(request)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result.String())
	}
}

func TestAliyunDnsProvider_UpdateDnsRecord(t *testing.T) {
	credential := auth.NewBasicCredential("LTAI5tSB1faJDcbNo3yXHCec", "24GruqqDNKYICbDMIZfNS9IstVw4IQ")

	aliyunDnsProvider := api.NewAlidnsProvider(credential)
	domain := "ssss.innas.cn"
	if aliyunDnsProvider.Support(record.A) {
		request := models.NewUpdateDomainRecordRequest().
			SetID("738123552390353920").
			SetDomain(domain).
			SetValue("192.168.2.181").
			SetRecordType(record.A)
		result, err := aliyunDnsProvider.UpdateRecord(request)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result.String())
	}
}

func TestAliyunDnsProvider_DeleteDnsRecord(t *testing.T) {
	credential := auth.NewBasicCredential("LTAI5tSB1faJDcbNo3yXHCec", "24GruqqDNKYICbDMIZfNS9IstVw4IQ")

	aliyunDnsProvider := api.NewAlidnsProvider(credential)
	if aliyunDnsProvider.Support(record.A) {
		request := models.NewDeleteDomainRecordRequest().
			SetID("738755789635027968")
		result, err := aliyunDnsProvider.DeleteRecord(request)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result.String())
	}

}
