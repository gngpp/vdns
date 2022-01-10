package alidns_model

import (
	"github.com/zf1976/vdns/lib/util/vjson"
)

type CreateDomainRecordResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s *CreateDomainRecordResponse) String() string {
	return vjson.PrettifyString(s)
}
