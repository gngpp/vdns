package rpc

import (
	"vdns/lib/api/conv"
	"vdns/lib/api/models"
)

type DnspodRpc struct {
	conv conv.DomainRecordResponseConverter
}

func NewDnspodRpc() Rpc {
	return &AlidnsRpc{
		conv: &conv.AlidnsDomainRecordResponseConvert{},
	}
}

func (_this *DnspodRpc) DoDescribeRequest(url string) (*models.DomainRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *DnspodRpc) DoCreateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *DnspodRpc) DoUpdateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *DnspodRpc) DoDeleteRequest(url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}
