package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	conv2 "vdns/lib/api/rpc/conv"
	"vdns/lib/auth"
	"vdns/lib/util/iotool"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
	"vdns/lib/util/vjson"
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
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareRpc) DoUpdateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareRpc) DoDeleteRequest(url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareRpc) DoDescribeCtxRequest(ctx context.Context, url string) (*models.DomainRecordsResponse, error) {
	fmt.Println("url:" + url)
	resp, err := vhttp.Get(url, strs.String(_this.credential.GetToken()))
	if err != nil {
		return nil, errs.NewVdnsFromError(err)
	}
	body := resp.Body
	defer iotool.ReadCloser(body)
	all, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, errs.NewVdnsFromError(err)
	}
	fmt.Println(string(all))
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

// request 统一请求接口
func (_this *CloudflareRpc) request(method string, url string, data interface{}, result interface{}) (err error) {
	jsonStr := make([]byte, 0)
	if data != nil {
		jsonStr, _ = json.Marshal(data)
	}
	req, err := http.NewRequest(
		method,
		url,
		bytes.NewBuffer(jsonStr),
	)
	if err != nil {
		log.Println("http.NewRequest失败. Error: ", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+_this.credential.GetToken())
	req.Header.Set("Content-Type", "application/json")

	client := vhttp.NewClient()
	resp, err := client.Do(req)
	if vhttp.IsOK(resp) {
		body := resp.Body
		defer body.Close()
		all, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(vjson.PrettifyString(string(all)))
	} else {
		fmt.Println(err)
	}
	return
}
