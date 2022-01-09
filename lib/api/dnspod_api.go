package api

import (
	"context"
	"net/url"
	"vdns/lib/api/action"
	"vdns/lib/api/models"
	"vdns/lib/api/parameter"
	"vdns/lib/api/rpc"
	"vdns/lib/auth"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
	"vdns/lib/standard/record"
	"vdns/util/vhttp"
)

func NewDnspodProvider(credential auth.Credential) VdnsRecordProvider {
	signatureComposer := compose.NewDnspodSignatureCompose()
	return &DnspodProvider{
		RequestAction:     action.NewDnspodAction(),
		signatureComposer: signatureComposer,
		rpc:               rpc.NewDnspodRpc(),
		api:               standard.DNSPOD_DNS_API.String(),
		credential:        credential,
		parameterProvider: parameter.NewDnspodParameterProvider(credential, signatureComposer),
	}
}

type DnspodProvider struct {
	*action.RequestAction
	api               *standard.Standard
	signatureComposer compose.SignatureComposer
	parameterProvider parameter.ParamaterProvider
	credential        auth.Credential
	rpc               rpc.VdnsRpc
}

func (_this *DnspodProvider) DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordsResponse, error) {
	describeParamater, err := _this.parameterProvider.LoadDescribeParamater(request, _this.Describe)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(describeParamater)
	ctx := context.WithValue(context.Background(), parameter.DNSPOC_PARAMETER_CONTEXT_DESCRIBE_KEY, request)
	return _this.rpc.DoDescribeCtxRequest(ctx, requestUrl)
}

func (_this *DnspodProvider) CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.parameterProvider.LoadCreateParamater(request, _this.Create)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoCreateRequest(requestUrl)
}

func (_this *DnspodProvider) UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.parameterProvider.LoadUpdateParamater(request, _this.Update)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoUpdateRequest(requestUrl)
}

func (_this *DnspodProvider) DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.parameterProvider.LoadDeleteParamater(request, _this.Delete)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoDeleteRequest(requestUrl)
}

func (_this *DnspodProvider) Support(recordType record.Type) bool {
	return record.Support(recordType)
}

func (_this *DnspodProvider) generateRequestUrl(paramater *url.Values) string {
	stringToSign := _this.signatureComposer.ComposeStringToSign(vhttp.HttpMethodGet, paramater)
	signature := _this.signatureComposer.GeneratedSignature(_this.credential.GetSecretKey(), stringToSign)
	return _this.signatureComposer.CanonicalizeRequestUrl(_this.api.StringValue(), signature, paramater)
}
