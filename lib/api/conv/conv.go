package conv

import (
	"net/http"
	"vdns/lib/api/models"
)

type DomainRecordResponseConverter interface {
	DescribeResponseConvert(resp *http.Response) (*models.DomainRecordsResponse, error)

	CreateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error)

	UpdateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error)

	DeleteResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error)
}
