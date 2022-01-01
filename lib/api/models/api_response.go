package models

import (
	"vdns/lib/standard/record"
	"vdns/vutil/vjson"
)

type DomainRecordResponse struct {
	TotalCount *int64    `json:"total_count,omitempty"`
	PageSize   *int64    `json:"page_size,omitempty"`
	PageNumber *int64    `json:"page_number,omitempty"`
	Records    []*Record `json:"records,omitempty"`
}

func (_this *DomainRecordResponse) String() string {
	return vjson.PrettifyString(_this)
}

type Record struct {
	ID         *string     `json:"id,omitempty"`
	RecordType record.Type `json:"record_type,omitempty"`
	RR         *string     `json:"rr,omitempty"`
	Domain     *string     `json:"domain,omitempty"`
	Value      *string     `json:"value,omitempty"`
	TTL        *int64      `json:"ttl,omitempty"`
}

func (_this *Record) String() string {
	return vjson.PrettifyString(_this)
}

type DomainStatusResponse struct {
	RequestId *string `json:"request_id,omitempty"`
	RecordId  *string `json:"record_id,omitempty"`
}

func (_this *DomainStatusResponse) String() string {
	return vjson.PrettifyString(_this)
}
