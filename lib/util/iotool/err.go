package iotool

import (
	"io"
	"vdns/lib/vlog"
)

func ReadCloser(io io.ReadCloser) {
	err := io.Close()
	if err != nil {
		vlog.Error(err)
	}
}
