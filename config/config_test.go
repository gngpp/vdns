package config

import (
	"fmt"
	"testing"
	"vdns/lib/util/file"
	"vdns/lib/util/vjson"
)

func TestCreate(t *testing.T) {
	config := NewVdnsConfig()
	println(vjson.PrettifyString(config))
}

func TestReadConfig(t *testing.T) {
	config, err := LoadVdnsConfig()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vjson.PrettifyString(config))
}

func TestCreateLogDir(t *testing.T) {
	err := file.MakeDir("/Users/gngpp/.vdns/logs")
	if err != nil {
		t.Error(err)
	}
}
