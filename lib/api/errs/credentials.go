package errs

import (
	"fmt"
	"vdns/vutil/strs"
)

type CredntialsError struct {
	Message *string
}

func (_this *CredntialsError) Error() string {
	return fmt.Sprintf("CredntialsError: %s", strs.StringValue(_this.Message))
}

func NewCredentialsError(message string) error {
	return &CredntialsError{
		Message: &message,
	}
}
