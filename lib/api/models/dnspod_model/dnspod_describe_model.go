package dnspod_model

import (
	"github.com/zf1976/vdns/lib/util/vjson"
)

type DescribeRecordListResponse struct {
	Response *struct {

		// 记录的数量统计信息
		RecordCountInfo *RecordCountInfo `json:"RecordCountInfo,omitempty" name:"RecordCountInfo"`

		// 获取的记录列表
		RecordList []*RecordListItem `json:"RecordList,omitempty" name:"RecordList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

		// 返回错误信息
		Error *Error `json:"Error" name:"Error"`
	} `json:"Response"`
}

func (s *DescribeRecordListResponse) String() string {
	return vjson.PrettifyString(s)
}

type RecordCountInfo struct {

	// 子域名数量
	SubdomainCount *int64 `json:"SubdomainCount,omitempty" name:"SubdomainCount"`

	// 列表返回的记录数
	ListCount *int64 `json:"ListCount,omitempty" name:"ListCount"`

	// 总的记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

func (s *RecordCountInfo) String() string {
	return vjson.PrettifyString(s)
}

type RecordListItem struct {

	// 记录Id
	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`

	// 记录值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 记录状态，启用：ENABLE，暂停：DISABLE
	Status *string `json:"Status,omitempty" name:"Status"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// 主机名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 记录线路
	Line *string `json:"Line,omitempty" name:"Line"`

	// 线路Id
	LineId *string `json:"LineId,omitempty" name:"LineId"`

	// 记录类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 记录权重，用于负载均衡记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 记录监控状态，正常：OK，告警：WARN，宕机：DOWN，未设置监控或监控暂停则为空
	MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// 记录备注说明
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 记录缓存时间
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// MX值，只有MX记录有
	// 注意：此字段可能返回 null，表示取不到有效值。
	MX *int64 `json:"MX,omitempty" name:"MX"`
}

func (s *RecordListItem) String() string {
	return vjson.PrettifyString(s)
}
