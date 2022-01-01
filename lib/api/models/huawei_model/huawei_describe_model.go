package huawei_model

import (
	"vdns/vutil/vjson"
)

// ListRecordSetsResponse Response Object
type ListRecordSetsResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Recordsets *[]ListRecordSetsWithTags `json:"recordsets,omitempty"`

	Metadata       *Metedata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o *ListRecordSetsResponse) String() string {
	return vjson.PrettifyString(o)
}

// PageLink 分页信息
type PageLink struct {
	// 当前资源的链接。

	Self *string `json:"self,omitempty"`
	// 下一页资源的链接。

	Next *string `json:"next,omitempty"`
}

func (o *PageLink) String() string {
	return vjson.PrettifyString(o)
}

type ListRecordSetsWithTags struct {
	// Record Set的ID。

	Id *string `json:"id,omitempty"`
	// Record Set的名称。

	Name *string `json:"name,omitempty"`
	// Record Set的描述信息。

	Description *string `json:"description,omitempty"`
	// 托管该记录的zone_id。

	ZoneId *string `json:"zone_id,omitempty"`
	// 托管该记录的zone_name。

	ZoneName *string `json:"zone_name,omitempty"`
	// 记录类型。

	Type *string `json:"type,omitempty"`
	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。

	Ttl *int32 `json:"ttl,omitempty"`
	// 域名解析后的值。

	Records *[]string `json:"records,omitempty"`
	// 创建时间。

	CreateAt *string `json:"create_at,omitempty"`
	// 更新时间。

	UpdateAt *string `json:"update_at,omitempty"`
	// 资源状态。

	Status *string `json:"status,omitempty"`
	// 标识是否由系统默认生成，系统默认生成的Record Set不能删除。

	Default *bool `json:"default,omitempty"`
	// 该Record Set所属的项目ID。

	ProjectId *string `json:"project_id,omitempty"`

	Links *PageLink `json:"links,omitempty"`
	// 资源标签。

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o *ListRecordSetsWithTags) String() string {
	return vjson.PrettifyString(o)
}

type Metedata struct {
	// 满足查询条件的资源总数，不受分页（即limit、offset参数）影响。

	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o *Metedata) String() string {
	return vjson.PrettifyString(o)
}

type Tag struct {
	// 键。最大长度36个unicode字符。 key不能为空。不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Key string `json:"key"`
	// 值。每个值最大长度43个unicode字符，可以为空字符串。 不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。

	Value *string `json:"value,omitempty"`
}

func (o *Tag) String() string {
	return vjson.PrettifyString(o)
}
