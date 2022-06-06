package parameter

import (
	"net/url"
	"strconv"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/standard"
	"vdns/lib/standard/msg"
	"vdns/lib/standard/record"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
)

func NewCloudflareParameter() Parameter {
	return &CloudflareParameter{
		version: standard.DNSPOD_API_VERSION.String(),
	}
}

type CloudflareParameter struct {
	version *standard.Standard
}

func (_this *CloudflareParameter) LoadDescribeParameter(request *model.DescribeDomainRecordsRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.CREATE_REQUEST_NOT_NIL)
	}

	// assert record type
	if !record.Support(request.RecordType) {
		return nil, errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
	}

	parameter := _this.loadCommonParameter(request.RecordType)

	// assert page number
	if request.PageNumber != nil {
		parameter.Set(CfParameterPage, strconv.FormatInt(*request.PageNumber, 10))
	}

	// assert page size
	if request.PageSize != nil {
		parameter.Set(CfParameterPerPage, strconv.FormatInt(*request.PageSize, 10))
	}

	// record value content
	if request.ValueKeyWord != nil {
		parameter.Set(CfParameterContent, *request.ValueKeyWord)
	}

	return parameter, nil
}

func (_this *CloudflareParameter) LoadCreateParameter(request *model.CreateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.CREATE_REQUEST_NOT_NIL)
	}
	parameter := _this.loadCommonParameter(request.RecordType)
	// assert domain
	domain, err := vhttp.CheckExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewVdnsFromError(err)
	}

	// assert record type
	if !record.Support(request.RecordType) {
		return nil, errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
	}

	// assert value
	if request.Value == nil {
		return nil, errs.NewVdnsError(msg.RECORD_VALUE_NOT_SUPPORT)
	}

	parameter.Set("name", domain.DomainName)
	parameter.Set("content", *request.Value)
	parameter.Set("ttl", "120")
	return parameter, nil
}

func (_this *CloudflareParameter) LoadUpdateParameter(request *model.UpdateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.CREATE_REQUEST_NOT_NIL)
	}
	parameter := _this.loadCommonParameter(request.RecordType)
	// assert domain
	domain, err := vhttp.CheckExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, err
	}
	if request.Value == nil {
		return nil, errs.NewVdnsError(msg.RECORD_VALUE_NOT_SUPPORT)
	}

	parameter.Set("name", domain.DomainName)
	parameter.Set("content", *request.Value)
	parameter.Set("ttl", "120")
	return parameter, nil
}

func (_this *CloudflareParameter) LoadDeleteParameter(request *model.DeleteDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.CREATE_REQUEST_NOT_NIL)
	}
	parameter := _this.loadCommonParameter("")
	return parameter, nil
}

func (_this *CloudflareParameter) loadCommonParameter(record record.Type) *url.Values {
	parameter := make(url.Values, 10)
	parameter.Set("match", "all")
	parameter.Set("type", record.String())
	return &parameter
}
