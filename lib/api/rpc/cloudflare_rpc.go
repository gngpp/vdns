package rpc

import (
	"context"
	"fmt"
	"io/ioutil"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	conv2 "vdns/lib/api/rpc/conv"
	"vdns/lib/auth"
	"vdns/lib/util/iotool"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
)

func NewCloudflareRpc(credential auth.Credential) VdnsRpc {
	return &CloudflareRpc{
		conv:       &conv2.DNSPodResponseConvert{},
		credential: credential,
	}
}

type CloudflareRpc struct {
	conv       conv2.VdnsResponseConverter
	credential auth.Credential
}

func (_this *CloudflareRpc) DoDescribeRequest(url string) (*models.DomainRecordsResponse, error) {
	return _this.DoDescribeCtxRequest(nil, url)
}

func (_this *CloudflareRpc) DoCreateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoCreateCtxRequest(nil, url)
}

func (_this *CloudflareRpc) DoUpdateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoUpdateCtxRequest(nil, url)
}

func (_this *CloudflareRpc) DoDeleteRequest(url string) (*models.DomainRecordStatusResponse, error) {
	return _this.DoDeleteCtxRequest(nil, url)
}

func (_this *CloudflareRpc) DoDescribeCtxRequest(_ context.Context, url string) (*models.DomainRecordsResponse, error) {
	resp, err := vhttp.Get(url, strs.String(_this.credential.GetToken()))
	if err != nil {
		return nil, errs.NewVdnsFromError(err)
	}
	if vhttp.IsOK(resp) {
		body := resp.Body
		defer iotool.ReadCloser(body)
		all, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		fmt.Println(string(all))
	}
	return nil, nil
}

func (_this *CloudflareRpc) DoCreateCtxRequest(ctx context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareRpc) DoUpdateCtxRequest(ctx context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareRpc) DoDeleteCtxRequest(ctx context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}
