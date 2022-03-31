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
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
)

type AliDNSParameterProvier struct {
	credential        auth.Credential
	signatureComposer compose.SignatureComposer
	version           *standard.Standard
}

func NewAliDNSParameterProvider(credential auth.Credential, signatureComposer compose.SignatureComposer) ParamaterProvider {
	return &AliDNSParameterProvier{
		credential:        credential,
		signatureComposer: signatureComposer,
		version:           standard.ALIDNS_API_VERSION.String(),
	}
}

func (_this *AliDNSParameterProvier) LoadDescribeParamater(request *models.DescribeDomainRecordsRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}

	// assert domain
	domain, err := vhttp.CheckExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewApiErrorFromError(err)
	}
	parameter := _this.loadCommonParamter(action)
	parameter.Set(ALIDNS_PARAMETER_DOAMIN_NAME, domain.DomainName)

	// assert record type
	if !record.Support(request.RecordType) {
		return nil, errs.NewVdnsError(msg.RECORD_TYPE_NOT_SUPPORT)
	}
	parameter.Set(ALIDNS_PARAMETER_TYPE_KEY_WORD, request.RecordType.String())

	// assert page size
	if request.PageSize != nil {
		parameter.Set(ALIDNS_PARAMETER_PAGE_SIZE, strconv.FormatInt(*request.PageSize, 10))
	}

	// assert page number
	if request.PageNumber != nil {
		parameter.Set(ALIDNS_PARAMETER_PAGE_NUMBER, strconv.FormatInt(*request.PageNumber, 10))
	}

	// record value keyword (fuzzy match before and after) pattern search, not case-sensitive.
	if request.ValueKeyWord != nil {
		parameter.Set(ALIDNS_PARAMETER_VALUE_KEY_WORD, *request.ValueKeyWord)
	}

	// the keywords recorded by the host, (fuzzy matching before and after) pattern search, are not case-sensitive.
	if request.RRKeyWord != nil {
		parameter.Set(ALIDNS_PARAMETER_RR_KEY_WORD, *request.RRKeyWord)
	} else if strs.NotEmpty(domain.SubDomain) {
		parameter.Set(ALIDNS_PARAMETER_RR_KEY_WORD, domain.SubDomain)
	}

	return parameter, nil
}

func (_this *AliDNSParameterProvier) LoadCreateParamater(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.CREATE_REQUEST_NOT_NIL)
	}

	// assert record type
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
	parameter.Set(ALIDNS_PARAMETER_DOAMIN_NAME, domain.DomainName)
	parameter.Set(ALIDNS_PARAMETER_TYPE, request.RecordType.String())
	parameter.Set(ALIDNS_PARAMETER_VALUE, *request.Value)

	// assert rr
	if strs.IsEmpty(domain.SubDomain) {
		parameter.Set(ALIDNS_PARAMETER_RR, record.PAN_ANALYSIS_RR_KEY_WORD.String())
	} else {
		parameter.Set(ALIDNS_PARAMETER_RR, domain.SubDomain)
	}

	return parameter, nil
}

func (_this *AliDNSParameterProvier) LoadUpdateParamater(request *models.UpdateDomainRecordRequest, action *string) (*url.Values, error) {
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

	parameter := _this.loadCommonParamter(action)
	parameter.Set(ALIDNS_PARAMETER_RECORD_ID, *request.ID)
	parameter.Set(ALIDNS_PARAMETER_DOAMIN_NAME, domain.DomainName)
	parameter.Set(ALIDNS_PARAMETER_TYPE, request.RecordType.String())
	parameter.Set(ALIDNS_PARAMETER_VALUE, *request.Value)

	// assert rr
	if strs.IsEmpty(domain.SubDomain) {
		parameter.Set(ALIDNS_PARAMETER_RR, record.PAN_ANALYSIS_RR_KEY_WORD.String())
	} else {
		parameter.Set(ALIDNS_PARAMETER_RR, domain.SubDomain)
	}

	return parameter, nil
}

func (_this *AliDNSParameterProvier) LoadDeleteParamater(request *models.DeleteDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewVdnsError(msg.CREATE_REQUEST_NOT_NIL)
	}

	// assert record id
	if request.ID == nil {
		return nil, errs.NewVdnsError(msg.RECORD_ID_NOT_SUPPORT)
	}

	parameter := _this.loadCommonParamter(action)
	parameter.Set(ALIDNS_PARAMETER_RECORD_ID, *request.ID)
	return parameter, nil
}

func (_this *AliDNSParameterProvier) loadCommonParamter(action *string) *url.Values {
	timestamp := time.Now().UTC().Format(time2.ALIYUN_FORMAT_ISO8601)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	parameter := make(url.Values, 10)
	parameter.Set(ALIDNS_PARAMETER_ACTION, strs.StringValue(action))
	parameter.Set(ALIDNS_PARAMETER_ACCESS_KEY_ID, _this.credential.GetSecretId())
	parameter.Set(ALIDNS_PARAMETER_FORMAT, ALIDNS_PARAMETER_JSON)
	parameter.Set(ALIDNS_PARAMETER_SIGNATURE_METHOD, _this.signatureComposer.SignatureMethod())
	parameter.Set(ALIDNS_PARAMETER_SIGNATURE_NONCE, nonce)
	parameter.Set(ALIDNS_PARAMETER_SIGNATURE_VERSION, _this.signatureComposer.SignerVersion())
	parameter.Set(ALIDNS_PARAMETER_TIMESTAMP, timestamp)
	parameter.Set(ALIDNS_PARAMETER_VERSION, _this.version.StringValue())
	return &parameter
}
