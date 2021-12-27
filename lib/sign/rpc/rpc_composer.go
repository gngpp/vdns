package rpc

import (
	"net/url"
	"vdns/vutil/vhttp"
)

const SEPARATOR = "&"

//goland:noinspection ALL
type RpcSignatureComposer interface {

	// ComposeStringToSign 组合签名必要参数
	ComposeStringToSign(method vhttp.HttpMethod, queries *url.Values) string

	// GeneratedSignature 生成签名
	GeneratedSignature(secret string, method vhttp.HttpMethod, queries *url.Values) string

	// SignatureMethod 签名方法
	SignatureMethod() string

	// SignerVersion 签名版本
	SignerVersion() string

	// CanonicalizeRequestUrl 生成规范请求URL
	CanonicalizeRequestUrl(urlPattern string, queries *url.Values) string
}
