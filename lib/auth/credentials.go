package auth

import "vdns/lib/util/strs"

type Credential interface {
	GetSecretId() string
	GetToken() string
	GetSecretKey() string
}

type UnifyCredential struct {
	secretId  *string
	secretKey *string
	token     *string
}

func (_this UnifyCredential) GetSecretId() string {
	return strs.StringValue(_this.secretId)
}

func (_this UnifyCredential) GetToken() string {
	return strs.StringValue(_this.token)
}

func (_this UnifyCredential) GetSecretKey() string {
	return strs.StringValue(_this.secretKey)
}

func NewUnifyCredential(secretId, secretKey, token string) Credential {
	return &UnifyCredential{
		secretId:  &secretId,
		secretKey: &secretKey,
		token:     &token,
	}
}
