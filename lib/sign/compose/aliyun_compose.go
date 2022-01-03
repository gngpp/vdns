package compose

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"strings"
	"vdns/vutil/strs"
	"vdns/vutil/vhttp"
)

func NewAlidnsSignatureCompose(separator string) SignatureComposer {
	return &AlidnsSignatureCompose{
		Separator:        strs.String(separator),
		signatureMethod:  strs.String("HMAC-SHA1"),
		signatureVersion: strs.String("1.0"),
	}
}

type AlidnsSignatureCompose struct {
	Separator        *string
	signatureMethod  *string
	signatureVersion *string
}

func (_this *AlidnsSignatureCompose) ComposeStringToSign(method vhttp.HttpMethod, queries *url.Values) string {
	// sort encode
	encode := queries.Encode()
	one := strings.ReplaceAll(encode, "+", "%20")
	two := strings.ReplaceAll(one, "*", "%2A")
	canonicalizedString := strings.ReplaceAll(two, "%7E", "~")
	stringToSign := strs.Concat(method.String(),
		strs.StringValue(_this.Separator),
		url.QueryEscape("/"),
		strs.StringValue(_this.Separator),
		url.QueryEscape(canonicalizedString),
	)
	log.Debugf("[AlidnsSignatureCompose] stringToSign value: %s", stringToSign)
	return stringToSign
}

func (_this *AlidnsSignatureCompose) GeneratedSignature(secret string, stringToSign string) string {
	secret = strs.Concat(secret, strs.StringValue(_this.Separator))
	// compose sign string
	hash := hmac.New(sha1.New, strs.ToBytes(secret))
	_, err := hash.Write(strs.ToBytes(stringToSign))
	if err != nil {
		log.Debugf("[AlidnsSignatureCompose] hash encrypt error: %s", err)
		return ""
	}
	// encode
	encodeBytes := hash.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(encodeBytes)
	log.Debugf("[AlidnsSignatureCompose] signature: %s", signature)
	return signature
}

func (*AlidnsSignatureCompose) CanonicalizeRequestUrl(urlPattern string, queries *url.Values) string {
	url := strs.Concat(urlPattern, "?", queries.Encode())
	log.Debugf("[AlidnsSignatureCompose] request url: %s", url)
	return url
}

func (_this *AlidnsSignatureCompose) SignatureMethod() string {
	return strs.StringValue(_this.signatureMethod)
}

func (_this *AlidnsSignatureCompose) SignerVersion() string {
	return strs.StringValue(_this.signatureVersion)
}
