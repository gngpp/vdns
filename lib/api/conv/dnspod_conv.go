package conv

import (
	"net/http"
	"vdns/lib/api/models"
)

type DnspodDomainRecordResponseConvert struct{}

func (_this *DnspodDomainRecordResponseConvert) DescribeResponseConvert(resp *http.Response) (*models.DomainRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *DnspodDomainRecordResponseConvert) CreateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *DnspodDomainRecordResponseConvert) UpdateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *DnspodDomainRecordResponseConvert) DeleteResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}
