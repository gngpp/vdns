package config

import (
	"testing"
	"vdns/lib/util/vjson"
)

func TestName(t *testing.T) {
	config := NewVdnsConfig()
	println(vjson.PrettifyString(config))
}
