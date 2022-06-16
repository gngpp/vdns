package vnet

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	v4Addr := GetIpv4Addr()
	fmt.Println(v4Addr)

	v6Addr := GetIpv6Addr()
	fmt.Println(v6Addr)
}
