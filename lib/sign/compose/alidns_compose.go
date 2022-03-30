package compose

import (
	"crypto/hmac"
	"encoding/base64"
	"net/url"
	"strings"
	"vdns/lib/sign/alg"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
	"vdns/lib/vlog"
)

func NewAliDNSSignatureCompose() SignatureComposer {
	return &AliDNSSignatureCompose{
		Separator:        strs.String(SEPARATOR),
		signatureMethod:  strs.String(alg.HMAC_SHA1),
		signatureVersion: strs.String("1.0"),
	}
}

type AliDNSSignatureCompose struct {
	Separator        *string
	signatureMethod  *string
	signatureVersion *string
}

func (_this *AliDNSSignatureCompose) ComposeStringToSign(method vhttp.HttpMethod, queries *url.Values) string {
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
	vlog.Debugf("[AliDNSSignatureCompose] stringToSign value: %s", stringToSign)
	return stringToSign
}

func (_this *AliDNSSignatureCompose) GeneratedSignature(secret string, stringToSign string) string {
	secret = strs.Concat(secret, strs.StringValue(_this.Separator))
	// compose sign string
	hash := hmac.New(alg.SignMethodMap[alg.HMAC_SHA1], strs.ToBytes(secret))
	_, err := hash.Write(strs.ToBytes(stringToSign))
	if err != nil {
		vlog.Debugf("[AliDNSSignatureCompose] hash encrypt error: %s", err)
		return ""
	}
	// encode
	encodeBytes := hash.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(encodeBytes)
	vlog.Debugf("[AliDNSSignatureCompose] signature: %s", signature)
	return signature
}

func (*AliDNSSignatureCompose) CanonicalizeRequestUrl(urlPattern, signature string, queries *url.Values) string {
	queries.Set("Signature", signature)
	requestUrl := strs.Concat(urlPattern, "?", queries.Encode())
	vlog.Debugf("[AliDNSSignatureCompose] request url: %s", requestUrl)
	return requestUrl
}

func (_this *AliDNSSignatureCompose) SignatureMethod() string {
	return strs.StringValue(_this.signatureMethod)
}

func (_this *AliDNSSignatureCompose) SignerVersion() string {
	return strs.StringValue(_this.signatureVersion)
}
