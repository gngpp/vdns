package vhttp

import (
	"errors"
	"regexp"
	"strings"
	"vdns/vutil/str"
)

var DomainRegexp = regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)

// ExtractDomain
// 提取顶级主域名跟域名记录
// example：
// www.baidu.com -> 顶级域名：baidu.com  记录：www
// a.b.baidu.com -> 顶级域名：baidu.com   记录：a.b
func ExtractDomain(domain string) ([]string, error) {
	if str.IsEmpty(domain) || !IsDomain(domain) {
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

// IsDomain Golang does not support Perl syntax ((?
func IsDomain(domain string) bool {
	return DomainRegexp.MatchString(domain)
}
