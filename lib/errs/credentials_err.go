package errs

type CredntialsError struct {
	Message string
}

func (c *CredntialsError) Error() string {
	return c.Message
}

func NewCredentialsError(message string) error {
	return &CredntialsError{
		Message: message,
	}
}
