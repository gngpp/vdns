package model

import (
	"vdns/lib/standard/record"
	"vdns/lib/util/vjson"
)

type DomainRecordsResponse struct {
	// 总数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 页大小
	PageSize *int64 `json:"page_size,omitempty"`
	// 当前页
	PageNumber *int64 `json:"page_number,omitempty"`
	// 列表总数
	ListCount *int64 `json:"list_count,omitempty"`
	// 记录列表
	Records []*Record `json:"records,omitempty"`
}

func (_this *DomainRecordsResponse) String() string {
	return vjson.PrettifyString(_this)
}

type Record struct {
	ID         *string     `json:"id,omitempty"`
	RecordType record.Type `json:"record_type,omitempty"`
	RR         *string     `json:"rr,omitempty"`
	Domain     *string     `json:"domain,omitempty"`
	Value      *string     `json:"value,omitempty"`
	Status     *string     `json:"status,omitempty"`
	TTL        *int64      `json:"ttl,omitempty"`
}

func (_this *Record) String() string {
	return vjson.PrettifyString(_this)
}

type DomainRecordStatusResponse struct {
	RequestId *string `json:"request_id,omitempty"`
	RecordId  *string `json:"record_id,omitempty"`
}

func (_this *DomainRecordStatusResponse) String() string {
	return vjson.PrettifyString(_this)
}
