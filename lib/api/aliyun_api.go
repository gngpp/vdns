package api

import (
	"net/url"
	"strconv"
	"time"
	"vdns/lib/auth"
	"vdns/lib/sign/rpc"
	"vdns/lib/standard"
	"vdns/lib/standard/record"
	time2 "vdns/lib/standard/time"
	"vdns/vutil/str"
	"vdns/vutil/vhttp"
)

func NewAliyunDnsProvider(credential *auth.BasisCredential) DnsProvider {
	return &AliyunDnsProvider{
		Action: &Action{
			describe: "DescribeDomainRecords",
			create:   "AddDomainRecord",
			update:   "UpdateDomainRecord",
			delete:   "DeleteDomainRecord",
		},
		compose: &rpc.AliyunRpcSignatureCompose{
			Separator: rpc.SEPARATOR,
		},
		api:       standard.ALIYUN_DNS_API,
		credntial: credential,
	}
}

type AliyunDnsProvider struct {
	*Action
	api       string
	compose   rpc.RpcSignatureComposer
	credntial auth.Credential
}

func (_this *AliyunDnsProvider) DescribeRecordList(request DnsRecordRequest) (*DnsRecordResponse, error) {
	panic("implement me")
}

func (_this *AliyunDnsProvider) CreateDnsRecord(request DnsRecordRequest) (*DnsRecordResponse, error) {
	panic("implement me")
}

func (_this *AliyunDnsProvider) UpdateDnsRecord(request DnsRecordRequest) (*DnsRecordResponse, error) {
	panic("implement me")
}

func (_this *AliyunDnsProvider) DeleteDnsRecord(request DnsRecordRequest) (*DnsRecordResponse, error) {
	panic("implement me")
}

func (_this *AliyunDnsProvider) Support(recordType record.Type) bool {
	return record.Support(string(recordType))
}

func (_this *AliyunDnsProvider) getParam(request DnsRecordRequest, action string) error {
	domain, err := vhttp.ExtractDomain(request.Domain)
	if err != nil {
		return err
	}
	timestamp := time.Now().UTC().Format(time2.ALIYUN_FORMAT_ISO8601)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	params := make(url.Values, 10)
	params.Set("Action", action)
	params.Set("AccessKeyId", _this.credntial.GetSecretId())
	params.Set("Format", "JSON")
	params.Set("SignatureMethod", _this.compose.SignatureMethod())
	params.Set("SignatureNonce", nonce)
	params.Set("SignatureVersion", _this.compose.SignerVersion())
	params.Set("Timestamp", timestamp)
	params.Set("Version", "2015-01-09")

	switch action {
	case "DescribeDomainRecords":
		{
			params.Set("PageSize", strconv.FormatInt(*request.PageSize, 10))
			params.Set("TypeKeyWord", string(request.RecordType))
			params.Set("DomainName", domain[0])
			if !str.IsEmpty(domain[1]) {
				params.Set("RRKeyWord", domain[1])
			}
		}
	case "AddDomainRecord":

	case "UpdateDomainRecord":

	case "DeleteDomainRecord":
	}

	return nil
}
