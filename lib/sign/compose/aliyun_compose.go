package compose

import (
	"crypto/hmac"
	"encoding/base64"
	"net/url"
	"strings"
	"vdns/lib/sign/alg"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
)

func NewAlidnsSignatureCompose() SignatureComposer {
	return &AlidnsSignatureCompose{
		Separator:        strs.String(SEPARATOR),
		signatureMethod:  strs.String(alg.HMAC_SHA1),
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
	vlog.Debugf("[AlidnsSignatureCompose] stringToSign value: %s", stringToSign)
	return stringToSign
}

func (_this *AlidnsSignatureCompose) GeneratedSignature(secret string, stringToSign string) string {
	secret = strs.Concat(secret, strs.StringValue(_this.Separator))
	// compose sign string
	hash := hmac.New(alg.SignMethodMap[alg.HMAC_SHA1], strs.ToBytes(secret))
	_, err := hash.Write(strs.ToBytes(stringToSign))
	if err != nil {
		vlog.Debugf("[AlidnsSignatureCompose] hash encrypt error: %s", err)
		return ""
	}
	// encode
	encodeBytes := hash.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(encodeBytes)
	vlog.Debugf("[AlidnsSignatureCompose] signature: %s", signature)
	return signature
}

func (*AlidnsSignatureCompose) CanonicalizeRequestUrl(urlPattern, signature string, queries *url.Values) string {
	queries.Set("Signature", signature)
	requestUrl := strs.Concat(urlPattern, "?", queries.Encode())
	vlog.Debugf("[AlidnsSignatureCompose] request url: %s", requestUrl)
	return requestUrl
}

func (_this *AlidnsSignatureCompose) SignatureMethod() string {
	return strs.StringValue(_this.signatureMethod)
}

func (_this *AlidnsSignatureCompose) SignerVersion() string {
	return strs.StringValue(_this.signatureVersion)
}
