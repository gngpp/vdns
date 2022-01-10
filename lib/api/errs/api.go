package errs

import (
	"fmt"
	"vdns/lib/util/strs"
)

type VdnsError struct {
	message *string
}

func NewVdnsError(message string) *VdnsError {
	return &VdnsError{message: &message}
}

func NewApiErrorFromError(e error) *VdnsError {
	err := e.Error()
	return &VdnsError{message: &err}
}

func (_this *VdnsError) Error() string {
	return fmt.Sprintf("[VdnsError]:\n%s", strs.StringValue(_this.message))
}
