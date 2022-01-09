package api

import (
	"net/url"
	"vdns/lib/api/action"
	"vdns/lib/api/models"
	"vdns/lib/api/parameter"
	"vdns/lib/api/rpc"
	"vdns/lib/auth"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
	"vdns/lib/standard/record"
	"vdns/vutil/vhttp"
)

func NewAlidnsProvider(credential auth.Credential) VdnsRecordProvider {
	signatureComposer := compose.NewAlidnsSignatureCompose()
	return &AlidnsProvider{
		RequestAction:     action.NewAlidnsAction(),
		signatureComposer: signatureComposer,
		rpc:               rpc.NewAlidnsRpc(),
		api:               standard.ALIYUN_DNS_API.String(),
		credential:        credential,
		parameterProvider: parameter.NewAlidnsParameterProvider(credential, signatureComposer),
	}
}

type AlidnsProvider struct {
	*action.RequestAction
	api               *standard.Standard
	signatureComposer compose.SignatureComposer
	parameterProvider parameter.ParamaterProvider
	credential        auth.Credential
	rpc               rpc.VdnsRpc
}

func (_this *AlidnsProvider) DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordsResponse, error) {
	paramater, err := _this.parameterProvider.LoadDescribeParamater(request, _this.Describe)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoDescribeRequest(requestUrl)
}

func (_this *AlidnsProvider) CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.parameterProvider.LoadCreateParamater(request, _this.Create)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoCreateRequest(requestUrl)
}

func (_this *AlidnsProvider) UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.parameterProvider.LoadUpdateParamater(request, _this.Update)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoUpdateRequest(requestUrl)
}

func (_this *AlidnsProvider) DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.parameterProvider.LoadDeleteParamater(request, _this.Delete)
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
