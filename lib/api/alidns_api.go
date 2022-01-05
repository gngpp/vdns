package api

import (
	"net/url"
	"vdns/lib/api/models"
	"vdns/lib/api/parameter"
	"vdns/lib/api/rpc"
	"vdns/lib/auth"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
	"vdns/lib/standard/record"
	"vdns/vutil/strs"
	"vdns/vutil/vhttp"
)

func NewAlidnsProvider(credential auth.Credential) DNSRecordProvider {
	signatureComposer := compose.NewAlidnsSignatureCompose()
	return &AlidnsProvider{
		Action: &Action{
			describe: strs.String("DescribeDomainRecords"),
			create:   strs.String("AddDomainRecord"),
			update:   strs.String("UpdateDomainRecord"),
			delete:   strs.String("DeleteDomainRecord"),
		},
		signatureComposer: signatureComposer,
		rpc:               rpc.NewAlidnsRpc(),
		api:               standard.ALIYUN_DNS_API.String(),
		credential:        credential,
		parameterProvider: parameter.NewAlidnsParameterProvider(credential, signatureComposer),
	}
}

type AlidnsProvider struct {
	*Action
	api               *standard.Standard
	signatureComposer compose.SignatureComposer
	parameterProvider parameter.ParamaterProvider
	credential        auth.Credential
	rpc               rpc.Rpc
}

func (_this *AlidnsProvider) DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordsResponse, error) {
	paramater, err := _this.parameterProvider.LoadDescribeParamater(request, _this.describe)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoDescribeRequest(requestUrl)
}

func (_this *AlidnsProvider) CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.parameterProvider.LoadCreateParamater(request, _this.create)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoCreateRequest(requestUrl)
}

func (_this *AlidnsProvider) UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.parameterProvider.LoadUpdateParamater(request, _this.update)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoUpdateRequest(requestUrl)
}

func (_this *AlidnsProvider) DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.parameterProvider.LoadDeleteParamater(request, _this.delete)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoDeleteRequest(requestUrl)
}

func (_this *AlidnsProvider) Support(recordType record.Type) bool {
	return record.Support(recordType)
}

func (_this *AlidnsProvider) generateRequestUrl(paramater *url.Values) string {
	stringToSign := _this.signatureComposer.ComposeStringToSign(vhttp.HttpMethodGet, paramater)
	signature := _this.signatureComposer.GeneratedSignature(_this.credential.GetSecretKey(), stringToSign)
	return _this.signatureComposer.CanonicalizeRequestUrl(_this.api.StringValue(), signature, paramater)
}
