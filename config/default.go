package config

import (
	"fmt"
	"os"
	"sync"
	"vdns/lib/api"
	"vdns/lib/auth"
	"vdns/lib/util/file"
	"vdns/lib/util/strs"
	"vdns/lib/util/vjson"
	"vdns/lib/vlog"
)

var conf *VdnsConfig
var rw sync.RWMutex

type Table interface {
	PrintTable() error
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
	conf = config
	return nil
}

func WriteVdnsProviderConfig(config *VdnsProviderConfig) error {
	rw.Lock()
	if strs.IsEmpty(config.Provider) {
		return fmt.Errorf("provider key cnanot been empty: %v", config.Provider)
	}
	conf.ProviderMap.Set(config.Provider, config)
	rw.Unlock()
	err := WriteVdnsConfig(conf)
	if err != nil {
		return err
	}
	return nil
}

func ReadVdnsConfig() (*VdnsConfig, error) {
	rw.Lock()
	defer rw.Unlock()
	if conf != nil {
		return conf, nil
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
	conf = &newConfig
	return conf, err
}

func ReadVdnsProviderConfig(providerKey string) (*VdnsProviderConfig, error) {
	c, err := ReadVdnsConfig()
	if err != nil {
		return nil, err
	}
	rw.RLock()
	defer rw.RUnlock()
	vdnsConfigProvider := c.ProviderMap.Get(providerKey)
	if vdnsConfigProvider == nil {
		return nil, fmt.Errorf("vdns provider configuration not found: %v", providerKey)
	}
	return vdnsConfigProvider, nil
}

func ReadVdnsProvider(providerKey string) (api.VdnsProvider, error) {
	c, err := ReadVdnsConfig()
	if err != nil {
		return nil, err
	}
	rw.RLock()
	defer rw.RUnlock()
	for key, c := range c.ProviderMap {
		if key == providerKey {
			credential := auth.NewUnifyCredential(c.Ak, c.Sk, c.Token)

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
