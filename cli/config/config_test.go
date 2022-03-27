package config

import (
	"fmt"
	"testing"
	"vdns/lib/util/file"
	"vdns/lib/util/vjson"
	"vdns/lib/vlog"
)

func TestConfig(t *testing.T) {
	read, err := file.Read(configPath)
	if err != nil {
		panic(err)
	}
	vlog.Infof("read:\n%v", read)
	var entity Config
	err = vjson.ByteArrayConvert([]byte(read), &entity)
	if err != nil {
		fmt.Println(err)
		return
	}
	vlog.Infof("entity:\n%v", entity)

}
