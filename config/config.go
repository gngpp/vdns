package config

import (
	"errors"
	"os"
	"vdns/lib/auth"
	"vdns/lib/homedir"
	"vdns/lib/util/file"
	"vdns/lib/util/strs"
	"vdns/lib/util/vjson"
	"vdns/lib/vlog"
)

var configPath string

type Configs map[string]*DNSConfig

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v Configs) Get(key string) *DNSConfig {
	if v == nil {
		return nil
	}
	vs := v[key]
	return vs
}

// Set sets the key to value. It replaces any existing
// values.
func (v Configs) Set(key string, value *DNSConfig) {
	v[key] = value
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (v Configs) Add(key string, value *DNSConfig) {
	config := v.Get(key)
	if config != nil {
		return
	}
	v.Set(key, value)
}

// Del deletes the values associated with key.
func (v Configs) Del(key string) {
	delete(v, key)
}

// Has checks whether a given key is set.
func (v Configs) Has(key string) bool {
	_, ok := v[key]
	return ok
}

type Config struct {
	ConfigsMap Configs
}

func NewConfig() *Config {
	config := Config{
		ConfigsMap: Configs{},
	}
	config.ConfigsMap.Add(ALIDNS_PROVIDER, NewDNSConfig(ALIDNS_PROVIDER))
	config.ConfigsMap.Add(DNSPOD_PROVIDER, NewDNSConfig(DNSPOD_PROVIDER))
	config.ConfigsMap.Add(HUAWERI_DNS_PROVIDER, NewDNSConfig(HUAWERI_DNS_PROVIDER))
	config.ConfigsMap.Add(CLOUDFLARE_PROVIDER, NewDNSConfig(CLOUDFLARE_PROVIDER))
	return &config
}

func ReadCredentials(key string) (auth.Credential, error) {
	config, err := ReadConfig()
	if err != nil {
		return nil, err
	}
	get := config.ConfigsMap.Get(key)
	if get == nil {
		return nil, errors.New("init credentials not found")
	}
	if key != CLOUDFLARE_PROVIDER {
		return auth.NewBasicCredential(*get.Ak, *get.Sk), nil
	} else {
		return auth.NewTokenCredential(*get.Token), nil
	}
}

type DNSConfig struct {
	Provider *string `json:"provider"`
	Ak       *string `json:"ak,omitempty"`
	Sk       *string `json:"sk,omitempty"`
	Token    *string `json:"token,omitempty"`
}

func NewDNSConfig(name string) *DNSConfig {
	return &DNSConfig{
		Provider: &name,
		Ak:       strs.String(""),
		Sk:       strs.String(""),
		Token:    strs.String(""),
	}
}

type DDNSConfig struct {
}

func ReadConfig() (*Config, error) {
	read, err := file.Read(configPath)
	if err != nil {
		panic(err)
	}
	vlog.Debugf("read config:\n%v", read)
	var entity Config
	err = vjson.ByteArrayConvert([]byte(read), &entity)
	if err != nil {
		vlog.Fatalf("read config error: %v", err)
	}
	return &entity, err
}

func WriteConfig(config *Config) error {
	open, err := os.Create(configPath)
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			vlog.Error(err)
		}
	}(open)
	if err != nil {
		return err
	}
	_, err = open.WriteString(vjson.PrettifyString(config))
	if err != nil {
		return err
	}
	return nil
}

func init() {

	dir, err := homedir.Dir()
	if err != nil {
		panic("system error")
	}

	//goland:noinspection SpellCheckingInspection
	WorkingDir := strs.Concat(dir, "/.vdns")
	if !file.Exist(WorkingDir) {
		err = file.MakeDir(WorkingDir)
		if err != nil {
			panic("creating working " + WorkingDir + " directory error: " + err.Error())
		}
	}

	configPath = strs.Concat(WorkingDir, "/config.json")
	if !file.Exist(configPath) {
		create, err := os.Create(configPath)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				vlog.Error(err)
			}
		}(create)
		if err != nil {
			panic("creating " + configPath + " config create error: " + err.Error())
		}
		config := NewConfig()
		_, err = create.WriteString(vjson.PrettifyString(config))
		if err != nil {
			panic("initializing " + configPath + " config create error: " + err.Error())
		}
	}
}
