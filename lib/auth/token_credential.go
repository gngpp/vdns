package auth

type TokenCredential struct {
	token string
}

func (_this *TokenCredential) GetSecretId() string {
	panic("unrealized")
}

func (_this *TokenCredential) GetToken() string {
	return _this.token
}

func (_this *TokenCredential) GetSecretKey() string {
	panic("unrealized")
}

func NewTokenCredential(token string) Credential {
	return &TokenCredential{
		token: token,
	}
}
