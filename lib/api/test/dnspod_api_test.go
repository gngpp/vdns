package test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"
	"vdns/lib/sign/alg"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
	"vdns/vlog"
	"vdns/vutil/vhttp"
	"vdns/vutil/vjson"
)

func Test(t *testing.T) {
	vlog.SetLevel(vlog.Level.DEBUG)
	paramater := make(url.Values)
	paramater.Set("Nonce", strconv.FormatInt(rand.Int63()+time.Now().UnixMilli(), 10))
	paramater.Set("Timestamp", strconv.FormatInt(time.Now().UnixMilli()/1000, 10))
	paramater.Set("SecretId", "AKIDDG4bqoXzB6l9G6k8F5exyiBNIwJQGo55")
	paramater.Set("Domain", "innas.work")
	paramater.Set("Action", "DescribeRecordList")
	paramater.Set("Limit", "1000")
	paramater.Set("Offset", "0")
	paramater.Set("Keyword", "222.217.152.145")
	paramater.Set("Version", "2021-03-23")
	paramater.Set("SignatureMethod", alg.HMAC_SHA256)
	composer := compose.NewDnspodSignatureCompose()
	stringToSign := composer.ComposeStringToSign(vhttp.HttpMethodGet, &paramater)
	signature := composer.GeneratedSignature("lGTP5UbS4PN4UdV6PRCgtqWv3tSlnY3G", stringToSign)
	requestUrl := composer.CanonicalizeRequestUrl(standard.DNSPOD_DNS_API.StringValue(), signature, &paramater)
	resp, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(vjson.PrettifyString(vjson.ToMap(string(bytes))))
}
