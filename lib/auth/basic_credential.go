package auth

import "vdns/vutil/strs"

type BasisCredential struct {
	secretId  *string
	secretKey *string
}

func (_this BasisCredential) GetSecretId() string {
	return strs.StringValue(_this.secretId)
}

func (_this BasisCredential) GetToken() string {
	return ""
}

func (_this BasisCredential) GetSecretKey() string {
	return strs.StringValue(_this.secretKey)
}

func NewBasicCredential(secretId, secretKey string) Credential {
	return &BasisCredential{
		secretId:  &secretId,
		secretKey: &secretKey,
	}
}
