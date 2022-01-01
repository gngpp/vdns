package dnspod_model

import "vdns/vutil/vjson"

type DeleteRecordResponse struct {
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (s *DeleteRecordResponse) String() string {
	return vjson.PrettifyString(s)
}
