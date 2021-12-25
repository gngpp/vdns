package auth

type Credential interface {
	GetSecretId() string
	GetToken() string
	GetSecretKey() string
}

//goland:noinspection ALL
func NewBasicCredential(secretId, secretKey string) Credential {
	return &BasisCredential{
		secretId:  secretId,
		secretKey: secretKey,
	}
}

//goland:noinspection ALL
func NewTokenCredential(token string) Credential {
	tokenCredntial := &TokenCredential{
		token: token,
	}
	return tokenCredntial
}
