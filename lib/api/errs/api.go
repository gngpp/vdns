package errs

import "vdns/vutil/str"

type ApiError struct {
	message *string
}

func NewApiError(message string) *ApiError {
	return &ApiError{message: &message}
}

func NewApiErrorFromError(e error) *ApiError {
	err := e.Error()
	return &ApiError{message: &err}
}

func (_this *ApiError) Error() string {
	return str.StringValue(_this.message)
}
