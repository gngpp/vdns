package auth

type BasisCredential struct {
	secretId  string
	secretKey string
}

func (b *BasisCredential) GetSecretId() string {
	return b.secretId
}

func (b *BasisCredential) GetToken() string {
	panic("unrealized")
}

func (b *BasisCredential) GetSecretKey() string {
	return b.secretKey
}
