package rpc

import (
	"context"
	"vdns/lib/api/model"
)

//goland:noinspection SpellCheckingInspection
type VdnsRpc interface {
	DoDescribeRequest(url string) (*model.DomainRecordsResponse, error)

	DoCreateRequest(url string) (*model.DomainRecordStatusResponse, error)

	DoUpdateRequest(url string) (*model.DomainRecordStatusResponse, error)

	DoDeleteRequest(url string) (*model.DomainRecordStatusResponse, error)

	DoDescribeCtxRequest(ctx context.Context, url string) (*model.DomainRecordsResponse, error)

	DoCreateCtxRequest(ctx context.Context, url string) (*model.DomainRecordStatusResponse, error)

	DoUpdateCtxRequest(ctx context.Context, url string) (*model.DomainRecordStatusResponse, error)

	DoDeleteCtxRequest(ctx context.Context, url string) (*model.DomainRecordStatusResponse, error)
}
