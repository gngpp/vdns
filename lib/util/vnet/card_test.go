package vnet

import (
	"fmt"
	"net"
	"testing"
)

func Test(t *testing.T) {
	data := map[string]bool{
		"127.0.0.1:9876":            true,
		"[::1]:9876":                true,
		"192.168.1.18:9876":         true,
		"172.16.1.18:9876":          true,
		"10.1.1.18:9876":            true,
		"100.0.0.1:9876":            false,
		"fe80::ba80:35ff:fe6b:4f7b": true,
		"[2409::1]:9876":            false,
		"223.5.5.5:9876":            false,
		"114.114.114.114":           false,
	}

	for key, value := range data {
		if IsPrivateAddr(key) != value {
			t.Errorf("%s verification failed\n", key)
		}
	}
}

func TestGetNetInterface(t *testing.T) {
	ipv4NetInterfaces, ipv6NetInterfaces, err := GetCardInterface()
	if err != nil {
		t.Error(err)
	}
	t.Log(ipv4NetInterfaces, ipv6NetInterfaces)

	interfaces, iotool := net.Interfaces()
	if iotool != nil {
		return
	}
	for i := range interfaces {
		inter := interfaces[i]
		addrs, iotool := inter.Addrs()
		if iotool != nil {
			return
		}
		fmt.Print(inter.Name + ":")
		fmt.Println(addrs)
	}
}
