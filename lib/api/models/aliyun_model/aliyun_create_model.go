package aliyun_model

import "vdns/vutil/vjson"

type CreateDomainRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s *CreateDomainRecordResponseBody) String() string {
	return vjson.PrettifyString(s)
}
