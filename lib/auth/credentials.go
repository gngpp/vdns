package auth

type Credential interface {
	GetSecretId() string
	GetToken() string
	GetSecretKey() string
}
