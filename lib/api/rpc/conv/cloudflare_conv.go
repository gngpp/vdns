package conv

import (
	"context"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	"vdns/lib/util/vhttp"
)

type CloudflareResponseConvert struct {
}

func (c CloudflareResponseConvert) DescribeResponseConvert(resp *http.Response) (*models.DomainRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudflareResponseConvert) CreateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	if vhttp.IsOK(resp) {

	} else {
		return nil, errs.NewVdnsError("")
	}
	//TODO implement me
	panic("implement me")
}

func (c CloudflareResponseConvert) UpdateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudflareResponseConvert) DeleteResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudflareResponseConvert) DescribeResponseCtxConvert(ctx context.Context, resp *http.Response) (*models.DomainRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudflareResponseConvert) CreateResponseCtxConvert(ctx context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudflareResponseConvert) UpdateResponseCtxConvert(ctx context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudflareResponseConvert) DeleteResponseCtxConvert(ctx context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}
