package config

import (
	"github.com/zf1976/vdns/lib/homedir"
	"github.com/zf1976/vdns/lib/util/file"
	"github.com/zf1976/vdns/lib/util/strs"
	"github.com/zf1976/vdns/lib/vlog"
	"github.com/zf1976/vdns/lib/vlog/timewriter"
	"io"
	"os"
)

var VConfig Config

type Config struct {
	username *string
	password *string
	port     *string
	log      *VLogConfig
	env      *Env
}

type VLogConfig struct {
	Level         string
	Dir           string
	Compress      bool
	ReserveDay    int
	LogFilePrefix string
}

type Env struct {
	workingPath *string
	logPath     *string
	config      *string
}

//goland:noinspection ALL
const (
	VDNS_WORKING_DIR_NAME = ".vdns"
	VDNS_LOGS_DIR_NAME    = "logs"
	VDNS_CONFIG_FILE_NAME = "config.json"
)

//goland:noinspection ALL
const (
	VDNS_WORKING_PATH_ENV = "vdns_working_path"
	VDNS_LOGS_PATH_ENV    = "vdns_logs_path"
	VDNS_CONFIG_PATH_ENV  = "vdns_config_path"
)

func GetWorkingPath() (string, error) {
	getenv := os.Getenv(VDNS_WORKING_PATH_ENV)
	if !strs.IsEmpty(getenv) {
		return getenv, nil
	}
	dir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return strs.Concat(dir, string(os.PathSeparator), VDNS_WORKING_DIR_NAME), nil
}

func GetLogPath() (string, error) {
	getenv := os.Getenv(VDNS_LOGS_PATH_ENV)
	if !strs.IsEmpty(getenv) {
		return getenv, nil
	}
	dir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return strs.Concat(dir, string(os.PathSeparator), VDNS_LOGS_DIR_NAME), nil
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
	return strs.Concat(dir, string(os.PathSeparator), VDNS_CONFIG_FILE_NAME), nil
}

func InitWorking() {

}

var log = vlog.Default()

func init() {
	InitWorking()
	timeWriter := &timewriter.TimeWriter{
		Dir:           "./logs",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vdns", // default is process name
	}
	writer := io.MultiWriter(os.Stdout, timeWriter)
	vlog.SetOutput(writer)
	workingPath, err := GetWorkingPath()
	if err != nil {
		panic(err)
	}
	if !file.IsDir(workingPath) {
		if err := file.MakeDir(workingPath); err != nil {
			panic(err)
		}
		log.Infof("[Init] working directory: %s\n", workingPath)
	} else {
		log.Infof("[Exist] working directory: %s exist\n", workingPath)
	}

	logPath, err := GetLogPath()
	if err != nil {
		panic(err)
	}
	if !file.IsDir(logPath) {
		if err := file.MakeDir(logPath); err != nil {
			panic(err)
		}
		log.Infof("[Init] logs directory: %s\n", logPath)
	} else {
		log.Infof("[Exist] logs directory: %s\n", logPath)
	}

	configPath, err := GetConfigPath()
	if err != nil {
		panic(err)
	}
	if !file.IsFile(configPath) {
		if err := file.Create(configPath); err != nil {
			panic(err)
		}
		log.Infof("[Init] config file: %s\n", configPath)
	} else {
		log.Infof("[Exist] config file: %s\n", configPath)
	}
}
