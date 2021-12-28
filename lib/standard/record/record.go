package record

type Type string

//goland:noinspection ALL
const (
	a           Type = "A"     // 将域名指向一个IPV4地址
	aaaa        Type = "AAAA"  // 将域名指向一个IPV6地址
	ns          Type = "NS"    // 将子域名指定其他DNS服务器解析
	mx          Type = "MX"    // 将域名指向邮件服务器地址
	cname       Type = "CNAME" // 将域名指向另外一个域名
	txt         Type = "TXT"   // 文本长度限制512，通常做SPF记录（反垃圾邮件）
	srv         Type = "SRV"   // 记录提供特定的服务的服务器
	ca          Type = "CA"    // CA证书颁发机构授权校验
	explicitUrl Type = "显性URL" // 将域名重定向到另外一个地址
	hiddenUrl   Type = "隐性URL" // 与显性URL类似，但是会隐藏真实目标地址
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
	CA           = ca
	EXPLICIT_URL = explicitUrl
	HIDDEN_URL   = hiddenUrl
)

func OfType(t Type) (Type, bool) {
	recordType, isOk := typeMap[t]
	return recordType, isOk
}

func Support(t Type) bool {
	_, isOk := OfType(t)
	return isOk
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
		ca:          CA,
		explicitUrl: EXPLICIT_URL,
		hiddenUrl:   HIDDEN_URL,
	}
}
