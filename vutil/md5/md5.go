package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Get16MD5Encode(data string) string {
	return GetMD5Encode(data)[8:24]
}

func SignWithPassword(name, password string) string {
	return Get16MD5Encode(fmt.Sprintf("vdns-%s-%s", password, name))
}

func SignWithToken(name, token string) string {
	return Get16MD5Encode(fmt.Sprintf("vdns-%s-%s", token, name))
}
