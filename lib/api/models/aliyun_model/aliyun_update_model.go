package aliyun_model

import "vdns/vutil/vjson"

type UpdateDomainRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s *UpdateDomainRecordResponseBody) String() string {
	return vjson.PrettifyString(s)
}
