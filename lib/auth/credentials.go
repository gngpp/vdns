package auth

import (
	"vdns/lib/api/errs"
	"vdns/vutil/str"
)

type Credential interface {
	GetSecretId() string
	GetToken() string
	GetSecretKey() string
}

//goland:noinspection ALL
func NewBasicCredential(secretId, secretKey string) (Credential, error) {
	if str.IsEmpty(secretId) {
		return nil, errs.NewCredentialsError("Access key ID cannot be empty")
	}
	if str.IsEmpty(secretKey) {
		return nil, errs.NewCredentialsError("Access key secret cannot be empty")
	}
	return &BasisCredential{
		secretId:  secretId,
		secretKey: secretKey,
	}, nil
}

//goland:noinspection ALL
func NewTokenCredential(token string) (Credential, error) {
	if str.IsEmpty(token) {
		return nil, errs.NewCredentialsError("Token secret cannot be empty")
	}
	tokenCredntial := &TokenCredential{
		token: token,
	}
	return tokenCredntial, nil
}
