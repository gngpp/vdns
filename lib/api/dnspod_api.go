package api

import (
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	"vdns/lib/api/parameter"
	"vdns/lib/api/rpc"
	"vdns/lib/auth"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
	"vdns/lib/standard/msg"
	"vdns/lib/standard/record"
	"vdns/vutil/strs"
)

func NewDnspodProvider(credential auth.Credential) DNSRecordProvider {
	if credential != nil {
		panic(errs.NewApiError(msg.SYSTEM_CREDENTIAL_NOT_NIL))
	}
	signatureComposer := compose.NewDnspodSignatureCompose()
	return &DnspodProvider{
		Action: &Action{
			describe: strs.String("DescribeRecordList"),
			create:   strs.String("CreateRecord"),
			update:   strs.String("DeleteRecord"),
			delete:   strs.String("ModifyRecord"),
		},
		signatureComposer: signatureComposer,
		rpc:               rpc.NewDnspodRpc(),
		api:               standard.DNSPOD_DNS_API.String(),
		credential:        credential,
		parameterProvider: parameter.NewDnspodParameterProvider(credential, signatureComposer),
	}
}

type DnspodProvider struct {
	*Action
	api               *standard.Standard
	signatureComposer compose.SignatureComposer
	parameterProvider parameter.ParamaterProvider
	credential        auth.Credential
	rpc               rpc.Rpc
}

func (d DnspodProvider) DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d DnspodProvider) CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d DnspodProvider) UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d DnspodProvider) DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d DnspodProvider) Support(recordType record.Type) bool {
	//TODO implement me
	panic("implement me")
}
