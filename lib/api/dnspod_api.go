package api

import (
	"context"
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

//goland:noinspection ALL
func NewDNSPodProvider(credential auth.Credential) VdnsProvider {
	signatureComposer := compose.NewDnspodSignatureCompose()
	return &DNSPodProvider{
		RequestAction:     action.NewDNSPodAction(),
		signatureComposer: signatureComposer,
		rpc:               rpc.NewDNSPodRpc(),
		api:               standard.DNSPOD_DNS_API.String(),
		credential:        credential,
		parameterProvider: parameter.NewDNSPodParameterProvider(credential, signatureComposer),
	}
}

type DNSPodProvider struct {
	*action.RequestAction
	api               *standard.Standard
	signatureComposer compose.SignatureComposer
	parameterProvider parameter.ParamaterProvider
	credential        auth.Credential
	rpc               rpc.VdnsRpc
}

func (_this *DNSPodProvider) DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordsResponse, error) {
	p, err := _this.parameterProvider.LoadDescribeParamater(request, _this.Describe)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(p)
	ctx := context.WithValue(context.Background(), parameter.DNSPOC_PARAMETER_CONTEXT_DESCRIBE_KEY, request)
	return _this.rpc.DoDescribeCtxRequest(ctx, requestUrl)
}

func (_this *DNSPodProvider) CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	p, err := _this.parameterProvider.LoadCreateParamater(request, _this.Create)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(p)
	return _this.rpc.DoCreateRequest(requestUrl)
}

func (_this *DNSPodProvider) UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	p, err := _this.parameterProvider.LoadUpdateParamater(request, _this.Update)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(p)
	return _this.rpc.DoUpdateRequest(requestUrl)
}

func (_this *DNSPodProvider) DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	p, err := _this.parameterProvider.LoadDeleteParamater(request, _this.Delete)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(p)
	return _this.rpc.DoDeleteRequest(requestUrl)
}

func (_this *DNSPodProvider) Support(recordType record.Type) error {
	support := record.Support(recordType)
	if support {
		return nil
	}
	return errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
}

func (_this *DNSPodProvider) generateRequestUrl(parameter *url.Values) string {
	stringToSign := _this.signatureComposer.ComposeStringToSign(vhttp.HttpMethodGet, parameter)
	signature := _this.signatureComposer.GeneratedSignature(_this.credential.GetSecretKey(), stringToSign)
	return _this.signatureComposer.CanonicalizeRequestUrl(_this.api.StringValue(), signature, parameter)
}

func (_this *DNSPodProvider) SetApi(api *standard.Standard) {
	_this.api = api
}

func (_this *DNSPodProvider) SetSignatureComposer(signatureComposer compose.SignatureComposer) {
	_this.signatureComposer = signatureComposer
}

func (_this *DNSPodProvider) SetParameterProvider(parameterProvider parameter.ParamaterProvider) {
	_this.parameterProvider = parameterProvider
}

func (_this *DNSPodProvider) SetCredential(credential auth.Credential) {
	_this.credential = credential
}

func (_this *DNSPodProvider) SetRpc(rpc rpc.VdnsRpc) {
	_this.rpc = rpc
}
