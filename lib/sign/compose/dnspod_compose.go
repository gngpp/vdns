package compose

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"sort"
	"vdns/vlog"
	"vdns/vutil/strs"
	"vdns/vutil/vhttp"
)

func NewDnspodSignatureCompose(separator string) SignatureComposer {
	return &DnspodSignatureCompose{
		Separator:       strs.String(separator),
		signatureMethod: strs.String("HmacSHA256"),
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
	hash := hmac.New(sha256.New, strs.ToBytes(secret))
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

func (_this *DnspodSignatureCompose) CanonicalizeRequestUrl(urlPattern string, queries *url.Values) string {
	url := strs.Concat(urlPattern, "?", queries.Encode())
	vlog.Debugf("[DnspodSignatureCompose] request url: %s", url)
	return url
}

func (_this *DnspodSignatureCompose) SignatureMethod() string {
	return strs.StringValue(_this.signatureMethod)
}

func (_this *DnspodSignatureCompose) SignerVersion() string {
	return ""
}
