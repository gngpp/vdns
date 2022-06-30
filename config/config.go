package config

import (
	"errors"
	"fmt"
	"github.com/liushuochen/gotable"
	"os"
	"strings"
	"vdns/lib/api"
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
	LogDir        string                `json:"logDir,omitempty"`
	LogCompress   bool                  `json:"logCompress,omitempty"`
	LogReserveDay int                   `json:"logReserveDay,omitempty"`
	LogFilePrefix string                `json:"logFilePrefix,omitempty"`
	ProviderMap   VdnsProviderConfigMap `json:"providerMap,omitempty"`
}

func NewVdnsConfig() *VdnsConfig {
	config := VdnsConfig{
		LogDir:        defaultLogDir,
		LogCompress:   true,
		LogFilePrefix: "vdns",
		LogReserveDay: 30,
		ProviderMap:   VdnsProviderConfigMap{},
	}
	config.ProviderMap.Add(AlidnsProvider, NewProviderConfig(AlidnsProvider))
	config.ProviderMap.Add(DnspodProvider, NewProviderConfig(DnspodProvider))
	config.ProviderMap.Add(HuaweiDnsProvider, NewProviderConfig(HuaweiDnsProvider))
	config.ProviderMap.Add(CloudflareProvider, NewProviderConfig(CloudflareProvider))
	return &config
}

func (_this *VdnsConfig) SetLogDir(dir *string) {
	_this.LogDir = *dir
}

func (_this *VdnsConfig) SetLogComporess(comporess bool) {
	_this.LogCompress = comporess
}

func (_this *VdnsConfig) SetReserveDay(i int) {
	_this.LogReserveDay = i
}

func (_this *VdnsConfig) SetLogFilePrefix(prefix *string) {
	_this.LogDir = *prefix
}

func (_this *VdnsConfig) ToVlogTimeWriter() *timewriter.TimeWriter {
	return &timewriter.TimeWriter{
		Dir:           _this.LogDir,
		Compress:      _this.LogCompress,
		ReserveDay:    _this.LogReserveDay,
		LogFilePrefix: _this.LogFilePrefix,
	}
}

func (_this VdnsConfig) GetProviderConfigList() []*VdnsProviderConfig {
	l := make([]*VdnsProviderConfig, 0)
	for _, p := range _this.ProviderMap {
		l = append(l, p)
	}
	return l
}

