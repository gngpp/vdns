package aliyun_model

import "vdns/vutil/vjson"

type DeleteDomainRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s *DeleteDomainRecordResponseBody) String() string {
	return vjson.PrettifyString(s)
}
