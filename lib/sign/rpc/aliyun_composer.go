package rpc

import (
	"crypto/hmac"
	"crypto/sha1"
	"hash"
	"net/url"
	"strings"
	"vdns/vlog"
	"vdns/vutil/str"
	"vdns/vutil/vhttp"
)

var log = vlog.Default()

func NewAliyunRpcSignatureCompose(separator string) RpcSignatureComposer {
	return &AliyunRpcSignatureCompose{
		Separator: str.String(separator),
	}
}

type AliyunRpcSignatureCompose struct {
	Separator *string
	sha1      hash.Hash
}

func (_this *AliyunRpcSignatureCompose) ComposeStringToSign(method vhttp.HttpMethod, queries *url.Values) string {
	// sort encode
	encode := queries.Encode()
	if log.IsDebugEnabled() {
		log.Debugf("aliyun api raw canonicalizedString value: %s", encode)
	}
	one := strings.ReplaceAll(encode, "+", "%20")
	two := strings.ReplaceAll(one, "*", "%2A")
	canonicalizedString := strings.ReplaceAll(two, "%7E", "~")
	return str.Concat(string(method),
		str.StringValue(_this.Separator),
		url.QueryEscape("/"),
		str.StringValue(_this.Separator),
		url.QueryEscape(canonicalizedString))
}

func (_this *AliyunRpcSignatureCompose) GeneratedSignature(secret string, stringToSign string) string {
	secret = str.Concat(secret, str.StringValue(_this.Separator))
	// compose sign string
	hash := hmac.New(sha1.New, str.ToBytes(secret))
	hash.Write(str.ToBytes(stringToSign))
	// encode
	encodeBytes := hash.Sum(nil)
	return str.ToString(encodeBytes)
}

func (*AliyunRpcSignatureCompose) SignatureMethod() string {
	return "HMAC-SHA1"
}

func (*AliyunRpcSignatureCompose) SignerVersion() string {
	return "1.0"
}

func (*AliyunRpcSignatureCompose) CanonicalizeRequestUrl(urlPattern string, queries *url.Values) string {
	return str.Concat(urlPattern, "?", queries.Encode())
}
