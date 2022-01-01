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
	"vdns/vlog"
	"vdns/vutil/str"
	"vdns/vutil/vhttp"
	"vdns/vutil/vjson"
)

func NewAliyunDnsProvider(credential auth.Credential) DNSRecordProvider {
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

func (_this *AliyunDnsProvider) DescribeRecords(request *models.DescribeDomainRecordsRequest) (*models.DomainRecordResponse, error) {
	paramater, err := _this.loadDescribeParamater(request, &_this.describe)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.doDescribeRequest(requestUrl)
}

func (_this *AliyunDnsProvider) CreateRecord(request *models.CreateDomainRecordRequest) (*models.DomainStatusResponse, error) {
	paramater, err := _this.loadCreateParamater(request, &_this.create)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.doCreateRequest(requestUrl)
}

func (_this *AliyunDnsProvider) UpdateRecord(request *models.UpdateDomainRecordRequest) (*models.DomainStatusResponse, error) {
	paramater, err := _this.loadUpdateParamater(request, &_this.update)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.doCreateRequest(requestUrl)
}

func (_this *AliyunDnsProvider) DeleteRecord(request *models.DeleteDomainRecordRequest) (*models.DomainStatusResponse, error) {
	paramater, err := _this.loadDeleteParamater(request, &_this.delete)
	if err != nil {
		return nil, err
	}
	requestUrl := _this.generateRequestUrl(paramater)
	return _this.doCreateRequest(requestUrl)
}

func (_this *AliyunDnsProvider) Support(recordType record.Type) bool {
	return record.Support(recordType)
}

func (_this *AliyunDnsProvider) doDescribeRequest(url string) (*models.DomainRecordResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainRecordResponse{}, err
	}
	if resp.StatusCode == http.StatusOK {
		bytes, _ := ioutil.ReadAll(resp.Body)
		var reponse = &aliyun_model.DescribeDomainRecordsResponse{}
		err := json.Unmarshal(bytes, reponse)
		if err != nil {
			return new(models.DomainRecordResponse), err
		}
		return conv.AiiyunBodyToResponse(reponse), nil
	} else {
		sdkError := errs.NewAliyunSDKError(nil)
		bytes, _ := ioutil.ReadAll(resp.Body)
		vjson.ByteArrayConver(bytes, sdkError)
		return nil, sdkError
	}
}

func (_this *AliyunDnsProvider) doCreateRequest(url string) (*models.DomainStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainStatusResponse{}, err
	}
	if resp.StatusCode == http.StatusOK {
		bytes, _ := ioutil.ReadAll(resp.Body)
		response := &models.DomainStatusResponse{}
		err := vjson.ByteArrayConver(bytes, response)
		if err != nil {
			vlog.Default().Fatal(err)
		}
		return response, nil
	} else {
		sdkError := errs.NewAliyunSDKError(nil)
		bytes, _ := ioutil.ReadAll(resp.Body)
		vjson.ByteArrayConver(bytes, sdkError)
		return nil, sdkError
	}
}

func (_this *AliyunDnsProvider) doUpdateRequest(url string) (*models.DomainStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainStatusResponse{}, err
	}
	if resp.StatusCode == http.StatusOK {
		bytes, _ := ioutil.ReadAll(resp.Body)
		response := &models.DomainStatusResponse{}
		err := vjson.ByteArrayConver(bytes, response)
		if err != nil {
			vlog.Default().Fatal(err)
		}
		return response, nil
	} else {
		sdkError := errs.NewAliyunSDKError(nil)
		bytes, _ := ioutil.ReadAll(resp.Body)
		vjson.ByteArrayConver(bytes, sdkError)
		return nil, sdkError
	}
}

func (_this *AliyunDnsProvider) doDeleteRequest(url string) (*models.DomainStatusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &models.DomainStatusResponse{}, err
	}
	if resp.StatusCode == http.StatusOK {
		bytes, _ := ioutil.ReadAll(resp.Body)
		response := &models.DomainStatusResponse{}
		err := vjson.ByteArrayConver(bytes, response)
		if err != nil {
			vlog.Default().Fatal(err)
		}
		return response, nil
	} else {
		sdkError := errs.NewAliyunSDKError(nil)
		bytes, _ := ioutil.ReadAll(resp.Body)
		vjson.ByteArrayConver(bytes, sdkError)
		return nil, sdkError
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

func (_this *AliyunDnsProvider) loadDescribeParamater(request *models.DescribeDomainRecordsRequest, action *string) (*url.Values, error) {
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

func (_this *AliyunDnsProvider) loadCreateParamater(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error) {
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

func (_this *AliyunDnsProvider) loadUpdateParamater(request *models.UpdateDomainRecordRequest, action *string) (*url.Values, error) {
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

func (_this *AliyunDnsProvider) loadDeleteParamater(request *models.DeleteDomainRecordRequest, action *string) (*url.Values, error) {
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
