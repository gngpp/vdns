package conv

import (
	"context"
	"net/http"
	"vdns/lib/api/model"
)

type VdnsResponseConverter interface {
	DescribeResponseConvert(resp *http.Response) (*model.DomainRecordsResponse, error)

	CreateResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error)

	UpdateResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error)

	DeleteResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error)

	DescribeResponseCtxConvert(ctx context.Context, resp *http.Response) (*model.DomainRecordsResponse, error)

	CreateResponseCtxConvert(ctx context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error)

	UpdateResponseCtxConvert(ctx context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error)

	DeleteResponseCtxConvert(ctx context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error)
}
