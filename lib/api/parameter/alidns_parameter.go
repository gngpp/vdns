package parameter

import (
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
	time2 "vdns/lib/standard/time"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
)

type AliDNSParameter struct {
	credential        auth.Credential
	signatureComposer compose.SignatureComposer
	version           *standard.Standard
}

func NewAliDNSParameter(credential auth.Credential, signatureComposer compose.SignatureComposer) Parameter {
	return &AliDNSParameter{
		credential:        credential,
		signatureComposer: signatureComposer,
		version:           standard.ALIDNS_API_VERSION.String(),
	}
}

func (_this *AliDNSParameter) LoadDescribeParameter(request *model.DescribeDomainRecordsRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}

	// assert domain
	domain, err := vhttp.CheckExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewVdnsFromError(err)
	}
	parameter := _this.loadCommonParameter(action)
	parameter.Set(AlidnsParameterDoaminName, domain.DomainName)

	// assert record type
	if !record.Support(request.RecordType) {
		return nil, errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
	}
	parameter.Set(AlidnsParameterTypeKeyWord, request.RecordType.String())

	// assert page size
	if request.PageSize != nil {
		parameter.Set(AlidnsParameterPageSize, strconv.FormatInt(*request.PageSize, 10))
	}

	// assert page number
	if request.PageNumber != nil {
		parameter.Set(AlidnsParameterPageNumber, strconv.FormatInt(*request.PageNumber, 10))
	}

	// record value keyword (fuzzy match before and after) pattern search, not case-sensitive.
	if request.ValueKeyWord != nil {
		parameter.Set(AlidnsParameterValueKeyWord, *request.ValueKeyWord)
	}

	// the keywords recorded by the host, (fuzzy matching before and after) pattern search, are not case-sensitive.
	if request.RRKeyWord != nil {
		parameter.Set(AlidnsParameterRrKeyWord, *request.RRKeyWord)
	} else if strs.NotEmpty(domain.SubDomain) {
		parameter.Set(AlidnsParameterRrKeyWord, domain.SubDomain)
	}

	return parameter, nil
}

func (_this *AliDNSParameter) LoadCreateParameter(request *model.CreateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.CREATE_REQUEST_NOT_NIL)
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
	parameter.Set(AlidnsParameterDoaminName, domain.DomainName)
	parameter.Set(AlidnsParameterType, request.RecordType.String())
	parameter.Set(AlidnsParameterValue, *request.Value)

	// assert rr
	if strs.IsEmpty(domain.SubDomain) {
		parameter.Set(AlidnsParameterRr, record.PAN_ANALYSIS_RR_KEY_WORD.String())
	} else {
		parameter.Set(AlidnsParameterRr, domain.SubDomain)
	}

	return parameter, nil
}

func (_this *AliDNSParameter) LoadUpdateParameter(request *model.UpdateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.CREATE_REQUEST_NOT_NIL)
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
		return nil, err
	}

	parameter := _this.loadCommonParameter(action)
	parameter.Set(AlidnsParameterRecordId, *request.ID)
	parameter.Set(AlidnsParameterDoaminName, domain.DomainName)
	parameter.Set(AlidnsParameterType, request.RecordType.String())
	parameter.Set(AlidnsParameterValue, *request.Value)

	// assert rr
	if strs.IsEmpty(domain.SubDomain) {
		parameter.Set(AlidnsParameterRr, record.PAN_ANALYSIS_RR_KEY_WORD.String())
	} else {
		parameter.Set(AlidnsParameterRr, domain.SubDomain)
	}

	return parameter, nil
}

func (_this *AliDNSParameter) LoadDeleteParameter(request *model.DeleteDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.CREATE_REQUEST_NOT_NIL)
	}

	// assert record id
	if request.ID == nil {
		return nil, errs.NewVdnsError(msg.RECORD_ID_NOT_SUPPORT)
	}

	parameter := _this.loadCommonParameter(action)
	parameter.Set(AlidnsParameterRecordId, *request.ID)
	return parameter, nil
}

func (_this *AliDNSParameter) loadCommonParameter(action *string) *url.Values {
	timestamp := time.Now().UTC().Format(time2.ALIYUN_FORMAT_ISO8601)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	parameter := make(url.Values, 10)
	parameter.Set(AlidnsParameterAction, strs.StringValue(action))
	parameter.Set(AlidnsParameterAccessKeyId, _this.credential.GetSecretId())
	parameter.Set(AlidnsParameterFormat, AlidnsParameterJson)
	parameter.Set(AlidnsParameterSignatureMethod, _this.signatureComposer.SignatureMethod())
	parameter.Set(AlidnsParameterSignatureNonce, nonce)
	parameter.Set(AlidnsParameterSignatureVersion, _this.signatureComposer.SignerVersion())
	parameter.Set(AlidnsParameterTimestamp, timestamp)
	parameter.Set(AlidnsParameterVersion, _this.version.StringValue())
	return &parameter
}
