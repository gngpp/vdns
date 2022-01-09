package errs

import (
	"fmt"
	"vdns/util/strs"
)

// AlidnsSDKError struct is used save error code and msg
type AlidnsSDKError struct {
	RequestId *string `json:"request_id,omitempty"`
	Recommend *string `json:"recommend,omitempty"`
	Code      *string `json:"code,omitempty"`
	Message   *string `json:"message,omitempty"`
	Stack     *string `json:"stack,omitempty"`
	errMsg    *string `json:"err_msg,omitempty"`
}

func (_this *AlidnsSDKError) Error() string {
	if _this.errMsg == nil {
		str := fmt.Sprintf("[AlidnsSDKError]:\n\tCode: %s\n\tMessage: %s\n",
			strs.StringValue(_this.Code),
			strs.StringValue(_this.Message),
		)
		_this.SetErrMsg(str)
	}
	return strs.StringValue(_this.errMsg)
}

// SetErrMsg Set ErrMsg by msg
func (_this *AlidnsSDKError) SetErrMsg(msg string) {
	_this.errMsg = strs.String(msg)
}
