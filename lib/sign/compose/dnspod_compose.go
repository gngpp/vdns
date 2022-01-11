package compose

import (
	"bytes"
	"crypto/hmac"
	"encoding/base64"
	"net/url"
	"sort"
	"vdns/lib/sign/alg"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
	"vdns/lib/vlog"
)

func NewDnspodSignatureCompose() SignatureComposer {
	return &DnspodSignatureCompose{
		Separator:       strs.String(SEPARATOR),
		signatureMethod: strs.String(alg.HMAC_SHA256),
	}
}

type DnspodSignatureCompose struct {
	Separator       *string
	signatureMethod *string
}

func (_this *DnspodSignatureCompose) ComposeStringToSign(method vhttp.HttpMethod, queries *url.Values) string {
	// sort encode
	buf := new(bytes.Buffer)
	buf.WriteString(strs.Concat(method.String(), "dnspod.tencentcloudapi.com/?"))
	// sort keys by ascii asc order
	keys := make([]string, 0, len(*queries))
	for k, _ := range *queries {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		k := keys[i]
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(queries.Get(k))
		buf.WriteString(strs.StringValue(_this.Separator))
	}
	buf.Truncate(buf.Len() - 1)
	stringToSign := buf.String()
	vlog.Debugf("[DnspodSignatureCompose] stringToSign value: %s", stringToSign)
	return stringToSign
}

func (_this *DnspodSignatureCompose) GeneratedSignature(secret string, stringToSign string) string {
	hash := hmac.New(alg.SignMethodMap[alg.HMAC_SHA256], strs.ToBytes(secret))
	_, err := hash.Write(strs.ToBytes(stringToSign))
	if err != nil {
		vlog.Debugf("[DnspodSignatureCompose] hash encrypt error: %s", err)
		return ""
	}
	encodeBytes := hash.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(encodeBytes)
	vlog.Debugf("[DnspodSignatureCompose] signature: %s", signature)
	return signature
}

func (_this *DnspodSignatureCompose) CanonicalizeRequestUrl(urlPattern, signature string, queries *url.Values) string {
	queries.Set("Signature", signature)
	requestUrl := strs.Concat(urlPattern, "?", queries.Encode())
	vlog.Debugf("[DnspodSignatureCompose] request url: %s", requestUrl)
	return requestUrl
}

func (_this *DnspodSignatureCompose) SignatureMethod() string {
	return strs.StringValue(_this.signatureMethod)
}

func (_this *DnspodSignatureCompose) SignerVersion() string {
	return ""
}
