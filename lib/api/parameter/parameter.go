package parameter

import (
	"net/url"
	"vdns/lib/api/models"
)

type ParamaterProvider interface {
	// LoadDescribeParameter 加载Describe参数
	LoadDescribeParameter(request *models.DescribeDomainRecordsRequest, action *string) (*url.Values, error)

	// LoadCreateParameter 加载Create参数
	LoadCreateParameter(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error)

	// LoadUpdateParameter 加载Update参数
	LoadUpdateParameter(request *models.UpdateDomainRecordRequest, action *string) (*url.Values, error)

	// LoadDeleteParameter 加载Delete参数
	LoadDeleteParameter(request *models.DeleteDomainRecordRequest, action *string) (*url.Values, error)
}
