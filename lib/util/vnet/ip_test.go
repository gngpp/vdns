package vnet

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	v4Addr := GetPubIpv4Addr()
	fmt.Println(v4Addr)

	v6Addr := GetPubIpv6Addr()
	fmt.Println(v6Addr)
}
