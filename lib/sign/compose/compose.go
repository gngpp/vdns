package compose

import (
	"github.com/zf1976/vdns/lib/util/vhttp"
	log "github.com/zf1976/vdns/lib/vlog"
	"net/url"
)

var vlog = log.Default()

const SEPARATOR = "&"

type SignatureComposer interface {

	// ComposeStringToSign 组合签名必要参数
	ComposeStringToSign(method vhttp.HttpMethod, queries *url.Values) string

	// GeneratedSignature 生成签名
	GeneratedSignature(secret string, stringToSign string) string

	// CanonicalizeRequestUrl 生成规范请求URL
	CanonicalizeRequestUrl(urlPattern, signature string, queries *url.Values) string

	// SignatureMethod 签名方法
	SignatureMethod() string

	// SignerVersion 签名版本
	SignerVersion() string
}
