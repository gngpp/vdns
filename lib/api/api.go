package api

import (
	"vdns/lib/standard/record"
)

type Action struct {
	create   string
	update   string
	describe string
	delete   string
}

type DnsRecordResponse struct {
	ID         string      `json:"id,omitempty"`
	RecordType record.Type `json:"record_type,omitempty"`
	RR         string      `json:"rr,omitempty"`
	Domain     string      `json:"domain,omitempty"`
	Value      string      `json:"value,omitempty"`
	TTL        *int64      `json:"ttl,omitempty"`
}

type DnsRecordRequest struct {
	ID         string      `json:"id,omitempty"`
	Domain     string      `json:"domain,omitempty"`
	Value      string      `json:"value,omitempty"`
	RecordType record.Type `json:"record_type,omitempty"`
	PageSize   *int64      `json:"page_size,omitempty"`
}

type DnsProvider interface {
	// DescribeRecordList 具体参数作用请看实现注释
	DescribeRecordList(request DnsRecordRequest) (*DnsRecordResponse, error)

	// CreateDnsRecord 具体参数作用请看实现注释
	CreateDnsRecord(request DnsRecordRequest) (*DnsRecordResponse, error)

	// UpdateDnsRecord 具体参数作用请看实现注释
	UpdateDnsRecord(request DnsRecordRequest) (*DnsRecordResponse, error)

	// DeleteDnsRecord 具体参数作用请看实现注释
	DeleteDnsRecord(request DnsRecordRequest) (*DnsRecordResponse, error)

	// Support 某些使用zone区域划分域名记录的DNS服务商，需强迫使用support
	Support(recordType record.Type) bool
}
