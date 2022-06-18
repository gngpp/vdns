package vhttp

import (
	"net/url"
)

func IsURL(rawUrl string) error {
	_, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return err
	}
	return nil
}
