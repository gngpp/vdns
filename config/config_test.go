package config

import (
	"testing"
	"vdns/lib/util/vjson"
)

func TestCreate(t *testing.T) {
	config := NewVdnsConfig()
	println(vjson.PrettifyString(config))
}

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig()
	if err != nil {
		t.Error(err)
	}
	println(vjson.PrettifyString(config))
}
