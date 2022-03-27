package testing

import (
	"testing"
	"vdns/lib/util/md5"
)

func Test(t *testing.T) {
	println(md5.GetMD5Encode("lkjkljklkj"))
}
