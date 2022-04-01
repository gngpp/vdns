package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"vdns/lib/standard"
	"vdns/lib/util/vhttp"
	"vdns/lib/util/vjson"
)

func TestName(t *testing.T) {

}

// 获得域名记录列表
func getZones(domain *vhttp.Domain) (err error) {
	err = request(
		"GET",
		fmt.Sprintf(standard.CLOUDFLARE_DNS_API.StringValue()+"?name=%s&status=%s&per_page=%s", domain.DomainName, "active", "50"),
		nil,
		nil,
	)
	return
}

// request 统一请求接口
func request(method string, url string, data interface{}, result interface{}) (err error) {
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
	req.Header.Set(vhttp.Authorization.String(), "Bearer "+"")
	req.Header.Set(vhttp.ContentType.String(), vhttp.ApplicationJson)

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
