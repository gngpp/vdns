package vhttp

import (
	"net/url"
)

func IsURL(rawUrl string) bool {
	_, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return false
	}
	return true
}
