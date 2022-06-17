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

type VdnsConfigProviderMap map[string]*VdnsConfigProvider

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v VdnsConfigProviderMap) Get(key string) *VdnsConfigProvider {
	if v == nil {
		return nil
	}
	vs := v[key]
	return vs
}

// Set sets the key to value. It replaces any existing
// values.
func (v VdnsConfigProviderMap) Set(key string, value *VdnsConfigProvider) {
	v[key] = value
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (v VdnsConfigProviderMap) Add(key string, value *VdnsConfigProvider) {
	config := v.Get(key)
	if config != nil {
		return
	}
	v.Set(key, value)
}

// Del deletes the values associated with key.
func (v VdnsConfigProviderMap) Del(key string) {
	delete(v, key)
}

// Has checks whether a given key is set.
func (v VdnsConfigProviderMap) Has(key string) bool {
	_, ok := v[key]
	return ok
}

type VdnsConfig struct {
	ConfigsMap VdnsConfigProviderMap
}

func NewVdnsConfig() *VdnsConfig {
	config := VdnsConfig{
		ConfigsMap: VdnsConfigProviderMap{},
	}
	config.ConfigsMap.Add(AlidnsProvider, NewProviderConfig(AlidnsProvider))
	config.ConfigsMap.Add(DnspodProvider, NewProviderConfig(DnspodProvider))
	config.ConfigsMap.Add(HuaweiDnsProvider, NewProviderConfig(HuaweiDnsProvider))
	config.ConfigsMap.Add(CloudflareProvider, NewProviderConfig(CloudflareProvider))
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
	if key != CloudflareProvider {
		return auth.NewBasicCredential(*get.Ak, *get.Sk), nil
	} else {
		return auth.NewTokenCredential(*get.Token), nil
	}
}

type VdnsConfigProvider struct {
	Provider *string `json:"provider"`
	Ak       *string `json:"ak,omitempty"`
	Sk       *string `json:"sk,omitempty"`
	Token    *string `json:"token,omitempty"`
}

func NewProviderConfig(name string) *VdnsConfigProvider {
	return &VdnsConfigProvider{
		Provider: &name,
		Ak:       strs.String(""),
		Sk:       strs.String(""),
		Token:    strs.String(""),
	}
}

type DDNSConfig struct {
}

func ReadConfig() (*VdnsConfig, error) {
	read, err := file.Read(configPath)
	if err != nil {
		panic(err)
	}
	vlog.Debugf("read config:\n%v", read)
	var entity VdnsConfig
	err = vjson.ByteArrayConvert([]byte(read), &entity)
	if err != nil {
		vlog.Fatalf("read config error: %v", err)
	}
	return &entity, err
}

func WriteConfig(config *VdnsConfig) error {
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
		panic("[Error] system error")
	}

	//goland:noinspection SpellCheckingInspection
	WorkingDir := strs.Concat(dir, "/.vdns")
	if !file.Exist(WorkingDir) {
		err = file.MakeDir(WorkingDir)
		if err != nil {
			panic("[Error] creating working " + WorkingDir + " directory error: " + err.Error())
		}
		vlog.Infof("[Init] working directory: %s\n", WorkingDir)
	}

	configPath = strs.Concat(WorkingDir, "/config.json")
	if !file.Exist(configPath) {
		create, err := os.Create(configPath)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				println(err)
			}
		}(create)
		if err != nil {
			panic("[Error] creating " + configPath + " config create error: " + err.Error())
		}
		config := NewVdnsConfig()
		_, err = create.WriteString(vjson.PrettifyString(config))
		if err != nil {
			panic("[Error] initializing " + configPath + " config create error: " + err.Error())
		}
		vlog.Infof("[Init] config file: %s\n", configPath)
	}
}
