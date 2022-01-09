package alidns_model

import "vdns/util/vjson"

type UpdateDomainRecordResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s *UpdateDomainRecordResponse) String() string {
	return vjson.PrettifyString(s)
}
