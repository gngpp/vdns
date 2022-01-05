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
	version           *standard.Standard
}

func NewAlidnsParameterProvider(credential auth.Credential, signatureComposer compose.SignatureComposer) ParamaterProvider {
	return &AlidnsParameterProvier{
		credential:        credential,
		signatureComposer: signatureComposer,
		version:           standard.ALIDNS_API_VERSION.String(),
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
	paramter.Set(ALIDNS_PARAMETER_DOAMIN_NAME, domain)

	// assert record type
	if !record.Support(request.RecordType) {
		return nil, errs.NewApiError(msg.RECORD_TYPE_NOT_SUPPORT)
	}
	paramter.Set(ALIDNS_PARAMETER_TYPE_KEY_WORD, request.RecordType.String())

	// assert page size
	if request.PageSize != nil {
		paramter.Set(ALIDNS_PARAMETER_PAGE_SIZE, strconv.FormatInt(*request.PageSize, 10))
	}

	// assert page number
	if request.PageNumber != nil {
		paramter.Set(ALIDNS_PARAMETER_PAGE_NUMBER, strconv.FormatInt(*request.PageNumber, 10))
	}

	// record value keyword (fuzzy match before and after) pattern search, not case-sensitive.
	if request.ValueKeyWord != nil {
		paramter.Set(ALIDNS_PARAMETER_VALUE_KEY_WORD, *request.ValueKeyWord)
	}

	// the keywords recorded by the host, (fuzzy matching before and after) pattern search, are not case-sensitive.
	if request.RRKeyWord != nil {
		paramter.Set(ALIDNS_PARAMETER_RR_KEY_WORD, *request.RRKeyWord)
	} else if strs.NotEmpty(rr) {
		paramter.Set(ALIDNS_PARAMETER_RR_KEY_WORD, rr)
	}

	return paramter, nil
}

func (_this *AlidnsParameterProvier) LoadCreateParamater(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.CREATE_REQUEST_NOT_NIL)
	}

	// assert record type
	if request.RecordType != nil && !record.Support(*request.RecordType) {
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
	paramter.Set(ALIDNS_PARAMETER_DOAMIN_NAME, domain)
	paramter.Set(ALIDNS_PARAMETER_TYPE, request.RecordType.String())
	paramter.Set(ALIDNS_PARAMETER_VALUE, *request.Value)

	// assert rr
	if strs.IsEmpty(rr) {
		paramter.Set(ALIDNS_PARAMETER_RR, record.PAN_ANALYSIS_RR_KEY_WORD.String())
	} else {
		paramter.Set(ALIDNS_PARAMETER_RR, rr)
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
	paramter.Set(ALIDNS_PARAMETER_RECORD_ID, *request.ID)
	paramter.Set(ALIDNS_PARAMETER_DOAMIN_NAME, domain)
	paramter.Set(ALIDNS_PARAMETER_TYPE, request.RecordType.String())
	paramter.Set(ALIDNS_PARAMETER_VALUE, *request.Value)

	// assert rr
	if strs.IsEmpty(rr) {
		paramter.Set(ALIDNS_PARAMETER_RR, record.PAN_ANALYSIS_RR_KEY_WORD.String())
	} else {
		paramter.Set(ALIDNS_PARAMETER_RR, rr)
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
	paramter.Set(ALIDNS_PARAMETER_RECORD_ID, *request.ID)
	return paramter, nil
}

func (_this *AlidnsParameterProvier) loadCommonParamter(action *string) *url.Values {
	timestamp := time.Now().UTC().Format(time2.ALIYUN_FORMAT_ISO8601)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	paramater := make(url.Values, 10)
	paramater.Set(ALIDNS_PARAMETER_ACTION, strs.StringValue(action))
	paramater.Set(ALIDNS_PARAMETER_ACCESS_KEY_ID, _this.credential.GetSecretId())
	paramater.Set(ALIDNS_PARAMETER_FORMAT, ALIDNS_PARAMETER_JSON)
	paramater.Set(ALIDNS_PARAMETER_SIGNATURE_METHOD, _this.signatureComposer.SignatureMethod())
	paramater.Set(ALIDNS_PARAMETER_SIGNATURE_NONCE, nonce)
	paramater.Set(ALIDNS_PARAMETER_SIGNATURE_VERSION, _this.signatureComposer.SignerVersion())
	paramater.Set(ALIDNS_PARAMETER_TIMESTAMP, timestamp)
	paramater.Set(ALIDNS_PARAMETER_VERSION, _this.version.StringValue())
	return &paramater
}
