package parameter

import (
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
	time2 "vdns/lib/standard/time"
	"vdns/vutil/strs"
	"vdns/vutil/vhttp"
)

type AlidnsParameterProvier struct {
	credential        auth.Credential
	signatureComposer compose.SignatureComposer
	version           *string
}

func NewAlidnsParameterProvider(credential auth.Credential, signatureComposer compose.SignatureComposer) ParamaterProvider {
	return &AlidnsParameterProvier{
		credential:        credential,
		signatureComposer: signatureComposer,
		version:           strs.String(standard.ALIDNS_API_VERSION),
	}
}

func (_this *AlidnsParameterProvier) LoadDescribeParamater(request *models.DescribeDomainRecordsRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}

	// assert domain
	extractDomain, err := vhttp.ExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewApiErrorFromError(err)
	}
	domain := extractDomain[0]
	rr := extractDomain[1]
	paramter := _this.loadCommonParamter(action)
	paramter.Set("DomainName", domain)

	// assert record type
	if !record.Support(request.RecordType) {
		return nil, errs.NewApiError(msg.RECORD_TYPE_NOT_SUPPORT)
	}
	paramter.Set("TypeKeyWord", request.RecordType.String())

	// assert page size
	if request.PageSize != nil {
		paramter.Set("PageSize", strconv.FormatInt(*request.PageSize, 10))
	}

	// assert page number
	if request.PageNumber != nil {
		paramter.Set("PageNumber", strconv.FormatInt(*request.PageNumber, 10))
	}

	// record value keyword (fuzzy match before and after) pattern search, not case-sensitive.
	if request.ValueKeyWord != nil {
		paramter.Set("ValueKeyWord", *request.ValueKeyWord)
	}

	// the keywords recorded by the host, (fuzzy matching before and after) pattern search, are not case-sensitive.
	if request.RRKeyWord != nil {
		paramter.Set("RRKeyWord", *request.RRKeyWord)
	} else if strs.NotEmpty(rr) {
		paramter.Set("RRKeyWord", rr)
	}

	return paramter, nil
}

func (_this *AlidnsParameterProvier) LoadCreateParamater(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.CREATE_REQUEST_NOT_NIL)
	}

	// assert record type
	if !record.Support(request.RecordType) {
		return nil, errs.NewApiError(msg.RECORD_TYPE_NOT_SUPPORT)
	}

	// assert value
	if request.Value == nil {
		return nil, errs.NewApiError(msg.RECORD_VALUE_NOT_SUPPORT)
	}

	// assert domain
	extractDomain, err := vhttp.ExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewApiErrorFromError(err)
	}
	domain := extractDomain[0]
	rr := extractDomain[1]
	paramter := _this.loadCommonParamter(action)
	paramter.Set("DomainName", domain)
	paramter.Set("Type", request.RecordType.String())
	paramter.Set("Value", *request.Value)

	// assert rr
	if strs.IsEmpty(rr) {
		paramter.Set("RR", standard.PanAnalysisRRKeyWord)
	} else {
		paramter.Set("RR", rr)
	}

	return paramter, nil
}

func (_this *AlidnsParameterProvier) LoadUpdateParamater(request *models.UpdateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.CREATE_REQUEST_NOT_NIL)
	}

	// assert record id
	if request.ID == nil {
		return nil, errs.NewApiError(msg.RECORD_ID_NOT_SUPPORT)
	}

	// assert record type
	if !record.Support(request.RecordType) {
		return nil, errs.NewApiError(msg.RECORD_TYPE_NOT_SUPPORT)
	}

	// assert value
	if request.Value == nil {
		return nil, errs.NewApiError(msg.RECORD_VALUE_NOT_SUPPORT)
	}

	// assert domain
	extractDomain, err := vhttp.ExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, err
	}
	domain := extractDomain[0]
	rr := extractDomain[1]

	paramter := _this.loadCommonParamter(action)
	paramter.Set("DomainName", domain)
	paramter.Set("RecordId", *request.ID)
	paramter.Set("Type", request.RecordType.String())
	paramter.Set("Value", *request.Value)

	// assert rr
	if strs.IsEmpty(rr) {
		paramter.Set("RR", standard.PanAnalysisRRKeyWord)
	} else {
		paramter.Set("RR", rr)
	}

	return paramter, nil
}

func (_this *AlidnsParameterProvier) LoadDeleteParamater(request *models.DeleteDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.CREATE_REQUEST_NOT_NIL)
	}

	// assert record id
	if request.ID == nil {
		return nil, errs.NewApiError(msg.RECORD_ID_NOT_SUPPORT)
	}

	paramter := _this.loadCommonParamter(action)
	paramter.Set("RecordId", *request.ID)
	return paramter, nil
}

func (_this *AlidnsParameterProvier) loadCommonParamter(action *string) *url.Values {
	timestamp := time.Now().UTC().Format(time2.ALIYUN_FORMAT_ISO8601)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	paramater := make(url.Values, 10)
	paramater.Set("Action", strs.StringValue(action))
	paramater.Set("AccessKeyId", _this.credential.GetSecretId())
	paramater.Set("Format", "JSON")
	paramater.Set("SignatureMethod", _this.signatureComposer.SignatureMethod())
	paramater.Set("SignatureNonce", nonce)
	paramater.Set("SignatureVersion", _this.signatureComposer.SignerVersion())
	paramater.Set("Timestamp", timestamp)
	paramater.Set("Version", "2015-01-09")
	return &paramater
}
