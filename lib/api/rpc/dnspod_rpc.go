package rpc

import (
	"context"
	"net/http"
	"vdns/lib/api/conv"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
)

func NewDNSPodRpc() VdnsRpc {
	return &DNSPodRpc{
		conv: &conv.DNSPodResponseConvert{},
	}
}

type DNSPodRpc struct {
	conv conv.VdnsResponseConverter
}

func (_this *DNSPodRpc) DoDescribeCtxRequest(ctx context.Context, url string) (*models.DomainRecordsResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordsResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.DescribeResponseCtxConvert(ctx, resp)
}

func (_this *DNSPodRpc) DoCreateCtxRequest(_ context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.CreateResponseConvert(resp)
}

func (_this *DNSPodRpc) DoUpdateCtxRequest(_ context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.UpdateResponseConvert(resp)
}

func (_this *DNSPodRpc) DoDeleteCtxRequest(_ context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.DeleteResponseConvert(resp)
}

func (_this *DNSPodRpc) DoDescribeRequest(url string) (*models.DomainRecordsResponse, error) {
	return _this.DoDescribeCtxRequest(nil, url)
}

func (_this *DNSPodRpc) DoCreateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoCreateCtxRequest(nil, url)
}

func (_this *DNSPodRpc) DoUpdateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoUpdateCtxRequest(nil, url)
}

func (_this *DNSPodRpc) DoDeleteRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoDeleteCtxRequest(nil, url)
}
