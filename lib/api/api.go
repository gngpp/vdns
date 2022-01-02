package api

import (
	"net/url"
	"vdns/lib/api/models"
	"vdns/lib/standard/record"
)

const DefaultRRKeyWord = "@"

type Action struct {
	create   *string
	update   *string
	describe *string
	delete   *string
}

type DNSRecordProvider interface {
	// DescribeRecords 具体参数作用请看实现注释
	DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordsResponse, error)

	// CreateRecord 具体参数作用请看实现注释
	CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainRecordStatusResponse, error)

	// UpdateRecord 具体参数作用请看实现注释
	UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainRecordStatusResponse, error)

	// DeleteRecord 具体参数作用请看实现注释
	DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainRecordStatusResponse, error)

	// Support 某些使用zone区域划分域名记录的DNS服务商，需强迫使用support
	Support(recordType record.Type) bool
}

type ParamaterProvider interface {
	// 加载Describe参数
	loadDescribeParamater(request *models.DescribeDomainRecordsRequest, action *string) (*url.Values, error)

	// 加载Create参数
	loadCreateParamater(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error)

	// 加载Update参数
	loadUpdateParamater(request *models.UpdateDomainRecordRequest, action *string) (*url.Values, error)

	// 加载Delete参数
	loadDeleteParamater(request *models.DeleteDomainRecordRequest, action *string) (*url.Values, error)
}
