package conv

import (
	"context"
	"net/http"
	"vdns/lib/api/models"
)

type VdnsResponseConverter interface {
	DescribeResponseConvert(resp *http.Response) (*models.DomainRecordsResponse, error)

	CreateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error)

	UpdateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error)

	DeleteResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error)

	DescribeResponseCtxConvert(ctx context.Context, resp *http.Response) (*models.DomainRecordsResponse, error)

	CreateResponseCtxConvert(ctx context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error)

	UpdateResponseCtxConvert(ctx context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error)

	DeleteResponseCtxConvert(ctx context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error)
}
