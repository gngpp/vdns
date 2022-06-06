package errs

import (
	"fmt"
)

type CloudFlareSDKError struct {
	Message string
}

func (_this *CloudFlareSDKError) Error() string {
	return fmt.Sprintf("[CloudFlareSDKError]:\n\tMessage: %s", _this.Message)
}

func NewCloudFlareSDKError(msg string) error {
	return &CloudFlareSDKError{msg}
}
