package config

import (
	"fmt"
	"testing"
	"vdns/vutil/file"
)

func TestName(t *testing.T) {
	println(GetConfigPath())
	println(getConfigPathDefault())
	println(GetLogPath())
	println(GetWorkingPath())
}

func Test(t *testing.T) {
	pathDefault, err := GetLogPath()
	if err != nil {
		fmt.Println(err)
	}
	println(file.IsFile(pathDefault))
}
