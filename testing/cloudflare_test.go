package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"vdns/lib/api"
	"vdns/lib/api/models"
	"vdns/lib/auth"
	"vdns/lib/standard/record"
	"vdns/lib/util/iotool"
	"vdns/lib/util/vhttp"
	"vdns/lib/util/vjson"
)

func TestName(t *testing.T) {
	//vlog.SetLevel(vlog.Level.DEBUG)
	credential := auth.NewTokenCredential("shww4JpWY1Ilp43DHDMwY8ja_aoPs-RSJwmTcobi")
	provider := api.NewCloudflareProvider(credential)
	err := provider.Support(record.A)
	if err != nil {
		fmt.Println(err)
	}
	recordsRequest := models.NewDescribeDomainRecordsRequest()
	recordsRequest.SetDomain("innas.work")
	recordsRequest.SetRecordType(record.A)
	records, err := provider.DescribeRecords(recordsRequest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vjson.PrettifyString(records))

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

	client := vhttp.NewClient()
	resp, err := client.Do(req)
	if vhttp.IsOK(resp) {
		body := resp.Body
		defer iotool.ReadCloser(body)
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
