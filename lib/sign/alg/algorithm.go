package alg

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
)

var SignMethodMap = map[string]func() hash.Hash{
	"HMAC-SHA1":   sha1.New,
	"HMAC-SHA256": sha256.New,
	"HMAC-MD5":    md5.New,
}

var HmacSha1 = sha1.New()
var HmacSha256 = sha256.New()
var HmacMD5 = md5.New()
