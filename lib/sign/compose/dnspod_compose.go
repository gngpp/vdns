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

func NewDNSPodSignatureCompose() SignatureComposer {
	return &DNSPodSignatureCompose{
		Separator:       strs.String(SEPARATOR),
		signatureMethod: strs.String(alg.HMAC_SHA256),
	}
}

type DNSPodSignatureCompose struct {
	Separator       *string
	signatureMethod *string
}

func (_this *DNSPodSignatureCompose) ComposeStringToSign(method vhttp.HttpMethod, queries *url.Values) string {
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
	vlog.Debugf("[DNSPodSignatureCompose] stringToSign value: %s", stringToSign)
	return stringToSign
}

func (_this *DNSPodSignatureCompose) GeneratedSignature(secret string, stringToSign string) string {
	hash := hmac.New(alg.SignMethodMap[alg.HMAC_SHA256], strs.ToBytes(secret))
	_, err := hash.Write(strs.ToBytes(stringToSign))
	if err != nil {
		vlog.Debugf("[DNSPodSignatureCompose] hash encrypt error: %s", err)
		return ""
	}
	encodeBytes := hash.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(encodeBytes)
	vlog.Debugf("[DNSPodSignatureCompose] signature: %s", signature)
	return signature
}

func (_this *DNSPodSignatureCompose) CanonicalizeRequestUrl(urlPattern, signature string, queries *url.Values) string {
	queries.Set("Signature", signature)
	requestUrl := strs.Concat(urlPattern, "?", queries.Encode())
	vlog.Debugf("[DNSPodSignatureCompose] request url: %s", requestUrl)
	return requestUrl
}

func (_this *DNSPodSignatureCompose) SignatureMethod() string {
	return strs.StringValue(_this.signatureMethod)
}

func (_this *DNSPodSignatureCompose) SignerVersion() string {
	return ""
}
