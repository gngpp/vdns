package auth

import (
	"github.com/zf1976/vdns/lib/util/strs"
)

type TokenCredential struct {
	token *string
}

func (_this *TokenCredential) GetSecretId() string {
	return ""
}

func (_this *TokenCredential) GetToken() string {
	return strs.StringValue(_this.token)
}

func (_this *TokenCredential) GetSecretKey() string {
	return ""
}

func NewTokenCredential(token string) Credential {
	return &TokenCredential{
		token: &token,
	}
}
