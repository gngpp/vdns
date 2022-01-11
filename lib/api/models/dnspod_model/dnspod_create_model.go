package dnspod_model

import (
	"vdns/lib/util/vjson"
)

type CreateRecordResponse struct {
	Response *struct {

		// 记录ID
		RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

		// 返回错误信息
		Error *Error `json:"Error" name:"Error"`
	} `json:"Response"`
}

func (s *CreateRecordResponse) String() string {
	return vjson.PrettifyString(s)
}
