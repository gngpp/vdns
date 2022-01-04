package alg

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
)

var SignMethodMap = map[string]func() hash.Hash{
	HMAC_SHA1:   sha1.New,
	HMAC_SHA256: sha256.New,
	HMAC_MD5:    md5.New,
}

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	HMAC_SHA1   = "HMAC-SHA1"
	HMAC_SHA256 = "HmacSHA256"
	HMAC_MD5    = "HMAC-MD5"
)
