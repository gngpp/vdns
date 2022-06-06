package api

import (
	"net/url"
	"vdns/lib/api/action"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/api/parameter"
	"vdns/lib/api/rpc"
	"vdns/lib/auth"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
	"vdns/lib/standard/msg"
	"vdns/lib/standard/record"
	"vdns/lib/util/vhttp"
)

func NewAliDNSProvider(credential auth.Credential) VdnsProvider {
	signatureComposer := compose.NewAliDNSSignatureCompose()
	return &AliDNSProvider{
		RequestAction:     action.NewAliDNSAction(),
		signatureComposer: signatureComposer,
		rpc:               rpc.NewAliDNSRpc(),
		api:               standard.ALIYUN_DNS_API.String(),
		credential:        credential,
		parameter:         parameter.NewAliDNSParameter(credential, signatureComposer),
	}
}

type AliDNSProvider struct {
	*action.RequestAction
	api               *standard.Standard
	signatureComposer compose.SignatureComposer
	parameter         parameter.Parameter
	credential        auth.Credential
	rpc               rpc.VdnsRpc
}

func (_this *AliDNSProvider) DescribeRecords(request *model.DescribeDomainRecordsRequest) (*model.DomainRecordsResponse, error) {
	p, err := _this.parameter.LoadDescribeParameter(request, _this.Describe)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(p.UrlValues)
	return _this.rpc.DoDescribeRequest(requestUrl)
}

func (_this *AliDNSProvider) CreateRecord(request *model.CreateDomainRecordRequest) (*model.DomainRecordStatusResponse, error) {
	p, err := _this.parameter.LoadCreateParameter(request, _this.Create)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(p.UrlValues)
	return _this.rpc.DoCreateRequest(requestUrl)
}

func (_this *AliDNSProvider) UpdateRecord(request *model.UpdateDomainRecordRequest) (*model.DomainRecordStatusResponse, error) {
	p, err := _this.parameter.LoadUpdateParameter(request, _this.Update)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(p.UrlValues)
	return _this.rpc.DoUpdateRequest(requestUrl)
}

func (_this *AliDNSProvider) DeleteRecord(request *model.DeleteDomainRecordRequest) (*model.DomainRecordStatusResponse, error) {
	p, err := _this.parameter.LoadDeleteParameter(request, _this.Delete)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(p.UrlValues)
	return _this.rpc.DoDeleteRequest(requestUrl)
}

func (_this *AliDNSProvider) Support(recordType record.Type) error {
	support := record.Support(recordType)
	if support {
		return nil
	}
	return errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
}

func (_this *AliDNSProvider) generateRequestUrl(parameter *url.Values) string {
	stringToSign := _this.signatureComposer.ComposeStringToSign(vhttp.HttpMethodGet, parameter)
	signature := _this.signatureComposer.GeneratedSignature(_this.credential.GetSecretKey(), stringToSign)
	return _this.signatureComposer.CanonicalizeRequestUrl(_this.api.StringValue(), signature, parameter)
}

func (_this *AliDNSProvider) SetApi(api *standard.Standard) {
	_this.api = api
}

func (_this *AliDNSProvider) SetSignatureComposer(signatureComposer compose.SignatureComposer) {
	_this.signatureComposer = signatureComposer
}

func (_this *AliDNSProvider) SetParameterProvider(parameterProvider parameter.Parameter) {
	_this.parameter = parameterProvider
}

func (_this *AliDNSProvider) SetCredential(credential auth.Credential) {
	_this.credential = credential
}

func (_this *AliDNSProvider) SetRpc(rpc rpc.VdnsRpc) {
	_this.rpc = rpc
}
