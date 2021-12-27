package errs

import (
	"encoding/json"
	"fmt"
	"strconv"
	"vdns/vutil/str"
)

// AliyunSDKError struct is used save error code and message
type AliyunSDKError struct {
	Code    *string
	Message *string
	Data    *string
	Stack   *string
	errMsg  *string
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

	if obj["message"] != nil {
		err.Message = str.String(obj["message"].(string))
	}
	if data := obj["data"]; data != nil {
		byt, _ := json.Marshal(data)
		err.Data = str.String(string(byt))
	}
	return err
}
