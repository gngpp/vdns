package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zf1976/vdns/lib/api/models"
	"github.com/zf1976/vdns/lib/auth"
	"github.com/zf1976/vdns/lib/standard"
	"github.com/zf1976/vdns/lib/standard/record"
	"github.com/zf1976/vdns/lib/util/vhttp"
	"github.com/zf1976/vdns/lib/util/vjson"
	"io/ioutil"
	"log"
	"net/http"
)

func NewCloudflareProvider(credential auth.Credential) VdnsProvider {
	return &CloudflareProvider{
		api:        standard.CLOUDFLARE_DNS_API.String(),
		credential: credential,
	}
}

type CloudflareProvider struct {
	api        *standard.Standard
	credential auth.Credential
}

func (_this *CloudflareProvider) DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareProvider) CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareProvider) UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareProvider) DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (_this *CloudflareProvider) Support(recordType record.Type) error {
	//TODO implement me
	panic("implement me")
}

// 获得域名记录列表
func (_this *CloudflareProvider) getZones(domain *vhttp.Domain) (err error) {
	err = _this.request(
		"GET",
		fmt.Sprintf(_this.api.StringValue()+"?name=%s&status=%s&per_page=%s", domain.DomainName, "active", "50"),
		nil,
		nil,
	)

	return
}

// request 统一请求接口
func (_this *CloudflareProvider) request(method string, url string, data interface{}, result interface{}) (err error) {
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

	client := vhttp.CreateClient()
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
