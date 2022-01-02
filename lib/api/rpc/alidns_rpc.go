package rpc

import (
	"net/http"
	"vdns/lib/api/conv"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
)

type AlidnsRpc struct {
	conv conv.DomainRecordResponseConverter
}

func NewAlidnsRpc() Rpc {
	return &AlidnsRpc{
		conv: &conv.AlidnsDomainRecordResponseConvert{},
	}
}

func (_this *AlidnsRpc) DoDescribeRequest(url string) (*models.DomainRecordsResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordsResponse{}, errs.NewApiErrorFromError(err)
	}
	return _this.conv.DescribeResponseConvert(resp)
}

func (_this *AlidnsRpc) DoCreateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewApiErrorFromError(err)
	}
	return _this.conv.UpdateResponseConvert(resp)
}

func (_this *AlidnsRpc) DoUpdateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewApiErrorFromError(err)
	}
	return _this.conv.UpdateResponseConvert(resp)
}

func (_this *AlidnsRpc) DoDeleteRequest(url string) (*models.DomainRecordStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordStatusResponse{}, errs.NewApiErrorFromError(err)
	}
	return _this.conv.DeleteResponseConvert(resp)
}
