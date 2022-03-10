package api

import (
	"vdns/lib/api/models"
	"vdns/lib/auth"
	"vdns/lib/standard"
	"vdns/lib/standard/record"
)

func NewHuaweiProvider(credential auth.Credential) VDNSProvider {
	return &HuaweiProvider{
		api:        standard.HUAWEI_DNS_API.String(),
		credential: credential,
	}
}

type HuaweiProvider struct {
	api        *standard.Standard
	credential auth.Credential
}

func (h HuaweiProvider) DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h HuaweiProvider) CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h HuaweiProvider) UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h HuaweiProvider) DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h HuaweiProvider) Support(recordType record.Type) error {
	//TODO implement me
	panic("implement me")
}
