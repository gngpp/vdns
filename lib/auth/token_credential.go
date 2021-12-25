package auth

type TokenCredential struct {
	token string
}

func (t *TokenCredential) GetSecretId() string {
	panic("unrealized")
}

func (t *TokenCredential) GetToken() string {
	return t.token
}

func (t *TokenCredential) GetSecretKey() string {
	panic("unrealized")
}
