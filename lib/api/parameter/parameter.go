package parameter

import (
	"net/url"
	"vdns/lib/api/models"
)

type ParamaterProvider interface {
	// LoadDescribeParamater 加载Describe参数
	LoadDescribeParamater(request *models.DescribeDomainRecordsRequest, action *string) (*url.Values, error)

	// LoadCreateParamater 加载Create参数
	LoadCreateParamater(request *models.CreateDomainRecordRequest, action *string) (*url.Values, error)

	// LoadUpdateParamater 加载Update参数
	LoadUpdateParamater(request *models.UpdateDomainRecordRequest, action *string) (*url.Values, error)

	// LoadDeleteParamater 加载Delete参数
	LoadDeleteParamater(request *models.DeleteDomainRecordRequest, action *string) (*url.Values, error)
}
