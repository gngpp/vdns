package api

import (
	"vdns/lib/api/models"
	"vdns/lib/standard/record"
)

type VdnsRecordProvider interface {
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
