package response

type QueryDomainListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDomainListResponseBody `json:"response,omitempty" xml:"response,omitempty" require:"true"`
}

type QueryDomainListResponseBody struct {
	PrePage        *bool                            `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	CurrentPageNum *int32                           `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalPageNum   *int32                           `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	Data           *QueryDomainListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	TotalItemNum   *int32                           `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	NextPage       *bool                            `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
}

type QueryDomainListResponseBodyData struct {
	Domain []*QueryDomainListResponseBodyDataDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

type QueryDomainListResponseBodyDataDomain struct {
	DomainAuditStatus      *string `json:"DomainAuditStatus,omitempty" xml:"DomainAuditStatus,omitempty"`
	DomainGroupId          *string `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	Remark                 *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	DomainGroupName        *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	RegistrationDate       *string `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DomainName             *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ExpirationDateStatus   *string `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	ExpirationDate         *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	RegistrantType         *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	ExpirationDateLong     *int64  `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	ExpirationCurrDateDiff *int32  `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	Premium                *bool   `json:"Premium,omitempty" xml:"Premium,omitempty"`
	RegistrationDateLong   *int64  `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	ProductId              *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	DomainStatus           *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainType             *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
}
