package parameter

import (
	"net/url"
	"vdns/lib/api/model"
)

type Parameter interface {
	// LoadDescribeParameter 加载Describe参数
	LoadDescribeParameter(request *model.DescribeDomainRecordsRequest, action *string) (*Value, error)

	// LoadCreateParameter 加载Create参数
	LoadCreateParameter(request *model.CreateDomainRecordRequest, action *string) (*Value, error)

	// LoadUpdateParameter 加载Update参数
	LoadUpdateParameter(request *model.UpdateDomainRecordRequest, action *string) (*Value, error)

	// LoadDeleteParameter 加载Delete参数
	LoadDeleteParameter(request *model.DeleteDomainRecordRequest, action *string) (*Value, error)
}

type Value struct {
	UrlValues  *url.Values
	JsonString any
}

func NewValue(values *url.Values, json any) *Value {
	return &Value{
		UrlValues:  values,
		JsonString: json,
	}
}
