package dnspod_model

import (
	"vdns/lib/util/vjson"
)

type Error struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

func (s *Error) String() string {
	return vjson.PrettifyString(s)
}
