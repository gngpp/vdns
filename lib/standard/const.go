package standard

type Standard string

func (_this Standard) StringValue() string {
	return string(_this)
}

func (_this Standard) String() *Standard {
	var v = Standard(_this.StringValue())
	return &v
}

//goland:noinspection ALL
const (
	ALIYUN_DNS_API     Standard = "https://alidns.aliyuncs.com/"
	DNSPOD_DNS_API     Standard = "https://dnspod.tencentcloudapi.com/"
	HUAWEI_DNS_API     Standard = "https://dns.myhuaweicloud.com/v2/zones"
	CLOUDFLARE_DNS_API Standard = "https://api.cloudflare.com/client/v4/zones"
)

//goland:noinspection ALL
const (
	ALIYUN_ENDPOINT_DOMAIN     Standard = "alidns.aliyuncs.com"
	DNSPOD_ENDPOINT_DOMAIN     Standard = "dnspod.tencentcloudapi.com"
	HUAWEI_ENDPOINT_DOMAIN     Standard = "dns.myhuaweicloud.com"
	CLOUDFLARE_ENDPOINT_DOMAIN Standard = "api.cloudflare.com"
)

// cloud api version
//goland:noinspection ALL
const (
	ALIDNS_API_VERSION Standard = "2015-01-09"
	DNSPOD_API_VERSION Standard = "2021-03-23"
)
