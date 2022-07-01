package vnet

import (
	"io/ioutil"
	"regexp"
	"vdns/config"
	"vdns/lib/util/iotool"
	"vdns/lib/util/vhttp"
	"vdns/lib/vlog"
)

const Ipv4Reg = `((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])`

const Ipv6Reg = `((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))`

func GetIpv4AddrForName(card string) []string {
	v4List, _, err := GetCardInterface()
	if err != nil {
		return []string{}
	}
	if v4List != nil {
		ipAddr := make([]string, 0)
		for _, i := range v4List {
			if i.Name == card && i.Ipv4() {
				for _, address := range i.Address {
					if IsPrivateAddr(address) {
						vlog.Debugf("ip %v is private address", address)
					} else {
						ipAddr = append(ipAddr, address)
					}
				}
			}
		}
		return ipAddr
	}
	return []string{}
}

func GetIpv6AddrForName(card string) []string {
	_, v6List, err := GetCardInterface()
	if err != nil {
		return []string{}
	}
	if v6List != nil {
		ipAddr := make([]string, 0)
		for _, i := range v6List {
			if i.Name == card && i.Ipv6() {
				for _, address := range i.Address {
					if IsPrivateAddr(address) {
						vlog.Debugf("ip %v is private address", address)
					} else {
						ipAddr = append(ipAddr, address)
					}
				}
			}
		}
		return ipAddr
	}
	return []string{}
}

func GetIpv4AddrForUrl(url string) string {
	return getIpAddr(Ipv4Reg, url)
}

func GetIpv6AddrForUrl(url string) string {
	return getIpAddr(Ipv6Reg, url)
}

func GetIpv4Addr() string {
	return getIpAddr(Ipv4Reg, config.V4Api1)
}

func GetIpv6Addr() string {
	return getIpAddr(Ipv6Reg, config.V6Api1)
}

func getIpAddr(reg string, url string) (result string) {
	resp, err := vhttp.Get(url, "")
	if err != nil {
		vlog.Error(err)
		return
	}

	body := resp.Body
	defer iotool.ReadCloser(body)
	bytes, err := ioutil.ReadAll(body)
	vlog.Debugf("request: %s body: %v", url, string(bytes))
	if err != nil {
		vlog.Error("Failed to read ip result! Query URL: ", url)
		return
	}
	comp := regexp.MustCompile(reg)
	return comp.FindString(string(bytes))
}
