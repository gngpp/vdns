package vhttp

import (
	"errors"
	"strings"
	"vdns/util/strs"
)

var rareDomainSuffix = []string{"com.cn", "org.cn", "net.cn", "ac.cn", "eu.org"}

type Domain struct {
	Domain    string
	SubDomain string
}

// CheckExtractDomain
// 提取顶级主域名跟域名记录
// example：
// www.baidu.com -> 顶级域名：baidu.com  记录：www
// a.b.baidu.com -> 顶级域名：baidu.com   记录：a.b
func CheckExtractDomain(domain string) (*Domain, error) {
	split := strings.Split(domain, ".")
	length := len(split)
	if strs.IsEmpty(domain) || length <= 1 {
		return nil, errors.New("the domain name does not meet the specification")
	}

	targetDomain := new(Domain)
	targetDomain.Domain = strs.Concat(split[length-2], ".", split[length-1])
	// add special suffix domain name processing
	for _, staticMainDomain := range rareDomainSuffix {
		if staticMainDomain == targetDomain.Domain {
			targetDomain.Domain = strs.Concat(split[length-3], ".", targetDomain.Domain)
			break
		}
	}

	index := strings.LastIndex(domain, targetDomain.Domain)
	if index > 0 {
		targetDomain.SubDomain = domain[:index-1]
	}

	return targetDomain, nil
}
