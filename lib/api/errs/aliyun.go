package errs

import (
	"encoding/json"
	"fmt"
	"strconv"
	"vdns/vutil/str"
)

// AliyunSDKError struct is used save error code and msg
type AliyunSDKError struct {
	RequestId *string `json:"request_id,omitempty"`
	Recommend *string `json:"recommend,omitempty"`
	Code      *string `json:"code,omitempty"`
	Message   *string `json:"message,omitempty"`
	Data      *string `json:"data,omitempty"`
	Stack     *string `json:"stack,omitempty"`
	errMsg    *string `json:"err_msg,omitempty"`
}

func (_this *AliyunSDKError) Error() string {
	if _this.errMsg == nil {
		str := fmt.Sprintf("AliyunSDKError:\n   Code: %s\n   Message: %s\n   Data: %s\n",
			str.StringValue(_this.Code),
			str.StringValue(_this.Message),
			str.StringValue(_this.Data),
		)
		_this.SetErrMsg(str)
	}
	return str.StringValue(_this.errMsg)
}

// SetErrMsg Set ErrMsg by msg
func (_this *AliyunSDKError) SetErrMsg(msg string) {
	_this.errMsg = str.String(msg)
}

// NewAliyunSDKError is used for shortly create AliyunSDKError object
func NewAliyunSDKError(obj map[string]interface{}) *AliyunSDKError {
	err := &AliyunSDKError{}
	if val, ok := obj["code"].(int); ok {
		err.Code = str.String(strconv.Itoa(val))
	} else if val, ok := obj["code"].(string); ok {
		err.Code = str.String(val)
	}

	if obj["msg"] != nil {
		err.Message = str.String(obj["msg"].(string))
	}
	if data := obj["data"]; data != nil {
		byt, _ := json.Marshal(data)
		err.Data = str.String(string(byt))
	}
	return err
}
