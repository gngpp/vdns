package api

import (
	"context"
	"github.com/zf1976/vdns/lib/api/action"
	"github.com/zf1976/vdns/lib/api/errs"
	"github.com/zf1976/vdns/lib/api/models"
	"github.com/zf1976/vdns/lib/api/parameter"
	"github.com/zf1976/vdns/lib/api/rpc"
	"github.com/zf1976/vdns/lib/auth"
	"github.com/zf1976/vdns/lib/sign/compose"
	"github.com/zf1976/vdns/lib/standard"
	"github.com/zf1976/vdns/lib/standard/msg"
	"github.com/zf1976/vdns/lib/standard/record"
	"github.com/zf1976/vdns/lib/util/vhttp"
	"net/url"
)

func NewDnspodProvider(credential auth.Credential) VdnsProvider {
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

func (_this *DnspodProvider) SetApi(api *standard.Standard) {
	_this.api = api
}

func (_this *DnspodProvider) SetSignatureComposer(signatureComposer compose.SignatureComposer) {
	_this.signatureComposer = signatureComposer
}

func (_this *DnspodProvider) SetParameterProvider(parameterProvider parameter.ParamaterProvider) {
	_this.parameterProvider = parameterProvider
}

func (_this *DnspodProvider) SetCredential(credential auth.Credential) {
	_this.credential = credential
}

func (_this *DnspodProvider) SetRpc(rpc rpc.VdnsRpc) {
	_this.rpc = rpc
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

func (_this *DnspodProvider) Support(recordType record.Type) error {
	support := record.Support(recordType)
	if support {
		return nil
	}
	return errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
}

func (_this *DnspodProvider) generateRequestUrl(paramater *url.Values) string {
	stringToSign := _this.signatureComposer.ComposeStringToSign(vhttp.HttpMethodGet, paramater)
	signature := _this.signatureComposer.GeneratedSignature(_this.credential.GetSecretKey(), stringToSign)
	return _this.signatureComposer.CanonicalizeRequestUrl(_this.api.StringValue(), signature, paramater)
}
