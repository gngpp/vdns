package vhttp

import (
	"net"
	"net/http"
	"time"
)

func NewClient() *http.Client {
	dialer := &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	return &http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			DisableKeepAlives:   false,
			DialContext:         dialer.DialContext,
			IdleConnTimeout:     30 * time.Second,
			TLSHandshakeTimeout: 30 * time.Second,
		},
	}
}
