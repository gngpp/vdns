package rpc

import (
	"context"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/api/parameter"
	"vdns/lib/api/rpc/conv"
	"vdns/lib/auth"
	"vdns/lib/util/vhttp"
)

func NewCloudflareRpc(credential auth.Credential) VdnsRpc {
	return &CloudflareRpc{
		conv:       &conv.CloudflareResponseConvert{},
		credential: credential,
	}
}

type CloudflareRpc struct {
	conv       conv.VdnsResponseConverter
	credential auth.Credential
}

func (_this *CloudflareRpc) DoDescribeRequest(url string) (*model.DomainRecordsResponse, error) {
	return _this.DoDescribeCtxRequest(nil, url)
}

func (_this *CloudflareRpc) DoCreateRequest(url string) (*model.DomainRecordStatusResponse, error) {
	return _this.DoCreateCtxRequest(nil, url)
}

func (_this *CloudflareRpc) DoUpdateRequest(url string) (*model.DomainRecordStatusResponse, error) {
	return _this.DoUpdateCtxRequest(nil, url)
}

func (_this *CloudflareRpc) DoDeleteRequest(url string) (*model.DomainRecordStatusResponse, error) {
	return _this.DoDeleteCtxRequest(nil, url)
}

func (_this *CloudflareRpc) DoDescribeCtxRequest(_ context.Context, url string) (*model.DomainRecordsResponse, error) {
	resp, err := vhttp.Get(url, _this.credential.GetToken())
	if err != nil {
		return &model.DomainRecordsResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.DescribeResponseConvert(resp)
}

func (_this *CloudflareRpc) DoCreateCtxRequest(ctx context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	data := ctx.Value(parameter.CfParameterContextCreateKey)
	resp, err := vhttp.Post(url, "application/json", data, _this.credential.GetToken())
	if err != nil {
		return &model.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.CreateResponseConvert(resp)
}

func (_this *CloudflareRpc) DoUpdateCtxRequest(ctx context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	data := ctx.Value(parameter.CfParameterContextUpdateKey)
	resp, err := vhttp.Put(url, "application/json", data, _this.credential.GetToken())
	if err != nil {
		return &model.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.UpdateResponseConvert(resp)
}

func (_this *CloudflareRpc) DoDeleteCtxRequest(_ context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	resp, err := vhttp.Delete(url, "application/json", nil, _this.credential.GetToken())
	if err != nil {
		return &model.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.UpdateResponseConvert(resp)
}
