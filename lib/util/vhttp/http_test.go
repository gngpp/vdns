package vhttp

import (
	"log"
	"testing"
)

func TestURL(t *testing.T) {
	a := IsURL("hi/there?")
	log.Print(a)

	b := IsURL("https://golang.cafe/")
	log.Print(b)

	c := IsURL("https://golang.org/index.html?#page1")
	log.Print(c)

	d := IsURL("golang.org")
	log.Print(d)
}
