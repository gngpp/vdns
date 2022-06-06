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

func (_this *CloudflareParameter) LoadDescribeParameter(request *model.DescribeDomainRecordsRequest, _ *string) (*Value, error) {
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
	return NewValue(parameter, nil), nil
}

func (_this *CloudflareParameter) LoadCreateParameter(request *model.CreateDomainRecordRequest, _ *string) (*Value, error) {
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

	parameter.Set(CfParameterName, domain.DomainName)
	parameter.Set(CfParameterContent, *request.Value)
	parameter.Set(CfParameterTTL, "120")
	return NewValue(parameter, nil), nil
}

func (_this *CloudflareParameter) LoadUpdateParameter(request *model.UpdateDomainRecordRequest, _ *string) (*Value, error) {
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

	parameter.Set(CfParameterName, domain.DomainName)
	parameter.Set(CfParameterContent, *request.Value)
	parameter.Set(CfParameterTTL, "120")
	return NewValue(parameter, nil), nil
}

func (_this *CloudflareParameter) LoadDeleteParameter(_ *model.DeleteDomainRecordRequest, _ *string) (*Value, error) {
	return NewValue(nil, nil), nil
}

func (_this *CloudflareParameter) loadCommonParameter(record record.Type) *url.Values {
	parameter := make(url.Values, 10)
	parameter.Set(CfParameterMatch, "all")
	parameter.Set(CfParameterType, record.String())
	return &parameter
}
