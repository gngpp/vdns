package record

type Type string

func (_this Type) String() string {
	//goland:noinspection GoRedundantConversion
	return string(_this)
}

// 类型名称	类型取值	类型定义	类型描述
// A记录		A	参考标准；RR值可为空，即@解析；不允许含有下划线；	IPv4地址格式
// NS记录	NS	参考标准；RR值不能为空；允许含有下划线；不支持泛解析	NameType形式
// MX记录	MX	参考标准；RR值可为空，即@解析；不允许含有下划线	NameType形式，且不可为IP地址。1-10，优先级依次递减。
// TXT记录	TXT	参考标准；另外，有效字符除字母、数字、“-”（中横杠）、还包括“_”(下划线)；RR值可为空，即@解析；允许含有下划线；不支持泛解析	字符串；长度小于512,合法字符：大小写字母，数字,空格，及以下字符：-~=:;/.@+^!*
// CNAME记录	CNAME	参考标准；另外，有效字符除字母、数字、“-”（中横杠）、还包括“_”(下划线)；RR值不允许为空（即@）；允许含有下划线	NameType形式，且不可为IP
// SRV记录	SRV	是一个name，且可含有下划线“_“和点“.”；允许含有下划线；可为空（即@）；不支持泛解析	priority:优先级，为0－65535之间的数字；weight：权重，为0－65535之间的数字；port：提供服务的端口号，为0－65535之间的数字 target：为提供服务的目标地址，为nameType，且存在。参考：http://en.wikipedia.org/wiki/SRV_record http://www.rfc-editor.org/rfc/rfc2782.txt
// AAAA记录	AAAA	参考标准；RR值可为空，即@解析；不允许含有下划线；	IPv6地址格式
// CAA记录	CAA	参考标准；RR值可为空，即@解析；不允许含有下划线；	格式为：[flag] [tag] [value]，是由一个标志字节的[flag],和一个被称为属性的标签[tag]-值[value]对组成。例如：@ 0 issue "symantec.com"或@ 0 iodef "mailto:admin@aliyun.com"
// 显性URL转发 REDIRECT_URL	参考标准；RR值可为空，即@解析	NameType或URL地址（区分大小写），长度最长为500字符，其中域名，如a.com，必须，大小写不敏感；协议：可选，如http、https，默认为http端口：可选，如81，默认为80；路径：可选，大小写敏感，如/path/to/，默认为/；文件名：可选，大小写敏感，如file.php，默认无；参数：可选，大小写敏感，如?user=my***，默认无。
// 隐性URL转发 FORWARD_URL	参考标准；RR值可为空，即@解析	NameType或URL地址（区分大小写），长度最长为500字符，其中域名，如a.com，必须，大小写不敏感；协议：可选，如http、https，默认为http端口：可选，如81，默认为80；路径：可选，大小写敏感，如/path/to/，默认为/；文件名：可选，大小写敏感，如file.php，默认无；参数：可选，大小写敏感，如?user=my***，默认无。
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
		caa:         CAA,
		redirectUrl: REDIRECT_URL,
		forwardUrl:  FORWARD_URL,
	}
}
