package config

import (
	"fmt"
	"os"
	"sync"
	"vdns/lib/api"
	"vdns/lib/auth"
	"vdns/lib/util/file"
	"vdns/lib/util/vjson"
	"vdns/lib/vlog"
)

var vdnsConfig *VdnsConfig
var rw sync.RWMutex

type Table interface {
	PrintTable() error
}

type IP struct {
	Type       string
	Enabled    bool
	Card       string
	OnCard     bool
	Api        string
	domainList []string
}

func WriteVdnsConfig(config *VdnsConfig) error {
	rw.Lock()
	defer rw.Unlock()
	f, err := os.Create(configPath)
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			vlog.Error(err)
		}
	}(f)
	if err != nil {
		return err
	}
	_, err = f.WriteString(vjson.PrettifyString(config))
	if err != nil {
		return err
	}
	return nil
}

func WriteVdnsProviderConfig(provierKey string, config *VdnsProviderConfig) error {
	rw.Lock()
	vdnsConfig.ProviderMap.Set(provierKey, config)
	rw.Unlock()
	err := WriteVdnsConfig(vdnsConfig)
	if err != nil {
		return err
	}
	return nil
}

func LoadVdnsConfig() (*VdnsConfig, error) {
	rw.Lock()
	defer rw.Unlock()
	if vdnsConfig != nil {
		return vdnsConfig, nil
	}

	read, err := file.Read(configPath)
	if err != nil {
		panic(err)
	}

	vlog.Debugf("read config:\n%v", read)

	var newConfig VdnsConfig
	err = vjson.ByteArrayConvert([]byte(read), &newConfig)
	if err != nil {
		vlog.Fatalf("read config error: %v", err)
	}
	vdnsConfig = &newConfig
	return vdnsConfig, err
}

func LoadVdnsProviderConfig(providerKey string) (*VdnsProviderConfig, error) {
	_, err := LoadVdnsConfig()
	if err != nil {
		return nil, err
	}
	rw.RLock()
	defer rw.RUnlock()
	vdnsConfigProvider := vdnsConfig.ProviderMap.Get(providerKey)
	if vdnsConfigProvider == nil {
		return nil, fmt.Errorf("vdns provider configuration not found: %v", providerKey)
	}
	return vdnsConfigProvider, nil
}

func LoadVdnsProvider(providerKey string) (api.VdnsProvider, error) {
	_, err := LoadVdnsConfig()
	if err != nil {
		return nil, err
	}
	rw.RLock()
	defer rw.RUnlock()
	for key, c := range vdnsConfig.ProviderMap {
		if key == providerKey {
			var credential auth.Credential
			if c.Provider != CloudflareProvider {
				credential = auth.NewBasicCredential(c.Ak, c.Sk)
			} else {
				credential = auth.NewTokenCredential(c.Token)
			}
			if err != nil {
				return nil, err
			}
			if providerKey == AlidnsProvider {
				return api.NewAliDNSProvider(credential), nil
			}
			if providerKey == DnspodProvider {
				return api.NewDNSPodProvider(credential), nil
			}
			if providerKey == CloudflareProvider {
				return api.NewCloudflareProvider(credential), nil
			}
			if providerKey == HuaweiDnsProvider {
				return api.NewHuaweiProvider(credential), nil
			}
		}
	}
	panic(fmt.Sprintf("provider configuration not found: %v", providerKey))
}
