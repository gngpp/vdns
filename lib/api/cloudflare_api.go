package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/url"
	"vdns/lib/api/action"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/api/model/cloudflare_model"
	"vdns/lib/api/parameter"
	"vdns/lib/api/rpc"
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
		RequestAction:        action.NewCloudflareAction(),
		api:                  standard.CLOUDFLARE_DNS_API.String(),
		credential:           credential,
		zonesMapUpdatedCount: 0,
		zonesMap:             make(map[string]string),
		parameter:            parameter.NewCloudflareParameter(),
		rpc:                  rpc.NewCloudflareRpc(credential),
	}
}

type CloudflareProvider struct {
	*action.RequestAction
	api                  *standard.Standard
	zonesMapUpdatedCount int8
	zonesMap             map[string]string
	credential           auth.Credential
	parameter            parameter.Parameter
	rpc                  rpc.VdnsRpc
}

func (_this *CloudflareProvider) DescribeRecords(request *model.DescribeDomainRecordsRequest) (*model.DomainRecordsResponse, error) {
	p, err := _this.parameter.LoadDescribeParameter(request, _this.Describe)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(nil, model.Domain(request), p.UrlValues)
	return _this.rpc.DoDescribeRequest(requestUrl)
}

func (_this *CloudflareProvider) CreateRecord(request *model.CreateDomainRecordRequest) (*model.DomainRecordStatusResponse, error) {
	p, err := _this.parameter.LoadCreateParameter(request, _this.Create)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(nil, model.Domain(request), nil)
	ctx := context.WithValue(context.Background(), parameter.CfParameterContextCreateKey, p.JsonString)
	return _this.rpc.DoCreateCtxRequest(ctx, requestUrl)
}

func (_this *CloudflareProvider) UpdateRecord(request *model.UpdateDomainRecordRequest) (*model.DomainRecordStatusResponse, error) {
	p, err := _this.parameter.LoadUpdateParameter(request, _this.Update)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(request.ID, model.Domain(request), nil)
	ctx := context.WithValue(context.Background(), parameter.CfParameterContextUpdateKey, p.JsonString)
	return _this.rpc.DoUpdateCtxRequest(ctx, requestUrl)
}

func (_this *CloudflareProvider) DeleteRecord(request *model.DeleteDomainRecordRequest) (*model.DomainRecordStatusResponse, error) {
	_, _ = _this.parameter.LoadDeleteParameter(request, _this.Update)
	requestUrl := _this.generateRequestUrl(request.ID, model.Domain(request), nil)
	return _this.rpc.DoDeleteRequest(requestUrl)
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
	do, err := vhttp.Get(_this.api.StringValue(), _this.credential.GetToken())
	if err != nil {
		return nil, err
	}
	body := do.Body
	defer iotool.ReadCloser(body)
	all, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	zones := cloudflare_model.NewCloudflareZones()
	err = vjson.ByteArrayConvert(all, zones)
	if err != nil {
		return nil, err
	}
	if len(zones.Errors) > 0 {
		return nil, errs.NewCloudFlareSDKError(vjson.PrettifyString(zones.Errors))
	}
	zonesMap := make(map[string]string)
	if len(zones.Result) != 0 {
		vlog.Debugf("cloudflare zones ->\n%s", vjson.PrettifyString(zones))
		for _, result := range zones.Result {
			zonesMap[result.Name] = result.Id
		}
		vlog.Debugf("cloudflare zones map ->\n%s", vjson.PrettifyString(zonesMap))
	}
	return zonesMap, nil
}

func (_this *CloudflareProvider) getZoneUrl(domain string) (string, error) {
	var identifier string
	if id, ok := _this.zonesMap[domain]; ok {
		if strs.IsEmpty(id) {
			return "", errs.NewVdnsError("Resolved primary domain name:" + domain + " does not exist")
		}
		identifier = id
	}
	return strs.Concat(_this.api.StringValue(), "/", identifier, "/dns_records"), nil
}

func (_this *CloudflareProvider) generateRequestUrl(identifier *string, domain model.Domain, parameter *url.Values) string {
	queryString := _this.toCanonicalizeStringQueryString(parameter)
	zoneUrl, err := _this.getZoneUrl(domain.GetDomain())
	if err != nil {
		vlog.Error(err)
		return ""
	}
	if strs.IsEmpty(queryString) {
		if identifier == nil || strs.IsEmpty(*identifier) {
			return zoneUrl
		} else {
			return strs.Concat(zoneUrl, "/", *identifier, "?")
		}
	} else {
		if identifier == nil || strs.IsEmpty(*identifier) {
			return strs.Concat(zoneUrl, "?", queryString)
		} else {
			return strs.Concat(zoneUrl, "/", *identifier, "?", queryString)
		}
	}
}

func (_this *CloudflareProvider) toCanonicalizeStringQueryString(parameter *url.Values) string {
	if parameter == nil {
		return ""
	}
	buf := new(bytes.Buffer)
	// sort keys by ascii asc order
	keys := make([]string, 0, len(*parameter))
	for k := range *parameter {
		keys = append(keys, k)
	}

	for i := range keys {
		k := keys[i]
		v := parameter.Get(k)
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
		buf.WriteString("&")
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}
