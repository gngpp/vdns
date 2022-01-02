package rpc

import "vdns/lib/api/models"

type Rpc interface {
	DoDescribeRequest(url string) (*models.DomainRecordsResponse, error)

	DoCreateRequest(url string) (*models.DomainRecordStatusResponse, error)

	DoUpdateRequest(url string) (*models.DomainRecordStatusResponse, error)

	DoDeleteRequest(url string) (*models.DomainRecordStatusResponse, error)
}
