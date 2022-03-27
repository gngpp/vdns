package config

import (
	"os"
	"vdns/lib/homedir"
	"vdns/lib/util/file"
	"vdns/lib/util/strs"
	"vdns/lib/util/vjson"
)

var configPath string

type Config struct {
	Configs []*DNSConfig
}

type DNSConfig struct {
	Name *string `json:"name"`
	AK   *string `json:"ak,omitempty"`
	SK   *string `json:"sk,omitempty"`
}

func init() {

	dir, err := homedir.Dir()
	if err != nil {
		panic("system error")
		return
	}

	WorkingDir := strs.Concat(dir, "/.vdns")
	if !file.Exist(WorkingDir) {
		err = file.MakeDir(WorkingDir)
		if err != nil {
			panic("creating working " + WorkingDir + " directory error: " + err.Error())
			return
		}
	}

	configPath = strs.Concat(WorkingDir, "/config.json")
	if !file.Exist(configPath) {
		file, err := os.Create(configPath)
		defer file.Close()
		if err != nil {
			panic("creating " + configPath + " config file error: " + err.Error())
			return
		}
		config := &Config{
			Configs: []*DNSConfig{},
		}
		_, err = file.WriteString(vjson.PrettifyString(config))
		if err != nil {
			panic("initializing " + configPath + " config file error: " + err.Error())
			return
		}
	}
}
