package rpc

import (
	"context"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	conv2 "vdns/lib/api/rpc/conv"
)

func NewAliDNSRpc() VdnsRpc {
	return &AliDNSRpc{
		conv: &conv2.AliDNSResponseConvert{},
	}
}

type AliDNSRpc struct {
	conv conv2.VdnsResponseConverter
}

func (_this *AliDNSRpc) DoDescribeCtxRequest(_ context.Context, url string) (*models.DomainRecordsResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordsResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.DescribeResponseConvert(resp)
}

func (_this *AliDNSRpc) DoCreateCtxRequest(_ context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.CreateResponseConvert(resp)
}

func (_this *AliDNSRpc) DoUpdateCtxRequest(_ context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.UpdateResponseConvert(resp)
}

func (_this *AliDNSRpc) DoDeleteCtxRequest(_ context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.DeleteResponseConvert(resp)
}

func (_this *AliDNSRpc) DoDescribeRequest(url string) (*models.DomainRecordsResponse, error) {
	return _this.DoDescribeCtxRequest(nil, url)
}

func (_this *AliDNSRpc) DoCreateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoCreateCtxRequest(nil, url)
}

func (_this *AliDNSRpc) DoUpdateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoUpdateCtxRequest(nil, url)
}

func (_this *AliDNSRpc) DoDeleteRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoDeleteCtxRequest(nil, url)
}
