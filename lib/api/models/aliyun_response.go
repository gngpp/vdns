package models

import "github.com/alibabacloud-go/tea/tea"

// DescribeDomainRecordsResponseBody aliyun response model
type DescribeDomainRecordsResponseBody struct {
	TotalCount    *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize      *int64                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainRecords *DescribeDomainRecordsResponseBodyDomainRecords `json:"DomainRecords,omitempty" xml:"DomainRecords,omitempty" type:"Struct"`
	PageNumber    *int64                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDomainRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsResponseBody) SetTotalCount(v int64) *DescribeDomainRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetPageSize(v int64) *DescribeDomainRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetRequestId(v string) *DescribeDomainRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetDomainRecords(v *DescribeDomainRecordsResponseBodyDomainRecords) *DescribeDomainRecordsResponseBody {
	s.DomainRecords = v
	return s
}

func (s *DescribeDomainRecordsResponseBody) SetPageNumber(v int64) *DescribeDomainRecordsResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeDomainRecordsResponseBodyDomainRecords struct {
	Record []*DescribeDomainRecordsResponseBodyDomainRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s DescribeDomainRecordsResponseBodyDomainRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRecordsResponseBodyDomainRecords) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsResponseBodyDomainRecords) SetRecord(v []*DescribeDomainRecordsResponseBodyDomainRecordsRecord) *DescribeDomainRecordsResponseBodyDomainRecords {
	s.Record = v
	return s
}

type DescribeDomainRecordsResponseBodyDomainRecordsRecord struct {
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

func (s DescribeDomainRecordsResponseBodyDomainRecordsRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRecordsResponseBodyDomainRecordsRecord) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetStatus(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Status = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetType(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Type = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetRemark(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Remark = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetTTL(v int64) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.TTL = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetRecordId(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetPriority(v int64) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Priority = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetRR(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.RR = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetDomainName(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetWeight(v int32) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Weight = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetValue(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Value = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetLine(v string) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Line = &v
	return s
}

func (s *DescribeDomainRecordsResponseBodyDomainRecordsRecord) SetLocked(v bool) *DescribeDomainRecordsResponseBodyDomainRecordsRecord {
	s.Locked = &v
	return s
}
