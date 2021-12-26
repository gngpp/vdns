package vhttp

import (
	"errors"
	"strings"
	"vdns/vutil/str"
)

// ExtractDomain
/**
 * 提取顶级主域名跟域名记录
 * 比如：www.baidu.com -> 顶级域名：baidu.com  记录：www 、   a.b.baidu.com -> 顶级域名：baidu.com   记录：a.b
 *
 * @param domain 域名
 * @return {@link String[]}
 */
func ExtractDomain(domain string) ([]string, error) {
	if str.IsEmpty(domain) {
		return nil, errors.New("the domain name does not meet the specification")
	}
	split := strings.Split(domain, ".")
	length := len(split)
	mainDomain := str.Concat(split[length-2], ".", split[length-1])
	rr := ""
	if length > 2 {
		index := strings.LastIndex(domain, mainDomain)
		rr = domain[:index-1]
	}
	return []string{mainDomain, rr}, nil
}
