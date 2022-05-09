package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"vdns/lib/api/conv"
	"vdns/lib/api/models"
	"vdns/lib/auth"
	"vdns/lib/util/vhttp"
	"vdns/lib/util/vjson"
)

func NewCloudflareRpc(credential auth.Credential) VdnsRpc {
	return &CloudflareRpc{
		conv:       &conv.DNSPodResponseConvert{},
		credential: credential,
	}
}

type CloudflareRpc struct {
	conv       conv.VdnsResponseConverter
	credential auth.Credential
}

func (c *CloudflareRpc) DoDescribeRequest(url string) (*models.DomainRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudflareRpc) DoCreateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudflareRpc) DoUpdateRequest(url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudflareRpc) DoDeleteRequest(url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudflareRpc) DoDescribeCtxRequest(ctx context.Context, url string) (*models.DomainRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudflareRpc) DoCreateCtxRequest(ctx context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudflareRpc) DoUpdateCtxRequest(ctx context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudflareRpc) DoDeleteCtxRequest(ctx context.Context, url string) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

// request 统一请求接口
func (c *CloudflareRpc) request(method string, url string, data interface{}, result interface{}) (err error) {
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
	req.Header.Set("Authorization", "Bearer "+c.credential.GetToken())
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
