package api

import (
	"net/url"
	"vdns/lib/api/action"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	"vdns/lib/api/parameter"
	"vdns/lib/api/rpc"
	"vdns/lib/auth"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
	"vdns/lib/standard/msg"
	"vdns/lib/standard/record"
	"vdns/lib/util/vhttp"
)

func NewAlidnsProvider(credential auth.Credential) VdnsProvider {
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

func (_this *AlidnsProvider) SetApi(api *standard.Standard) {
	_this.api = api
}

func (_this *AlidnsProvider) SetSignatureComposer(signatureComposer compose.SignatureComposer) {
	_this.signatureComposer = signatureComposer
}

func (_this *AlidnsProvider) SetParameterProvider(parameterProvider parameter.ParamaterProvider) {
	_this.parameterProvider = parameterProvider
}

func (_this *AlidnsProvider) SetCredential(credential auth.Credential) {
	_this.credential = credential
}

func (_this *AlidnsProvider) SetRpc(rpc rpc.VdnsRpc) {
	_this.rpc = rpc
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

func (_this *AlidnsProvider) Support(recordType record.Type) error {
	support := record.Support(recordType)
	if support {
		return nil
	}
	return errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
}

func (_this *AlidnsProvider) generateRequestUrl(paramater *url.Values) string {
	stringToSign := _this.signatureComposer.ComposeStringToSign(vhttp.HttpMethodGet, paramater)
	signature := _this.signatureComposer.GeneratedSignature(_this.credential.GetSecretKey(), stringToSign)
	return _this.signatureComposer.CanonicalizeRequestUrl(_this.api.StringValue(), signature, paramater)
}
