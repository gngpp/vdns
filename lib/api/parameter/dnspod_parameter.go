package parameter

import (
	"math/rand"
	"net/url"
	"strconv"
	"time"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/auth"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
	"vdns/lib/standard/msg"
	"vdns/lib/standard/record"
	"vdns/lib/util/convert"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
)

type DNSPodParameter struct {
	credential        auth.Credential
	signatureComposer compose.SignatureComposer
	version           *standard.Standard
}

func NewDNSPodParameter(credential auth.Credential, signatureComposer compose.SignatureComposer) Parameter {
	return &DNSPodParameter{
		credential:        credential,
		signatureComposer: signatureComposer,
		version:           standard.DNSPOD_API_VERSION.String(),
	}
}

func (_this *DNSPodParameter) LoadDescribeParameter(request *model.DescribeDomainRecordsRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}
	// assert domain
	domain, err := vhttp.CheckExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewVdnsFromError(err)
	}
	parameter := _this.loadCommonParameter(action)
	parameter.Set(DnspodParameterDomain, domain.DomainName)

	// assert record type
	if !record.Support(request.RecordType) {
		return nil, errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
	}
	parameter.Set(DnspodParameterRecordType, request.RecordType.String())

	// assert page size
	if request.PageSize != nil {
		parameter.Set(DnspodParameterLimit, convert.AsStringValue(request.PageSize))
	}

	// assert offset start from 0
	if request.PageNumber != nil {
		parameter.Set(DnspodParameterOffset, convert.AsStringValue(*request.PageNumber-1))
	}

	// search and parse records by keyword, currently supports searching for host headers and record values
	if request.ValueKeyWord != nil {
		parameter.Set(DnspodParameterKeyWord, *request.ValueKeyWord)
	}

	// assert rr key word
	if request.RRKeyWord != nil {
		parameter.Set(DnspodParameterSubdomain1, *request.RRKeyWord)
	} else if strs.NotEmpty(domain.SubDomain) {
		parameter.Set(DnspodParameterSubdomain1, domain.SubDomain)
	}
	return parameter, nil
}

func (_this *DNSPodParameter) LoadCreateParameter(request *model.CreateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}

	// assert record type
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
		return nil, errs.NewVdnsFromError(err)
	}
	parameter := _this.loadCommonParameter(action)
	parameter.Set(DnspodParameterDomain, domain.DomainName)
	parameter.Set(DnspodParameterRecordType, request.RecordType.String())
	parameter.Set(DnspodParameterValue, strs.StringValue(request.Value))
	parameter.Set(DnspodParameterRecordLine, DnspodParameterDefault)

	// assert rr
	if strs.IsEmpty(domain.SubDomain) {
		parameter.Set(DnspodParameterSubdomain2, record.PAN_ANALYSIS_RR_KEY_WORD.String())
	} else {
		parameter.Set(DnspodParameterSubdomain2, domain.SubDomain)
	}
	return parameter, nil
}

func (_this *DNSPodParameter) LoadUpdateParameter(request *model.UpdateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}

	// assert record id
	if request.ID == nil {
		return nil, errs.NewVdnsError(msg.RECORD_ID_NOT_SUPPORT)
	}

	// assert record type
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
		return nil, errs.NewVdnsFromError(err)
	}
	parameter := _this.loadCommonParameter(action)
	parameter.Set(DnspodParameterRecordId, *request.ID)
	parameter.Set(DnspodParameterDomain, domain.DomainName)
	parameter.Set(DnspodParameterRecordType, request.RecordType.String())
	parameter.Set(DnspodParameterValue, strs.StringValue(request.Value))
	parameter.Set(DnspodParameterRecordLine, DnspodParameterDefault)

	// assert rr
	if strs.IsEmpty(domain.SubDomain) {
		parameter.Set(DnspodParameterSubdomain2, record.PAN_ANALYSIS_RR_KEY_WORD.String())
	} else {
		parameter.Set(DnspodParameterSubdomain2, domain.SubDomain)
	}

	return parameter, nil
}

func (_this *DNSPodParameter) LoadDeleteParameter(request *model.DeleteDomainRecordRequest, action *string) (*url.Values, error) {
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
		return nil, errs.NewVdnsFromError(err)
	}
	parameter := _this.loadCommonParameter(action)
	parameter.Set(DnspodParameterRecordId, *request.ID)
	parameter.Set(DnspodParameterDomain, domain.DomainName)
	return parameter, nil
}

func (_this *DNSPodParameter) loadCommonParameter(action *string) *url.Values {
	parameter := make(url.Values, 10)
	nonce := strconv.FormatInt(rand.Int63()+time.Now().UnixMilli(), 10)
	timestamp := strconv.FormatInt(time.Now().UnixMilli()/1000, 10)
	parameter.Set(DnspodParameterAction, strs.StringValue(action))
	parameter.Set(DnspodParameterNonce, nonce)
	parameter.Set(DnspodParameterTimestamp, timestamp)
	parameter.Set(DnspodParameterSecretId, _this.credential.GetSecretId())
	parameter.Set(DnspodParameterSignaturemethod, _this.signatureComposer.SignatureMethod())
	parameter.Set(DnspodParameterVersion, _this.version.StringValue())

	return &parameter
}
