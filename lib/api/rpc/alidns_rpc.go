package rpc

import (
	"context"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/api/rpc/conv"
	"vdns/lib/vlog"
)

func NewAliDNSRpc() VdnsRpc {
	return &AliDNSRpc{
		conv: &conv.AliDNSResponseConvert{},
	}
}

type AliDNSRpc struct {
	conv conv.VdnsResponseConverter
}

func (_this *AliDNSRpc) DoDescribeCtxRequest(_ context.Context, url string) (*model.DomainRecordsResponse, error) {
	vlog.Debugf("[AliDNSRpc] request url: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return &model.DomainRecordsResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.DescribeResponseConvert(resp)
}

func (_this *AliDNSRpc) DoCreateCtxRequest(_ context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	vlog.Debugf("[AliDNSRpc] request url: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return &model.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.CreateResponseConvert(resp)
}

func (_this *AliDNSRpc) DoUpdateCtxRequest(_ context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	vlog.Debugf("[AliDNSRpc] request url: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return &model.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.UpdateResponseConvert(resp)
}

func (_this *AliDNSRpc) DoDeleteCtxRequest(_ context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	vlog.Debugf("[AliDNSRpc] request url: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return &model.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.DeleteResponseConvert(resp)
}

func (_this *AliDNSRpc) DoDescribeRequest(url string) (*model.DomainRecordsResponse, error) {
	return _this.DoDescribeCtxRequest(nil, url)
}

func (_this *AliDNSRpc) DoCreateRequest(url string) (*model.DomainRecordStatusResponse, error) {
	return _this.DoCreateCtxRequest(nil, url)
}

func (_this *AliDNSRpc) DoUpdateRequest(url string) (*model.DomainRecordStatusResponse, error) {
	return _this.DoUpdateCtxRequest(nil, url)
}

func (_this *AliDNSRpc) DoDeleteRequest(url string) (*model.DomainRecordStatusResponse, error) {
	return _this.DoDeleteCtxRequest(nil, url)
}
