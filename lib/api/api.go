package api

import (
	"vdns/lib/api/models"
	"vdns/lib/standard/record"
)

const DefaultRRKeyWord = "@"

type Action struct {
	create   string
	update   string
	describe string
	delete   string
}

type DnsProvider interface {
	// DescribeRecords 具体参数作用请看实现注释
	DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.ApiDomainRecordResponse, error)

	// CreateDnsRecord 具体参数作用请看实现注释
	CreateDnsRecord(request *models.CreateDomainRecordRequest) (*models.ApiDomainRecordResponse, error)

	// UpdateDnsRecord 具体参数作用请看实现注释
	UpdateDnsRecord(request *models.UpdateDomainRecordRequest) (*models.ApiDomainRecordResponse, error)

	// DeleteDnsRecord 具体参数作用请看实现注释
	DeleteDnsRecord(request *models.DeleteDomainRecordRequest) (*models.ApiDomainRecordResponse, error)

	// Support 某些使用zone区域划分域名记录的DNS服务商，需强迫使用support
	Support(recordType record.Type) bool
}
