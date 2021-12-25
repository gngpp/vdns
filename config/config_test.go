package config

import (
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
	println("")
	println(file.CurrentDir())
}
