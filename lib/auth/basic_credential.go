package auth

type BasisCredential struct {
	secretId  string
	secretKey string
}

func (_this BasisCredential) GetSecretId() string {
	return _this.secretId
}

func (_this BasisCredential) GetToken() string {
	panic("unrealized")
}

func (_this BasisCredential) GetSecretKey() string {
	return _this.secretKey
}

func NewBasicCredential(secretId, secretKey string) Credential {
	return &BasisCredential{
		secretId:  secretId,
		secretKey: secretKey,
	}
}
