package models

type BaseRequest struct {
	httpMethod string
	scheme     string
	rootDomain string
	domain     string
	path       string
	params     map[string]string
	formParams map[string]string

	service string
	version string
	action  string

	contentType string
	body        []byte
}

type DescribeDomainListRequest struct {
	*BaseRequest

	// 域名分组类型，默认为ALL。可取值为ALL，MINE，SHARE，ISMARK，PAUSE，VIP，RECENT，SHARE_OUT。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 记录开始的偏移, 第一条记录为 0, 依次类推。默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 要获取的域名数量, 比如获取20个, 则为20。默认值为3000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分组ID, 获取指定分组的域名
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 根据关键字搜索域名
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}
