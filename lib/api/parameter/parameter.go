package parameter

import (
	"net/url"
	"vdns/lib/api/model"
)

type Parameter interface {
	// LoadDescribeParameter 加载Describe参数
	LoadDescribeParameter(request *model.DescribeDomainRecordsRequest, action *string) (*url.Values, error)

	// LoadCreateParameter 加载Create参数
	LoadCreateParameter(request *model.CreateDomainRecordRequest, action *string) (*url.Values, error)

	// LoadUpdateParameter 加载Update参数
	LoadUpdateParameter(request *model.UpdateDomainRecordRequest, action *string) (*url.Values, error)

	// LoadDeleteParameter 加载Delete参数
	LoadDeleteParameter(request *model.DeleteDomainRecordRequest, action *string) (*url.Values, error)
}
