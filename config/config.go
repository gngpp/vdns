package config

import (
	"os"
	"vdns/lib/homedir"
	"vdns/vutil/str"
)

//goland:noinspection ALL
const (
	VDNS_WORKING_NAME = ".vdns"
	VDNS_LOGS_NAME    = "logs"
	VDNS_CONFIG_NAME  = "config.json"
)

//goland:noinspection ALL
const (
	VDNS_WORKING_PATH_ENV = "vdns_working_path"
	VDNS_LOGS_PATH_ENV    = "vdns_logs_path"
	VDNS_CONFIG_PATH_ENV  = "vdns_config_path"
)

//goland:noinspection ALL
func GetWorkingPath() (string, error) {
	getenv := os.Getenv(VDNS_WORKING_PATH_ENV)
	if !str.IsEmpty(getenv) {
		return getenv, nil
	}
	dir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return str.Concat(dir, string(os.PathSeparator), VDNS_WORKING_NAME), nil
}

//goland:noinspection ALL
func GetLogPath() (string, error) {
	getenv := os.Getenv(VDNS_LOGS_PATH_ENV)
	if !str.IsEmpty(getenv) {
		return getenv, nil
	}
	dir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return str.Concat(dir, string(os.PathSeparator), VDNS_LOGS_NAME), nil
}

func GetConfigPath() (string, error) {
	configFilePath := os.Getenv(VDNS_CONFIG_PATH_ENV)
	if configFilePath != "" {
		return configFilePath, nil
	}
	return getConfigPathDefault()
}

func getConfigPathDefault() (string, error) {
	dir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return str.Concat(dir, string(os.PathSeparator), VDNS_CONFIG_NAME), nil
}
