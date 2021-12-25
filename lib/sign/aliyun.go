package sign

import (
	"net/url"
	"vdns/vutil/vhttp"
)

type AliyunRpcSignatureCompose struct {
}

func (a *AliyunRpcSignatureCompose) composeStringToSign(method vhttp.HttpMethod, query url.Values) string {
	//TODO implement me
	panic("implement me")
}

func (a *AliyunRpcSignatureCompose) toSignatureUrl() string {
	//TODO implement me
	panic("implement me")
}

func (a *AliyunRpcSignatureCompose) signatureMethod() string {
	//TODO implement me
	panic("implement me")
}

func (a *AliyunRpcSignatureCompose) signerVersion() string {
	//TODO implement me
	panic("implement me")
}

func (a *AliyunRpcSignatureCompose) canonicalizeRequestUrl() string {
	//TODO implement me
	panic("implement me")
}
