package vhttp

import (
	"fmt"
	"testing"
)

func TestExtractDomain(t *testing.T) {
	domain, err := CheckExtractDomain("w.www.com.cn")
	if err != nil {
		fmt.Println(domain)
	}
	fmt.Println(*domain)

	domain, err = CheckExtractDomain("w.www.baidu.cn")
	if err != nil {
		fmt.Println(domain)
	}
	fmt.Println(*domain)

	domain, err = CheckExtractDomain("w.www.baidu.cn")
	if err != nil {
		fmt.Println(domain)
	}
	fmt.Println(*domain)

	domain, err = CheckExtractDomain("innas.work")
	if err != nil {
		fmt.Println(domain)
	}
	fmt.Println(*domain)

	domain, err = CheckExtractDomain("www.innas.work")
	if err != nil {
		fmt.Println(domain)
	}
	fmt.Println(*domain)
}
