package compose

import (
	"crypto/hmac"
	"crypto/sha1"
	"hash"
	"net/url"
	"strings"
	"vdns/vlog"
	"vdns/vutil/strs"
	"vdns/vutil/vhttp"
)

var log = vlog.Default()

func NewAlidnsSignatureCompose(separator string) SignatureComposer {
	return &AlidnsSignatureCompose{
		Separator: strs.String(separator),
	}
}

type AlidnsSignatureCompose struct {
	Separator *string
	sha1      hash.Hash
}

func (_this *AlidnsSignatureCompose) ComposeStringToSign(method vhttp.HttpMethod, queries *url.Values) string {
	// sort encode
	encode := queries.Encode()
	if log.IsDebugEnabled() {
		log.Debugf("aliyun api raw canonicalizedString value: %s", encode)
	}
	one := strings.ReplaceAll(encode, "+", "%20")
	two := strings.ReplaceAll(one, "*", "%2A")
	canonicalizedString := strings.ReplaceAll(two, "%7E", "~")
	return strs.Concat(string(method),
		strs.StringValue(_this.Separator),
		url.QueryEscape("/"),
		strs.StringValue(_this.Separator),
		url.QueryEscape(canonicalizedString))
}

func (_this *AlidnsSignatureCompose) GeneratedSignature(secret string, stringToSign string) string {
	secret = strs.Concat(secret, strs.StringValue(_this.Separator))
	// compose sign string
	hash := hmac.New(sha1.New, strs.ToBytes(secret))
	hash.Write(strs.ToBytes(stringToSign))
	// encode
	encodeBytes := hash.Sum(nil)
	return strs.ToString(encodeBytes)
}

func (*AlidnsSignatureCompose) SignatureMethod() string {
	return "HMAC-SHA1"
}

func (*AlidnsSignatureCompose) SignerVersion() string {
	return "1.0"
}

func (*AlidnsSignatureCompose) CanonicalizeRequestUrl(urlPattern string, queries *url.Values) string {
	return strs.Concat(urlPattern, "?", queries.Encode())
}
