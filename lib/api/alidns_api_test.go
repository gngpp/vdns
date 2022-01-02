package api

import (
	"fmt"
	"testing"
	"vdns/lib/api/models"
	"vdns/lib/auth"
	"vdns/lib/standard/record"
)

func Test_AliyunDnsProvider_DescribeDnsRecord(t *testing.T) {
	credential, err := auth.NewBasicCredential("LTAI5tDsfyHKxratweTjho87", "DtmJl1m0RnZcemU6GtdoWbBTM0Izg5")
	if err != nil {
		fmt.Println(err)
		return
	}
	aliyunDnsProvider := NewAlidnsProvider(credential)
	domain := "innas.cn"
	if aliyunDnsProvider.Support(record.A) {
		request := models.NewDescribeDomainRecordsRequest().
			SetDomain(domain).
			SetRecordType(record.A).
			SetPageSize(10).
			SetPageNumber(1)
		resp, err := aliyunDnsProvider.DescribeRecords(request)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.String())
	}
}

func TestAliyunDnsProvider_CreateDnsRecord(t *testing.T) {
	credential, err := auth.NewBasicCredential("asda", "DtmJl1m0RnZcemU6GtdoWbBTM0Izg5")
	if err != nil {
		fmt.Println(err)
		return
	}
	aliyunDnsProvider := NewAlidnsProvider(credential)
	domain := "hanbi.innas.cn"
	if aliyunDnsProvider.Support(record.A) {
		request := models.NewCreateDomainRecordRequest().
			SetDomain(domain).
			SetValue("192.168.2.4").
			SetRecordType(record.A)
		result, err := aliyunDnsProvider.CreateRecord(request)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result.String())
	}
}

func TestAliyunDnsProvider_UpdateDnsRecord(t *testing.T) {
	credential, err := auth.NewBasicCredential("LTAI5tDsfyHKxratweTjho87", "DtmJl1m0RnZcemU6GtdoWbBTM0Izg5")
	if err != nil {
		fmt.Println(err)
		return
	}
	aliyunDnsProvider := NewAlidnsProvider(credential)
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
	credential, err := auth.NewBasicCredential("LTAI5tDsfyHKxratweTjho87", "DtmJl1m0RnZcemU6GtdoWbBTM0Izg5")
	if err != nil {
		fmt.Println(err)
		return
	}
	aliyunDnsProvider := NewAlidnsProvider(credential)
	if aliyunDnsProvider.Support(record.A) {
		request := models.NewDeleteDomainRecordRequest().
			SetID("738619282108471296")
		result, err := aliyunDnsProvider.DeleteRecord(request)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result.String())
	}

}
