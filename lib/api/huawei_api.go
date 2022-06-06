package api

import (
	"vdns/lib/api/model"
	"vdns/lib/auth"
	"vdns/lib/standard"
	"vdns/lib/standard/record"
)

func NewHuaweiProvider(credential auth.Credential) VdnsProvider {
	return &HuaweiProvider{
		api:        standard.HUAWEI_DNS_API.String(),
		credential: credential,
	}
}

type HuaweiProvider struct {
	api        *standard.Standard
	credential auth.Credential
}

func (h HuaweiProvider) DescribeRecords(request *model.DescribeDomainRecordsRequest) (*model.DomainRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h HuaweiProvider) CreateRecord(request *model.CreateDomainRecordRequest) (*model.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h HuaweiProvider) UpdateRecord(request *model.UpdateDomainRecordRequest) (*model.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h HuaweiProvider) DeleteRecord(request *model.DeleteDomainRecordRequest) (*model.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h HuaweiProvider) Support(recordType record.Type) error {
	//TODO implement me
	panic("implement me")
}
