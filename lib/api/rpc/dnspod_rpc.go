package rpc

import (
	"context"
	"github.com/zf1976/vdns/lib/api/conv"
	"github.com/zf1976/vdns/lib/api/errs"
	"github.com/zf1976/vdns/lib/api/models"
	"net/http"
)

func NewDnspodRpc() VdnsRpc {
	return &DnspodRpc{
		conv: &conv.DnspodResponseConvert{},
	}
}

type DnspodRpc struct {
	conv conv.VdnsResponseConverter
}

func (_this *DnspodRpc) DoDescribeCtxRequest(ctx context.Context, url string) (*models.DomainRecordsResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordsResponse{}, errs.NewApiErrorFromError(err)
	}
	return _this.conv.DescribeResponseCtxConvert(ctx, resp)
}

func (_this *DnspodRpc) DoCreateCtxRequest(_ context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewApiErrorFromError(err)
	}
	return _this.conv.CreateResponseConvert(resp)
}

func (_this *DnspodRpc) DoUpdateCtxRequest(_ context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewApiErrorFromError(err)
	}
	return _this.conv.UpdateResponseConvert(resp)
}

func (_this *DnspodRpc) DoDeleteCtxRequest(_ context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewApiErrorFromError(err)
	}
	return _this.conv.DeleteResponseConvert(resp)
}

func (_this *DnspodRpc) DoDescribeRequest(url string) (*models.DomainRecordsResponse, error) {
	return _this.DoDescribeCtxRequest(nil, url)
}

func (_this *DnspodRpc) DoCreateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoCreateCtxRequest(nil, url)
}

func (_this *DnspodRpc) DoUpdateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoUpdateCtxRequest(nil, url)
}

func (_this *DnspodRpc) DoDeleteRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoDeleteCtxRequest(nil, url)
}
