package api

import (
	"encoding/base64"
	"net/url"
	"strconv"
	"time"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	"vdns/lib/api/rpc"
	"vdns/lib/auth"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
	"vdns/lib/standard/msg"
	"vdns/lib/standard/record"
	time2 "vdns/lib/standard/time"
	"vdns/vutil/strs"
	"vdns/vutil/vhttp"
)

func NewAlidnsProvider(credential auth.Credential) DNSRecordProvider {
	return &AlidnsProvider{
		Action: &Action{
			describe: strs.String("DescribeDomainRecords"),
			create:   strs.String("AddDomainRecord"),
			update:   strs.String("UpdateDomainRecord"),
			delete:   strs.String("DeleteDomainRecord"),
		},
		compose:   compose.NewAlidnsSignatureCompose(compose.SEPARATOR),
		rpc:       rpc.NewAlidnsRpc(),
		api:       strs.String(standard.ALIYUN_DNS_API),
		credntial: credential,
	}
}

type AlidnsProvider struct {
	*Action
	api       *string
	compose   compose.SignatureComposer
	rpc       rpc.Rpc
	credntial auth.Credential
}

func (_this *AlidnsProvider) DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordsResponse, error) {
	paramater, err := _this.loadDescribeParamater(request, _this.describe)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoDescribeRequest(requestUrl)
}

func (_this *AlidnsProvider) CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.loadCreateParamater(request, _this.create)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoCreateRequest(requestUrl)
}

func (_this *AlidnsProvider) UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.loadUpdateParamater(request, _this.update)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoUpdateRequest(requestUrl)
}

func (_this *AlidnsProvider) DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainRecordStatusResponse, error) {
	paramater, err := _this.loadDeleteParamater(request, _this.delete)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.rpc.DoDeleteRequest(requestUrl)
}

func (_this *AlidnsProvider) Support(recordType record.Type) bool {
	return record.Support(recordType)
}

func (_this *AlidnsProvider) generateRequestUrl(paramater *url.Values) string {
	stringToSign := _this.compose.ComposeStringToSign(vhttp.HttpMethodGet, paramater)
	signature := _this.compose.GeneratedSignature(_this.credntial.GetSecretKey(), stringToSign)
	_this.SetSignature(signature, paramater)
	return _this.compose.CanonicalizeRequestUrl(strs.StringValue(_this.api), paramater)
}

func (_this *AlidnsProvider) SetSignature(signature string, queries *url.Values) {
	queries.Set("Signature", base64.StdEncoding.EncodeToString(strs.ToBytes(signature)))
}

func (_this *AlidnsProvider) loadDescribeParamater(request *models.DescribeDomainRecordsRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}
	extractDomain, err := vhttp.ExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewApiErrorFromError(err)
	}
	domain := extractDomain[0]
	rr := extractDomain[1]
	paramter := _this.getCommonParamter(action)
	if request.PageSize != nil {
		paramter.Set("PageSize", strconv.FormatInt(*request.PageSize, 10))
	}
	if request.PageNumber != nil {
		paramter.Set("PageNumber", strconv.FormatInt(*request.PageNumber, 10))
	}
	if record.Support(request.RecordType) {
		paramter.Set("TypeKeyWord", request.RecordType.String())
	} else {
		return nil, errs.NewApiError(msg.RECORD_TYPE_NOT_SUPPORT)
	}
	// Record value keyword (fuzzy match before and after) pattern search, not case-sensitive.
	if request.ValueKeyWord != nil {
		paramter.Set("ValueKeyWord", *request.ValueKeyWord)
	}
	paramter.Set("DomainName", domain)
	// The keywords recorded by the host, (fuzzy matching before and after) pattern search, are not case-sensitive.
	condition := true
	if request.RRKeyWord != nil {
		condition = false
		paramter.Set("RRKeyWord", *request.RRKeyWord)
	}
	if condition && strs.NotEmpty(rr) {
		paramter.Set("RRKeyWord", rr)
	}
	return paramter, nil
}

func (_this *AlidnsProvider) loadCreateParamater(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.CREATE_REQUEST_NOT_NIL)
	}
	if !record.Support(request.RecordType) {
		return nil, errs.NewApiError(msg.RECORD_TYPE_NOT_SUPPORT)

	}
	if request.Value == nil {
		return nil, errs.NewApiError(msg.RECORD_VALUE_NOT_SUPPORT)
	}
	extractDomain, err := vhttp.ExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, errs.NewApiErrorFromError(err)
	}
	domain := extractDomain[0]
	rr := extractDomain[1]
	paramter := _this.getCommonParamter(action)
	paramter.Set("Type", request.RecordType.String())
	paramter.Set("Value", *request.Value)
	paramter.Set("DomainName", domain)
	paramter.Set("RR", rr)
	return paramter, nil
}

func (_this *AlidnsProvider) loadUpdateParamater(request *models.UpdateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.CREATE_REQUEST_NOT_NIL)
	}
	if request.ID == nil {
		return nil, errs.NewApiError(msg.RECORD_ID_NOT_SUPPORT)
	}
	if !record.Support(request.RecordType) {
		return nil, errs.NewApiError(msg.RECORD_TYPE_NOT_SUPPORT)

	}
	if request.Value == nil {
		return nil, errs.NewApiError(msg.RECORD_VALUE_NOT_SUPPORT)
	}
	extractDomain, err := vhttp.ExtractDomain(strs.StringValue(request.Domain))
	if err != nil {
		return nil, err
	}
	domain := extractDomain[0]
	rr := extractDomain[1]
	paramter := _this.getCommonParamter(action)
	paramter.Set("RecordId", *request.ID)
	paramter.Set("Type", request.RecordType.String())
	paramter.Set("Value", *request.Value)
	paramter.Set("DomainName", domain)
	paramter.Set("RR", rr)
	return paramter, nil
}

func (_this *AlidnsProvider) loadDeleteParamater(request *models.DeleteDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.CREATE_REQUEST_NOT_NIL)
	}
	paramter := _this.getCommonParamter(action)
	if request.ID == nil {
		return nil, errs.NewApiError(msg.RECORD_ID_NOT_SUPPORT)
	}
	paramter.Set("RecordId", *request.ID)
	return paramter, nil
}

func (_this *AlidnsProvider) getCommonParamter(action *string) *url.Values {
	timestamp := time.Now().UTC().Format(time2.ALIYUN_FORMAT_ISO8601)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	params := make(url.Values, 10)
	params.Set("Action", strs.StringValue(action))
	params.Set("AccessKeyId", _this.credntial.GetSecretId())
	params.Set("Format", "JSON")
	params.Set("SignatureMethod", _this.compose.SignatureMethod())
	params.Set("SignatureNonce", nonce)
	params.Set("SignatureVersion", _this.compose.SignerVersion())
	params.Set("Timestamp", timestamp)
	params.Set("Version", "2015-01-09")
	return &params
}
