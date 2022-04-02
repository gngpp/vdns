package config

// provides
//goland:noinspection GoUnusedConst,GoSnakeCaseUsage,SpellCheckingInspection
const (
	ALIDNS_PROVIDER      = "AliDNS"
	DNSPOD_PROVIDER      = "DNSPod"
	CLOUDFLARE_PROVIDER  = "Cloudflare"
	HUAWERI_DNS_PROVIDER = "HuaweiDNS"
)

// ipv4 api endpoint
//goland:noinspection GoUnusedConst
const (
	V4Api1 = "https://ddns.oray.com/checkip"
	V4Api2 = "https://api-ipv4.ip.sb/ip"
	V4Api3 = "https://api.ip.sb/ip"
	V4Api4 = "https://myip.ipip.net"
	V4Api5 = "https://api-ipv4.ip.sb/ip"
)

// ipv6 api endpoint
//goland:noinspection GoUnusedConst
const (
	V6Api1 = "https://api-ipv6.ip.sb/ip"
	V6Api2 = "https://v6.myip.la/json"
	V6Api3 = "https://speed.neu6.edu.cn/getIP.php"
)
