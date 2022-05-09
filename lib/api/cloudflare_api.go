package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	"vdns/lib/auth"
	"vdns/lib/standard"
	"vdns/lib/standard/msg"
	"vdns/lib/standard/record"
	"vdns/lib/util/iotool"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
	"vdns/lib/util/vjson"
	"vdns/lib/vlog"
)

func NewCloudflareProvider(credential auth.Credential) VdnsProvider {
	return &CloudflareProvider{
		api:                  standard.CLOUDFLARE_DNS_API.String(),
		credential:           credential,
		zonesMapUpdatedCount: 0,
		zonesMap:             make(map[string]string),
	}
}

type CloudflareProvider struct {
	api                  *standard.Standard
	zonesMapUpdatedCount int8
	zonesMap             map[string]string
	credential           auth.Credential
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
	// The zones may be updated, and the API will be initialized every 5 times.
	if len(_this.zonesMap) == 0 || (_this.zonesMapUpdatedCount > 5) {
		zonesMap, err := _this.getZones()
		if err != nil {
			return errs.NewVdnsFromError(err)
		}

		if _this.zonesMapUpdatedCount > 5 {
			_this.zonesMapUpdatedCount = 0
		}

		_this.zonesMapUpdatedCount += 1
		_this.zonesMap = zonesMap
	}

	support := record.Support(recordType)
	if support {
		return nil
	}
	return errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
}

// 获得域名区域列表
func (_this *CloudflareProvider) getZones() (map[string]string, error) {
	do, err := vhttp.Get(_this.api.StringValue(), strs.String(_this.credential.GetToken()))
	if err != nil {
		return nil, err
	}
	body := do.Body
	defer iotool.ReadCloser(body)
	all, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	zones := models.NewCloudflareZones()
	err = vjson.ByteArrayConvert(all, zones)
	if err != nil {
		return nil, err
	}
	zonesMap := make(map[string]string)
	if len(zones.Result) != 0 {
		for _, result := range zones.Result {
			zonesMap[result.Name] = result.Id
		}
		vlog.Debugf("cloudflare: zones ->\n %s", vjson.PrettifyString(zones))
	}
	return zonesMap, nil
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
