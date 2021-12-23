package common

import (
	"reflect"
)

type RecordType string

//goland:noinspection ALL
const (
	a           RecordType = "A"     // 将域名指向一个IPV4地址
	aaaa        RecordType = "AAAA"  // 将域名指向一个IPV6地址
	ns          RecordType = "NS"    // 将子域名指定其他DNS服务器解析
	mx          RecordType = "MX"    // 将域名指向邮件服务器地址
	canme       RecordType = "CNAME" // 将域名指向另外一个域名
	txt         RecordType = "TXT"   // 文本长度限制512，通常做SPF记录（反垃圾邮件）
	srv         RecordType = "SRV"   // 记录提供特定的服务的服务器
	ca          RecordType = "CA"    // CA证书颁发机构授权校验
	explicitUrl RecordType = "显性URL" // 将域名重定向到另外一个地址
	hiddenUrl   RecordType = "隐性URL" // 与显性URL类似，但是会隐藏真实目标地址
)

var Record *RecordEnum
var recordMap map[string]RecordType

//goland:noinspection ALL
type RecordEnum struct {
	A            RecordType
	AAAA         RecordType
	NS           RecordType
	MX           RecordType
	CNAME        RecordType
	TXT          RecordType
	SRV          RecordType
	CA           RecordType
	EXPLICIT_URL RecordType
	HIDDEN_URL   RecordType
}

func (t RecordEnum) OfType(value string) bool {
	_, isOk := recordMap[value]
	return isOk
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
	enumMap := make(map[string]RecordType)
	for i := 0; i < of.NumField(); i++ {
		structFieldName := of.Field(i).Name
		if _, b := of.FieldByName(structFieldName); b {
			enumMap[structFieldName] = RecordType(structFieldName)
		}
	}
	recordMap = enumMap
}
