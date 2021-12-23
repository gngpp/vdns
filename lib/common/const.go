package common

type api string

//goland:noinspection ALL
const (
	ALIYUN_DNS_API     api = "https://alidns.aliyuncs.com/"
	DNSPOD_DNS_API     api = "https://dnspod.tencentcloudapi.com/"
	HUAWEI_DNS_API     api = "https://dns.myhuaweicloud.com/v2/zones"
	CLOUDFLARE_DNS_API api = "https://api.cloudflare.com/client/v4/zones"
)
