package models

import "vdns/lib/standard/record"

// on aliyun document https://help.aliyun.com/document_detail/29776.html

// DescribeDomainRecordsRequest
// RRKeyWord 记录关键字
// ValueKeyWord 记录值关键字
// PageSize 分页大小
// PageNumber 当前页数
type DescribeDomainRecordsRequest struct {
	Domain       *string     `json:"domain,omitempty"`
	RecordType   record.Type `json:"record_type,omitempty"`
	RRKeyWord    *string     `json:"rr_key_word,omitempty"`
	ValueKeyWord *string     `json:"value_key_word,omitempty"`
	PageNumber   *int64      `json:"page_number,omitempty"`
	PageSize     *int64      `json:"page_size,omitempty"`
}

func (_this *DescribeDomainRecordsRequest) SetDomain(domain string) *DescribeDomainRecordsRequest {
	_this.Domain = &domain
	return _this
}

func (_this *DescribeDomainRecordsRequest) SetRecordType(recordType record.Type) *DescribeDomainRecordsRequest {
	_this.RecordType = recordType
	return _this
}

func (_this *DescribeDomainRecordsRequest) SetRRKeyWord(RRKeyWord string) *DescribeDomainRecordsRequest {
	_this.RRKeyWord = &RRKeyWord
	return _this
}

func (_this *DescribeDomainRecordsRequest) SetValueKeyWord(ValueKeyWord string) *DescribeDomainRecordsRequest {
	_this.ValueKeyWord = &ValueKeyWord
	return _this
}

func (_this *DescribeDomainRecordsRequest) SetPageNumber(PageNumber int64) *DescribeDomainRecordsRequest {
	_this.PageNumber = &PageNumber
	return _this
}

func (_this *DescribeDomainRecordsRequest) SetPageSize(PageSize int64) *DescribeDomainRecordsRequest {
	_this.PageSize = &PageSize
	return _this
}

func NewDescribeDomainRecordsRequest() *DescribeDomainRecordsRequest {
	defaultPageSize := int64(10)
	defaultPageNumber := int64(1)
	return &DescribeDomainRecordsRequest{
		PageSize:   &defaultPageSize,
		PageNumber: &defaultPageNumber,
	}
}

type CreateDomainRecordRequest struct {
	Domain     *string      `json:"domain,omitempty"`
	Value      *string      `json:"value,omitempty"`
	RecordType *record.Type `json:"record_type,omitempty"`
}

func (_this *CreateDomainRecordRequest) SetDomain(domain string) *CreateDomainRecordRequest {
	_this.Domain = &domain
	return _this
}

func (_this *CreateDomainRecordRequest) SetValue(value string) *CreateDomainRecordRequest {
	_this.Value = &value
	return _this
}

func (_this *CreateDomainRecordRequest) SetRecordType(recordType record.Type) *CreateDomainRecordRequest {
	_this.RecordType = &recordType
	return _this
}

func NewCreateDomainRecordRequest() *CreateDomainRecordRequest {
	return &CreateDomainRecordRequest{}
}

type UpdateDomainRecordRequest struct {
	ID         *string     `json:"id,omitempty"`
	Domain     *string     `json:"domain,omitempty"`
	Value      *string     `json:"value,omitempty"`
	RecordType record.Type `json:"record_type,omitempty"`
}

func (_this *UpdateDomainRecordRequest) SetID(id string) *UpdateDomainRecordRequest {
	_this.ID = &id
	return _this
}

func (_this *UpdateDomainRecordRequest) SetDomain(domain string) *UpdateDomainRecordRequest {
	_this.Domain = &domain
	return _this
}

func (_this *UpdateDomainRecordRequest) SetValue(value string) *UpdateDomainRecordRequest {
	_this.Value = &value
	return _this
}

func (_this *UpdateDomainRecordRequest) SetRecordType(recordType record.Type) *UpdateDomainRecordRequest {
	_this.RecordType = recordType
	return _this
}

func NewUpdateDomainRecordRequest() *UpdateDomainRecordRequest {
	return &UpdateDomainRecordRequest{}
}

type DeleteDomainRecordRequest struct {
	ID     *string `json:"id,omitempty"`
	Domain *string `json:"domain,omitempty"`
}

func (_this *DeleteDomainRecordRequest) SetID(id string) *DeleteDomainRecordRequest {
	_this.ID = &id
	return _this
}

func (_this *DeleteDomainRecordRequest) SetDomain(domain string) *DeleteDomainRecordRequest {
	_this.Domain = &domain
	return _this
}

func NewDeleteDomainRecordRequest() *DeleteDomainRecordRequest {
	return &DeleteDomainRecordRequest{}
}
