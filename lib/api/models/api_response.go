package models

import (
	"github.com/zf1976/vdns/lib/standard/record"
	"github.com/zf1976/vdns/lib/util/vjson"
)

type DomainRecordsResponse struct {
	TotalCount *int64    `json:"total_count,omitempty"`
	PageSize   *int64    `json:"page_size,omitempty"`
	PageNumber *int64    `json:"page_number,omitempty"`
	ListCount  *int64    `json:"list_count,omitempty"`
	Records    []*Record `json:"records,omitempty"`
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
