package vhttp

import (
	"fmt"
	"log"
	"testing"
)

func TestURL(t *testing.T) {
	a := IsURL("hi/there?")
	log.Print(a)

	b := IsURL("http://golang.cafe/")
	log.Print(b)

	c := IsURL("http://golang.org/index.html?#page1")
	log.Print(c)

	d := IsURL("golang.org")
	log.Print(d)
}

func TestDomain(t *testing.T) {
	domName := "www.golang.org"

	if !IsDomain(domName) {
		fmt.Printf("Domain Name %s is invalid\n", domName)
	} else {
		fmt.Printf("Domain Name %s is VALID\n", domName)
	}

	domName = "www.socketloop,.com"

	if !IsDomain(domName) {
		fmt.Printf("Domain Name %s is invalid\n", domName)
	} else {
		fmt.Printf("Domain Name %s is VALID\n", domName)
	}

	domName = "subdomain-socketloop.com"

	if !IsDomain(domName) {
		fmt.Printf("Domain Name %s is invalid\n", domName)
	} else {
		fmt.Printf("Domain Name %s is VALID\n", domName)
	}

	domName = "-socketloop.com" // invalid starts with hyphen

	if !IsDomain(domName) {
		fmt.Printf("Domain Name %s is invalid\n", domName)
	} else {
		fmt.Printf("Domain Name %s is VALID\n", domName)
	}

	domName = "socketloop.co_" // invalid ends with underscore

	if !IsDomain(domName) {
		fmt.Printf("Domain Name %s is invalid\n", domName)
	} else {
		fmt.Printf("Domain Name %s is VALID\n", domName)
	}

	domName = "subdomain.socketloop.com"

	if !IsDomain(domName) {
		fmt.Printf("Domain Name %s is invalid\n", domName)
	} else {
		fmt.Printf("Domain Name %s is VALID\n", domName)
	}

}
