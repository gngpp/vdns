package api

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"vdns/lib/api/conv"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	"vdns/lib/api/models/aliyun_model"
	"vdns/lib/auth"
	"vdns/lib/sign/rpc"
	"vdns/lib/standard"
	"vdns/lib/standard/msg"
	"vdns/lib/standard/record"
	time2 "vdns/lib/standard/time"
	"vdns/vutil/str"
	"vdns/vutil/vhttp"
)

func NewAliyunDnsProvider(credential auth.Credential) DnsProvider {
	return &AliyunDnsProvider{
		Action: &Action{
			describe: "DescribeDomainRecords",
			create:   "AddDomainRecord",
			update:   "UpdateDomainRecord",
			delete:   "DeleteDomainRecord",
		},
		compose:   rpc.NewAliyunRpcSignatureCompose(rpc.SEPARATOR),
		api:       str.String(standard.ALIYUN_DNS_API),
		credntial: credential,
	}
}

type AliyunDnsProvider struct {
	*Action
	api       *string
	compose   rpc.RpcSignatureComposer
	credntial auth.Credential
}

func (_this *AliyunDnsProvider) DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.ApiDomainRecordResponse, error) {
	paramater, err := _this.getDescribeParamater(request, &_this.describe)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.doRequest(requestUrl)
}

func (_this *AliyunDnsProvider) CreateDnsRecord(request *models.CreateDomainRecordRequest) (*models.ApiDomainRecordResponse, error) {
	paramater, err := _this.getCreateParamater(request, &_this.create)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.doRequest(requestUrl)
}

func (_this *AliyunDnsProvider) UpdateDnsRecord(request *models.UpdateDomainRecordRequest) (*models.ApiDomainRecordResponse, error) {
	panic("implement me")
}

func (_this *AliyunDnsProvider) DeleteDnsRecord(request *models.DeleteDomainRecordRequest) (*models.ApiDomainRecordResponse, error) {
	panic("implement me")
}

func (_this *AliyunDnsProvider) Support(recordType record.Type) bool {
	return record.Support(recordType)
}

func (_this *AliyunDnsProvider) doRequest(url string) (*models.ApiDomainRecordResponse, error) {
	resp, _ := http.Get(url)
	if resp.StatusCode == http.StatusOK {
		bytes, _ := ioutil.ReadAll(resp.Body)
		var reponse = &aliyun_model.AliyunDescribeDomainRecordsResponseBody{}
		err := json.Unmarshal(bytes, reponse)
		if err != nil {
			return new(models.ApiDomainRecordResponse), err
		}
		return conv.AiiyunBodyToResponse(reponse), nil
	} else {
		err := errs.NewAliyunSDKError(nil)
		bytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bytes, &err)
		return nil, err
	}
}

func (_this *AliyunDnsProvider) generateRequestUrl(paramater *url.Values) string {
	stringToSign := _this.compose.ComposeStringToSign(vhttp.HttpMethodGet, paramater)
	signature := _this.compose.GeneratedSignature(_this.credntial.GetSecretKey(), stringToSign)
	_this.SetSignature(signature, paramater)
	return _this.compose.CanonicalizeRequestUrl(str.StringValue(_this.api), paramater)
}

func (_this *AliyunDnsProvider) SetSignature(signature string, queries *url.Values) {
	queries.Set("Signature", base64.StdEncoding.EncodeToString(str.ToBytes(signature)))
}

func (_this *AliyunDnsProvider) getDescribeParamater(request *models.DescribeDomainRecordsRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.DESCRIBE_REQUEST_NOT_NIL)
	}
	extractDomain, err := vhttp.ExtractDomain(str.StringValue(request.Domain))
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
	if condition && str.NotEmpty(rr) {
		paramter.Set("RRKeyWord", rr)
	}
	return paramter, nil
}

func (_this *AliyunDnsProvider) getCreateParamater(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.CREATE_REQUEST_NOT_NIL)
	}
	if !record.Support(request.RecordType) {
		return nil, errs.NewApiError(msg.RECORD_TYPE_NOT_SUPPORT)

	}
	if request.Value == nil {
		return nil, errs.NewApiError(msg.RECORD_VALUE_NOT_SUPPORT)
	}
	extractDomain, err := vhttp.ExtractDomain(str.StringValue(request.Domain))
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

func (_this *AliyunDnsProvider) getUpdateParamater(request *models.UpdateDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.CREATE_REQUEST_NOT_NIL)
	}
	if request.ID != nil {
		return nil, errs.NewApiError(msg.RECORD_ID_NOT_SUPPORT)
	}
	if !record.Support(request.RecordType) {
		return nil, errs.NewApiError(msg.RECORD_TYPE_NOT_SUPPORT)

	}
	if request.Value == nil {
		return nil, errs.NewApiError(msg.RECORD_VALUE_NOT_SUPPORT)
	}
	extractDomain, err := vhttp.ExtractDomain(str.StringValue(request.Domain))
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

func (_this *AliyunDnsProvider) getDeleteParamater(request *models.DeleteDomainRecordRequest, action *string) (*url.Values, error) {
	if request == nil {
		return nil, errs.NewApiError(msg.CREATE_REQUEST_NOT_NIL)
	}
	paramter := _this.getCommonParamter(action)
	if request.ID != nil {
		return nil, errs.NewApiError(msg.RECORD_ID_NOT_SUPPORT)
	}
	paramter.Set("RecordId", *request.ID)
	return paramter, nil
}

func (_this *AliyunDnsProvider) getCommonParamter(action *string) *url.Values {
	timestamp := time.Now().UTC().Format(time2.ALIYUN_FORMAT_ISO8601)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	params := make(url.Values, 10)
	params.Set("Action", *action)
	params.Set("AccessKeyId", _this.credntial.GetSecretId())
	params.Set("Format", "JSON")
	params.Set("SignatureMethod", _this.compose.SignatureMethod())
	params.Set("SignatureNonce", nonce)
	params.Set("SignatureVersion", _this.compose.SignerVersion())
	params.Set("Timestamp", timestamp)
	params.Set("Version", "2015-01-09")
	return &params
}
