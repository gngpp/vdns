package aliyun_model

import (
	"github.com/alibabacloud-go/tea/tea"
	"vdns/vutil/vjson"
)

// AliyunDescribeDomainRecordsResponseBody aliyun response model
type AliyunDescribeDomainRecordsResponseBody struct {
	TotalCount    *int64                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize      *int64                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainRecords *AliyunDescribeDomainRecordsResponseBodyDomainRecords `json:"DomainRecords,omitempty" xml:"DomainRecords,omitempty" type:"Struct"`
	PageNumber    *int64                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s *AliyunDescribeDomainRecordsResponseBody) String() string {
	return vjson.PrettifyString(s)
}

type AliyunDescribeDomainRecordsResponseBodyDomainRecords struct {
	Record []*AliyunDescribeDomainRecordsResponseBodyDomainRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s *AliyunDescribeDomainRecordsResponseBodyDomainRecords) String() string {
	return vjson.PrettifyString(s)
}

type AliyunDescribeDomainRecordsResponseBodyDomainRecordsRecord struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TTL        *int64  `json:"TTL,omitempty" xml:"TTL,omitempty"`
	RecordId   *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Priority   *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RR         *string `json:"RR,omitempty" xml:"RR,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Weight     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Line       *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Locked     *bool   `json:"Locked,omitempty" xml:"Locked,omitempty"`
}

func (s *AliyunDescribeDomainRecordsResponseBodyDomainRecordsRecord) String() string {
	return tea.Prettify(s)
}

type EResponse struct {
	RequestId string `json:"RequestId"`
	RecordId  string `json:"RecordId"`
}
