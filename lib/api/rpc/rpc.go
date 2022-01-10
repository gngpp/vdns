package rpc

import (
	"context"
	"vdns/lib/api/models"
)

type VdnsRpc interface {
	DoDescribeRequest(url string) (*models.DomainRecordsResponse, error)

	DoCreateRequest(url string) (*models.DomainRecordStatusResponse, error)

	DoUpdateRequest(url string) (*models.DomainRecordStatusResponse, error)

	DoDeleteRequest(url string) (*models.DomainRecordStatusResponse, error)

	DoDescribeCtxRequest(ctx context.Context, url string) (*models.DomainRecordsResponse, error)

	DoCreateCtxRequest(ctx context.Context, url string) (*models.DomainRecordStatusResponse, error)

	DoUpdateCtxRequest(ctx context.Context, url string) (*models.DomainRecordStatusResponse, error)

	DoDeleteCtxRequest(ctx context.Context, url string) (*models.DomainRecordStatusResponse, error)
}
