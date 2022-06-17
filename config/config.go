package config

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
	"os"
	"vdns/lib/auth"
	"vdns/lib/homedir"
	"vdns/lib/util/convert"
	"vdns/lib/util/file"
	"vdns/lib/util/strs"
	"vdns/lib/util/vjson"
	"vdns/lib/vlog"
	"vdns/lib/vlog/timewriter"
)

var configPath string
var workspacePath string
var defaultLogDir string

type VdnsProviderConfigMap map[string]*VdnsProviderConfig

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v VdnsProviderConfigMap) Get(key string) *VdnsProviderConfig {
	if v == nil {
		return nil
	}
	vs := v[key]
	return vs
}

// Set sets the key to value. It replaces any existing
// values.
func (v VdnsProviderConfigMap) Set(key string, value *VdnsProviderConfig) {
	v[key] = value
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (v VdnsProviderConfigMap) Add(key string, value *VdnsProviderConfig) {
	config := v.Get(key)
	if config != nil {
		return
	}
	v.Set(key, value)
}

// Del deletes the values associated with key.
func (v VdnsProviderConfigMap) Del(key string) {
	delete(v, key)
}

// Has checks whether a given key is set.
func (v VdnsProviderConfigMap) Has(key string) bool {
	_, ok := v[key]
	return ok
}

type VdnsConfig struct {
	Dir           string
	Compress      bool
	ReserveDay    int
	LogFilePrefix string
	ProviderMap   VdnsProviderConfigMap
}

func (_this *VdnsConfig) ToVlogTimeWriter() *timewriter.TimeWriter {
	return &timewriter.TimeWriter{
		Dir:           _this.Dir,
		Compress:      _this.Compress,
		ReserveDay:    _this.ReserveDay,
		LogFilePrefix: _this.LogFilePrefix,
	}
}

func (_this *VdnsConfig) PrintTable() error {
	t1, err := gotable.Create("Log Path", "Log Comporess", "Log ReserveDay", "Log File Prefix")
	if err != nil {
		return err
	}
	err = t1.AddRow([]string{_this.Dir, convert.AsStringValue(_this.Compress), convert.AsStringValue(_this.ReserveDay), _this.LogFilePrefix})

	t2, err := gotable.CreateByStruct(new(VdnsProviderConfig))
	if err != nil {
		return err
	}
	for key, p := range _this.ProviderMap {
		if p != nil {
			err := t2.AddRow([]string{*p.Provider, *p.Ak, *p.Sk, *p.Token})
			if err != nil {
				return err
			}
		} else {
			err := t1.AddRow([]string{key})
			if err != nil {
				return err
			}
		}
	}
	if err != nil {
		return err
	}
	fmt.Println(t1)
	fmt.Print(t2)
	return err
}

func NewVdnsConfig() *VdnsConfig {
	config := VdnsConfig{
		Dir:           defaultLogDir,
		Compress:      true,
		LogFilePrefix: "vdns",
		ReserveDay:    30,
		ProviderMap:   VdnsProviderConfigMap{},
	}
	config.ProviderMap.Add(AlidnsProvider, NewProviderConfig(AlidnsProvider))
	config.ProviderMap.Add(DnspodProvider, NewProviderConfig(DnspodProvider))
	config.ProviderMap.Add(HuaweiDnsProvider, NewProviderConfig(HuaweiDnsProvider))
	config.ProviderMap.Add(CloudflareProvider, NewProviderConfig(CloudflareProvider))
	return &config
}

type VdnsProviderConfig struct {
	Provider *string `json:"provider"`
	Ak       *string `json:"ak,omitempty"`
	Sk       *string `json:"sk,omitempty"`
	Token    *string `json:"token,omitempty"`
}

func (_this *VdnsProviderConfig) loadCredential() (auth.Credential, error) {
	if *_this.Provider != CloudflareProvider {
		return auth.NewBasicCredential(*_this.Ak, *_this.Sk), nil
	} else {
		return auth.NewTokenCredential(*_this.Token), nil
	}
}

func (_this *VdnsProviderConfig) PrintTable() (*table.Table, error) {
	t, err := gotable.CreateByStruct(new(VdnsProviderConfig))
	if err != nil {
		return nil, err
	}
	err = t.AddRow([]string{*_this.Provider, *_this.Ak, *_this.Sk, *_this.Token})
	if err != nil {
		return nil, err
	}
	return t, nil
}

func NewProviderConfig(name string) *VdnsProviderConfig {
	return &VdnsProviderConfig{
		Provider: &name,
		Ak:       strs.String(""),
		Sk:       strs.String(""),
		Token:    strs.String(""),
	}
}

func init() {

	dir, err := homedir.Dir()
	if err != nil {
		panic("[Error] system error")
	}

	//goland:noinspection SpellCheckingInspection
	workspacePath = strs.Concat(dir, "/.vdns")
	if !file.Exist(workspacePath) {
		err = file.MakeDir(workspacePath)
		if err != nil {
			panic("[Error] creating workspace " + workspacePath + " directory error: " + err.Error())
		}
		vlog.Infof("[Init] workspace directory: %s\n", workspacePath)
	}

	configPath = strs.Concat(workspacePath, "/config.json")
	defaultLogDir = strs.Concat(workspacePath, "/logs")
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
		defaultLogDir = config.Dir
		vlog.Infof("[Init] config file: %s\n", configPath)
	}
	if !file.Exist(defaultLogDir) {
		err = file.MakeDir(defaultLogDir)
		if err != nil {
			panic("[Error] creating log dir: " + defaultLogDir + " error: " + err.Error())
		}
	}
}
