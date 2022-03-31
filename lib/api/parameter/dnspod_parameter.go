package parameter

import (
	"math/rand"
	"net/url"
	"strconv"
	"time"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	"vdns/lib/auth"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
	"vdns/lib/standard/msg"
	"vdns/lib/standard/record"
	"vdns/lib/util/convert"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
)

type DNSPodParameterProvider struct {
	credential        auth.Credential
	signatureComposer compose.SignatureComposer
	version           *standard.Standard
}

func NewDNSPodParameterProvider(credential auth.Credential, signatureComposer compose.SignatureComposer) ParamaterProvider {
	return &DNSPodParameterProvider{
		credential:        credential,
		signatureComposer: signatureComposer,
		version:           standard.DNSPOD_API_VERSION.String(),
	}
}

func (_this *DNSPodParameterProvider) LoadDescribeParameter(request *models.DescribeDomainRecordsRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}
	// assert domain
	domain, err := vhttp.CheckExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewApiErrorFromError(err)
	}
	parameter := _this.loadCommonParamter(action)
	parameter.Set(DNSPOD_PARAMETER_DOMAIN, domain.DomainName)

	// assert record type
	if !record.Support(request.RecordType) {
		return nil, errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
	}
	parameter.Set(DNSPOD_PARAMETER_RECORD_TYPE, request.RecordType.String())

	// assert page size
	if request.PageSize != nil {
		parameter.Set(DNSPOD_PARAMETER_LIMIT, convert.AsStringValue(request.PageSize))
	}

	// assert offset start from 0
	if request.PageNumber != nil {
		parameter.Set(DNSPOD_PARAMETER_OFFSET, convert.AsStringValue(*request.PageNumber-1))
	}

	// search and parse records by keyword, currently supports searching for host headers and record values
	if request.ValueKeyWord != nil {
		parameter.Set(DNSPOD_PARAMETER_KEY_WORD, *request.ValueKeyWord)
	}

	// assert rr key word
	if request.RRKeyWord != nil {
		parameter.Set(DNSPOD_PARAMETER_SUBDOMAIN_1, *request.RRKeyWord)
	} else if strs.NotEmpty(domain.SubDomain) {
		parameter.Set(DNSPOD_PARAMETER_SUBDOMAIN_1, domain.SubDomain)
	}
	return parameter, nil
}

func (_this *DNSPodParameterProvider) LoadCreateParameter(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}

	// assert record type assert
	if request.RecordType != nil && !record.Support(*request.RecordType) {
		return nil, errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
	}

	// assert value
	if request.Value == nil {
		return nil, errs.NewVdnsError(msg.RECORD_VALUE_NOT_SUPPORT)
	}

	// assert domain
	domain, err := vhttp.CheckExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewApiErrorFromError(err)
	}
	parameter := _this.loadCommonParamter(action)
	parameter.Set(DNSPOD_PARAMETER_DOMAIN, domain.DomainName)
	parameter.Set(DNSPOD_PARAMETER_RECORD_TYPE, request.RecordType.String())
	parameter.Set(DNSPOD_PARAMETER_VALUE, strs.StringValue(request.Value))
	parameter.Set(DNSPOD_PARAMETER_RECORD_LINE, DNSPOD_PARAMETER_DEFAULT)

	// assert rr
	if strs.IsEmpty(domain.SubDomain) {
		parameter.Set(DNSPOD_PARAMETER_SUBDOMAIN_2, record.PAN_ANALYSIS_RR_KEY_WORD.String())
	} else {
		parameter.Set(DNSPOD_PARAMETER_SUBDOMAIN_2, domain.SubDomain)
	}
	return parameter, nil
}

func (_this *DNSPodParameterProvider) LoadUpdateParameter(request *models.UpdateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}

	// assert record id
	if request.ID == nil {
		return nil, errs.NewVdnsError(msg.RECORD_ID_NOT_SUPPORT)
	}

	// assert record type assert
	if !record.Support(request.RecordType) {
		return nil, errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
	}

	// assert value
	if request.Value == nil {
		return nil, errs.NewVdnsError(msg.RECORD_VALUE_NOT_SUPPORT)
	}

	// assert domain
	domain, err := vhttp.CheckExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewApiErrorFromError(err)
	}
	parameter := _this.loadCommonParamter(action)
	parameter.Set(DNSPOD_PARAMETER_RECORD_ID, *request.ID)
	parameter.Set(DNSPOD_PARAMETER_DOMAIN, domain.DomainName)
	parameter.Set(DNSPOD_PARAMETER_RECORD_TYPE, request.RecordType.String())
	parameter.Set(DNSPOD_PARAMETER_VALUE, strs.StringValue(request.Value))
	parameter.Set(DNSPOD_PARAMETER_RECORD_LINE, DNSPOD_PARAMETER_DEFAULT)

	// assert rr
	if strs.IsEmpty(domain.SubDomain) {
		parameter.Set(DNSPOD_PARAMETER_SUBDOMAIN_2, record.PAN_ANALYSIS_RR_KEY_WORD.String())
	} else {
		parameter.Set(DNSPOD_PARAMETER_SUBDOMAIN_2, domain.SubDomain)
	}

	return parameter, nil
}

func (_this *DNSPodParameterProvider) LoadDeleteParameter(request *models.DeleteDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}

	// assert record id
	if request.ID == nil {
		return nil, errs.NewVdnsError(msg.RECORD_ID_NOT_SUPPORT)
	}

	// assert domain
	domain, err := vhttp.CheckExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewApiErrorFromError(err)
	}
	parameter := _this.loadCommonParamter(action)
	parameter.Set(DNSPOD_PARAMETER_RECORD_ID, *request.ID)
	parameter.Set(DNSPOD_PARAMETER_DOMAIN, domain.DomainName)
	return parameter, nil
}

func (_this *DNSPodParameterProvider) loadCommonParamter(action *string) *url.Values {
	parameter := make(url.Values, 10)
	nonce := strconv.FormatInt(rand.Int63()+time.Now().UnixMilli(), 10)
	timestamp := strconv.FormatInt(time.Now().UnixMilli()/1000, 10)
	parameter.Set(DNSPOD_PARAMETER_ACTION, strs.StringValue(action))
	parameter.Set(DNSPOD_PARAMETER_NONCE, nonce)
	parameter.Set(DNSPOD_PARAMETER_TIMESTAMP, timestamp)
	parameter.Set(DNSPOD_PARAMETER_SECRET_ID, _this.credential.GetSecretId())
	parameter.Set(DNSPOD_PARAMETER_SIGNATUREMETHOD, _this.signatureComposer.SignatureMethod())
	parameter.Set(DNSPOD_PARAMETER_VERSION, _this.version.StringValue())

	return &parameter
}
