package models

import "vdns/lib/standard/record"

type DnsRecordResponse struct {
	ID         *string      `json:"id,omitempty"`
	RecordType *record.Type `json:"record_type,omitempty"`
	RR         *string      `json:"rr,omitempty"`
	Domain     *string      `json:"domain,omitempty"`
	Value      *string      `json:"value,omitempty"`
	TTL        *int64       `json:"ttl,omitempty"`
}
