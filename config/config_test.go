package config

import (
	"testing"
	"vdns/lib/util/vjson"
)

func TestName(t *testing.T) {
	config := NewConfig()
	println(vjson.PrettifyString(config))
}
