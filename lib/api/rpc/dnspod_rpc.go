package rpc

import (
	"context"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/api/rpc/conv"
	"vdns/lib/vlog"
)

func NewDNSPodRpc() VdnsRpc {
	return &DNSPodRpc{
		conv: &conv.DNSPodResponseConvert{},
	}
}

type DNSPodRpc struct {
	conv conv.VdnsResponseConverter
}

func (_this *DNSPodRpc) DoDescribeCtxRequest(ctx context.Context, url string) (*model.DomainRecordsResponse, error) {
	vlog.Debugf("[DNSPodRpc] request url: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return &model.DomainRecordsResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.DescribeResponseCtxConvert(ctx, resp)
}

func (_this *DNSPodRpc) DoCreateCtxRequest(_ context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	vlog.Debugf("[DNSPodRpc] request url: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return &model.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.CreateResponseConvert(resp)
}

func (_this *DNSPodRpc) DoUpdateCtxRequest(_ context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	vlog.Debugf("[DNSPodRpc] request url: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return &model.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.UpdateResponseConvert(resp)
}

func (_this *DNSPodRpc) DoDeleteCtxRequest(_ context.Context, url string) (*model.DomainRecordStatusResponse, error) {
	vlog.Debugf("[DNSPodRpc] request url: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return &model.DomainRecordStatusResponse{}, errs.NewVdnsFromError(err)
	}
	return _this.conv.DeleteResponseConvert(resp)
}

func (_this *DNSPodRpc) DoDescribeRequest(url string) (*model.DomainRecordsResponse, error) {
	return _this.DoDescribeCtxRequest(nil, url)
}

func (_this *DNSPodRpc) DoCreateRequest(url string) (*model.DomainRecordStatusResponse, error) {
	return _this.DoCreateCtxRequest(nil, url)
}

func (_this *DNSPodRpc) DoUpdateRequest(url string) (*model.DomainRecordStatusResponse, error) {
	return _this.DoUpdateCtxRequest(nil, url)
}

func (_this *DNSPodRpc) DoDeleteRequest(url string) (*model.DomainRecordStatusResponse, error) {
	return _this.DoDeleteCtxRequest(nil, url)
}
