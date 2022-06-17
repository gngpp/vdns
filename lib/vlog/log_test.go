package vlog

import (
	"testing"
	"vdns/lib/vlog/timewriter"
)

func TestLog(t *testing.T) {
	timeWriter := &timewriter.TimeWriter{
		Dir:           "./logs",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vlog",
	}
	log := New(timeWriter)
	log.Info("hello vlog")
}
