package record

import (
	"reflect"
)

type Type string

//goland:noinspection ALL
const (
	a           Type = "A"     // 将域名指向一个IPV4地址
	aaaa        Type = "AAAA"  // 将域名指向一个IPV6地址
	ns          Type = "NS"    // 将子域名指定其他DNS服务器解析
	mx          Type = "MX"    // 将域名指向邮件服务器地址
	canme       Type = "CNAME" // 将域名指向另外一个域名
	txt         Type = "TXT"   // 文本长度限制512，通常做SPF记录（反垃圾邮件）
	srv         Type = "SRV"   // 记录提供特定的服务的服务器
	ca          Type = "CA"    // CA证书颁发机构授权校验
	explicitUrl Type = "显性URL" // 将域名重定向到另外一个地址
	hiddenUrl   Type = "隐性URL" // 与显性URL类似，但是会隐藏真实目标地址
)

var Record *RecordEnum
var recordMap map[string]Type

//goland:noinspection ALL
type RecordEnum struct {
	A            Type
	AAAA         Type
	NS           Type
	MX           Type
	CNAME        Type
	TXT          Type
	SRV          Type
	CA           Type
	EXPLICIT_URL Type
	HIDDEN_URL   Type
}

func OfType(value string) (bool, Type) {
	recordType, isOk := recordMap[value]
	return isOk, recordType
}

func Support(value string) bool {
	_, ok := OfType(value)
	return ok
}

func init() {
	Record = &RecordEnum{
		A:            a,
		AAAA:         aaaa,
		NS:           ns,
		MX:           mx,
		CNAME:        canme,
		TXT:          txt,
		SRV:          srv,
		CA:           ca,
		EXPLICIT_URL: explicitUrl,
		HIDDEN_URL:   hiddenUrl,
	}
	of := reflect.TypeOf(Record)
	if reflect.Ptr == of.Kind() {
		of = of.Elem()
	}
	enumMap := make(map[string]Type)
	for i := 0; i < of.NumField(); i++ {
		structFieldName := of.Field(i).Name
		if _, b := of.FieldByName(structFieldName); b {
			enumMap[structFieldName] = Type(structFieldName)
		}
	}
	recordMap = enumMap
}