func (_this *VdnsConfig) PrintTable() error {
	t1, err := gotable.Create("Log Path", "Log Comporess", "Log ReserveDay", "Log File Prefix")
	if err != nil {
		return err
	}

	err = t1.AddRow([]string{_this.LogDir, convert.AsStringValue(_this.LogCompress), convert.AsStringValue(_this.LogReserveDay), _this.LogFilePrefix})
	if err != nil {
		return err
	}

	t2, err := gotable.Create("Provider", "Ak", "Sk", "Token")
	if err != nil {
		return err
	}

	t3, err := gotable.Create("Provider", "Type", "Enabled", "OnCard", "Card", "Api", "DomainList")
	if err != nil {
		return err
	}

	for key, p := range _this.ProviderMap {
		if p != nil {
			err := t2.AddRow([]string{p.Provider, p.Ak, p.Sk, p.Token})
			if err != nil {
				return err
			}
			err = t3.AddRow([]string{p.Provider, p.V4.Type, convert.AsStringValue(p.V4.Enabled), convert.AsStringValue(p.V4.OnCard), p.V4.Card, p.V4.Api, strings.Join(p.V4.DomainList, ",")})
			if err != nil {
				return err
			}
			err = t3.AddRow([]string{p.Provider, p.V6.Type, convert.AsStringValue(p.V6.Enabled), convert.AsStringValue(p.V6.OnCard), p.V6.Card, p.V6.Api, strings.Join(p.V6.DomainList, ",")})
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
	fmt.Printf("---Log Config---\n%v", t1)
	fmt.Printf("---Provider Config---\n%v", t2)
	fmt.Printf("---Get Ip Config---\n%v", t3)
	return nil
}

type IP struct {
	Type       string   `json:"type,omitempty"`
	Enabled    bool     `json:"enabled,omitempty"`
	Card       string   `json:"card,omitempty"`
	OnCard     bool     `json:"onCard,omitempty"`
	Api        string   `json:"api,omitempty"`
	DomainList []string `json:"domainList,omitempty"`
}

type VdnsProviderConfig struct {
	Provider string `json:"provider"`
	Ak       string `json:"ak,omitempty"`
	Sk       string `json:"sk,omitempty"`
	Token    string `json:"token,omitempty"`
	V4       IP     `json:"v4,omitempty"`
	V6       IP     `json:"v6,omitempty"`
}

func NewProviderConfig(name string) *VdnsProviderConfig {
	return &VdnsProviderConfig{
		Provider: name,
		Ak:       "",
		Sk:       "",
		Token:    "",
		V4: IP{
			Type:       "ipv4",
			Enabled:    false,
			OnCard:     true,
			Api:        "",
			DomainList: []string{},
		},
		V6: IP{
			Type:       "ipv6",
			Enabled:    false,
			OnCard:     true,
			Api:        "",
			DomainList: []string{},
		},
	}
}

func (_this *VdnsProviderConfig) SetAk(ak *string) {
	_this.Ak = *ak
}

func (_this *VdnsProviderConfig) SetSK(sk *string) {
	_this.Sk = *sk
}

func (_this *VdnsProviderConfig) SetToken(token *string) {
	_this.Token = *token
}

func (_this *VdnsProviderConfig) PrintTable() error {
	t1, err := gotable.Create("Provider", "Ak", "Sk", "Token")
	if err != nil {
		return err
	}
	t2, err := gotable.Create("Provider", "Type", "Enabled", "OnCard", "Card", "Api", "DomainList")
	if err != nil {
		return err
	}

	err = t1.AddRow([]string{_this.Provider, _this.Ak, _this.Sk, _this.Token})
	if err != nil {
		return err
	}
	err = t2.AddRow([]string{_this.Provider, _this.V4.Type, convert.AsStringValue(_this.V4.Enabled), convert.AsStringValue(_this.V4.OnCard), _this.V4.Card, _this.V4.Api, strings.Join(_this.V4.DomainList, ",")})
	if err != nil {
		return err
	}
	err = t2.AddRow([]string{_this.Provider, _this.V6.Type, convert.AsStringValue(_this.V6.Enabled), convert.AsStringValue(_this.V6.OnCard), _this.V6.Card, _this.V6.Api, strings.Join(_this.V6.DomainList, ",")})
	if err != nil {
		return err
	}

	fmt.Printf("---Provider Config---\n%v", t1)
	fmt.Printf("---Get Ip Config---\n%v", t2)
	return nil
}

func (_this VdnsProviderConfig) ToVdnsProvider() (api.VdnsProvider, error) {
	credential := auth.NewUnifyCredential(_this.Ak, _this.Sk, _this.Token)
	if _this.Provider == AlidnsProvider {
		return api.NewAliDNSProvider(credential), nil
	}
	if _this.Provider == DnspodProvider {
		return api.NewDNSPodProvider(credential), nil
	}
	if _this.Provider == CloudflareProvider {
		return api.NewCloudflareProvider(credential), nil
	}
	if _this.Provider == HuaweiDnsProvider {
		return api.NewHuaweiProvider(credential), nil
	}
	return nil, errors.New("provider not found")
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
		vlog.Debugf("[Init] workspace directory: %s\n", workspacePath)
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
		vlog.Debugf("[Init] config file: %s\n", configPath)
	}
	if !file.Exist(defaultLogDir) {
		err = file.MakeDir(defaultLogDir)
		if err != nil {
			panic("[Error] creating log dir: " + defaultLogDir + " error: " + err.Error())
		}
	}
}
