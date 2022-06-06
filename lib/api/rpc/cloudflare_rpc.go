package rpc

import (
	"context"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/api/rpc/conv"
	"vdns/lib/auth"
	"vdns/lib/util/strs"
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
	resp, err := vhttp.Get(url, strs.String(_this.credential.GetToken()))
	if err != nil {
		return &model.DomainRecordsResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.DescribeResponseConvert(resp)
}

func (_this *CloudflareRpc) DoCreateCtxRequest(ctx context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareRpc) DoUpdateCtxRequest(ctx context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareRpc) DoDeleteCtxRequest(ctx context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}
