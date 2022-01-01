package api

import (
	"net/url"
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

type DNSRecordProvider interface {
	// DescribeRecords 具体参数作用请看实现注释
	DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordResponse, error)

	// CreateRecord 具体参数作用请看实现注释
	CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainStatusResponse, error)

	// UpdateRecord 具体参数作用请看实现注释
	UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainStatusResponse, error)

	// DeleteRecord 具体参数作用请看实现注释
	DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainStatusResponse, error)

	// Support 某些使用zone区域划分域名记录的DNS服务商，需强迫使用support
	Support(recordType record.Type) bool
}

type DNSRequest interface {
	doDescribeRequest(url string) (*models.DomainRecordResponse, error)

	doCreateRequest(url string) (*models.DomainStatusResponse, error)

	doUpdateRequest(url string) (*models.DomainStatusResponse, error)

	doDeleteRequest(url string) (*models.DomainStatusResponse, error)

	loadDescribeParamater(request *models.DescribeDomainRecordsRequest, action *string) (*url.Values, error)

	loadCreateParamater(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error)

	loadUpdateParamater(request *models.UpdateDomainRecordRequest, action *string) (*url.Values, error)

	loadDeleteParamater(request *models.DeleteDomainRecordRequest, action *string) (*url.Values, error)
}
