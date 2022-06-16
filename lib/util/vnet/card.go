package vnet

import (
	"fmt"
	"net"
	"strings"
)

// Interface 本机网络
type Interface struct {
	Name    string
	Address []string
	ipv4    bool
	ipv6    bool
}

func (_this *Interface) Ipv4() bool {
	return _this.ipv4 && !_this.ipv6
}

func (_this *Interface) Ipv6() bool {
	return _this.ipv6 && !_this.ipv4
}

// GetInterface 获得网卡地址
func GetInterface() (ipv4NetInterfaces []Interface, ipv6NetInterfaces []Interface, err error) {
	allNetInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, iotool:", err.Error())
		return ipv4NetInterfaces, ipv6NetInterfaces, err
	}

	for i := 0; i < len(allNetInterfaces); i++ {
		if (allNetInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := allNetInterfaces[i].Addrs()
			var ipv4 []string
			var ipv6 []string

			for _, address := range addrs {
				if ipNet, ok := address.(*net.IPNet); ok && ipNet.IP.IsGlobalUnicast() {
					_, bits := ipNet.Mask.Size()
					if bits == 128 && !ipNet.IP.IsLoopback() {
						to16 := ipNet.IP.To16()
						if to16 != nil {
							ipv6 = append(ipv6, to16.String())
						}
					}
					if bits == 32 && !ipNet.IP.IsLoopback() {
						to4 := ipNet.IP.To4()
						if to4 != nil {
							ipv4 = append(ipv4, to4.String())
						}
					}
				}
			}

			if len(ipv4) > 0 {
				ipv4NetInterfaces = append(
					ipv4NetInterfaces,
					Interface{
						Name:    allNetInterfaces[i].Name,
						Address: ipv4,
						ipv4:    true,
					},
				)
			}

			if len(ipv6) > 0 {
				ipv6NetInterfaces = append(
					ipv6NetInterfaces,
					Interface{
						Name:    allNetInterfaces[i].Name,
						Address: ipv6,
						ipv6:    true,
					},
				)
			}

		}
	}

	return ipv4NetInterfaces, ipv6NetInterfaces, nil
}

func IsPrivateAddr(remoteAddr string) bool {
	lastIndex := strings.LastIndex(remoteAddr, ":")
	if lastIndex < 1 {
		return false
	}

	remoteAddr = remoteAddr[:lastIndex]

	// ipv6
	if strings.HasPrefix(remoteAddr, "[") && strings.HasSuffix(remoteAddr, "]") {
		remoteAddr = remoteAddr[1 : len(remoteAddr)-1]
	}

	if ip := net.ParseIP(remoteAddr); ip != nil {
		if ip.IsLoopback() {
			return true
		}

		_, ipNet192, _ := net.ParseCIDR("192.168.0.0/16")
		if ipNet192.Contains(ip) {
			return true
		}

		_, ipNet172, _ := net.ParseCIDR("172.16.0.0/12")
		if ipNet172.Contains(ip) {
			return true
		}

		_, ipNet10, _ := net.ParseCIDR("10.0.0.0/8")
		if ipNet10.Contains(ip) {
			return true
		}

		_, ipNetFE, _ := net.ParseCIDR("fe80::/10")
		if ipNetFE.Contains(ip) {
			return true
		}

		_, ipNetV6FD, _ := net.ParseCIDR("fd00::/8")
		if ipNetV6FD.Contains(ip) {
			return true
		}

	}

	// localhost
	if remoteAddr == "localhost" {
		return true
	}

	return false
}
