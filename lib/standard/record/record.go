package record

import "vdns/util/strs"

type Type string

func (_this Type) String() string {
	//goland:noinspection GoRedundantConversion
	return string(_this)
}

//goland:noinspection GoSnakeCaseUsage
const PAN_ANALYSIS_RR_KEY_WORD Type = "@"

// reference https://help.aliyun.com/document_detail/29805.html?spm=a2c4g.11186623.0.0.30e73067AXxwak
// 类型名称		类型取值
// A记录			A
// NS记录		NS
// MX记录		MX
// TXT记录		TXT
// CNAME记录		CNAME
// SRV记录		SRV
// AAAA记录		AAAA
// CAA记录		CAA
// 显性URL转发 	REDIRECT_URL
// 隐性URL转发 	FORWARD_URL
const (
	a           Type = "A"            // 将域名指向一个IPV4地址
	aaaa        Type = "AAAA"         // 将域名指向一个IPV6地址
	ns          Type = "NS"           // 将子域名指定其他DNS服务器解析
	mx          Type = "MX"           // 将域名指向邮件服务器地址
	cname       Type = "CNAME"        // 将域名指向另外一个域名
	txt         Type = "TXT"          // 文本长度限制512，通常做SPF记录（反垃圾邮件）
	srv         Type = "SRV"          // 记录提供特定的服务的服务器
	caa         Type = "CAA"          // CA证书颁发机构授权校验
	redirectUrl Type = "REDIRECT_URL" // 将域名重定向到另外一个地址
	forwardUrl  Type = "FORWARD_URL"  // 与显性URL类似，但是会隐藏真实目标地址
)

var typeMap map[Type]Type

//goland:noinspection ALL
const (
	A            = a
	AAAA         = aaaa
	NS           = ns
	MX           = mx
	CNAME        = cname
	TXT          = txt
	SRV          = srv
	CAA          = caa
	REDIRECT_URL = redirectUrl
	FORWARD_URL  = forwardUrl
)

func OfType(t Type) (Type, bool) {
	recordType, isOk := typeMap[t]
	return recordType, isOk
}

func Support(t Type) bool {
	_, isOk := OfType(t)
	// allow to be empty
	return isOk || strs.IsEmpty(t.String())
}

func init() {
	typeMap = map[Type]Type{
		a:           A,
		aaaa:        AAAA,
		ns:          NS,
		mx:          MX,
		cname:       CNAME,
		txt:         TXT,
		srv:         SRV,
		caa:         CAA,
		redirectUrl: REDIRECT_URL,
		forwardUrl:  FORWARD_URL,
	}
}
