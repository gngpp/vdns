<p align="center">
	<a target="_blank" href="https://github.com/gngpp/vdns/blob/main/LICENSE">
		<img src="https://img.shields.io/badge/license-WTFPL-blue.svg"/>
	</a>
	<a target="_blank" href="https://go.dev/">
		<img src="https://img.shields.io/github/go-mod/go-version/gngpp/vdns"/>
	</a>
	<a target="_blank" href="https://github.com/gngpp/vdns/actions/workflows/sync.yaml">
		<img src="https://github.com/gngpp/vdns/actions/workflows/sync.yaml/badge.svg"/>
	</a>
	<a target="_blank" href="https://github.com/gngpp/vdns/actions/workflows/release.yml">
		<img src="https://github.com/gngpp/vdns/actions/workflows/release.yml/badge.svg"/>
	</a>
	<a target="_blank" href="https://github.com/gngpp/vdns/releases/latest">
		<img alt="GitHub Release" src="https://img.shields.io/github/v/release/gngpp/vdns.svg?logo=github">
	</a>
</p>

# vdns
`vdns`支持多平台`DNS`解析操作, 同时以服务方式支持`DDNS`，
- 支持阿里云、腾讯云、华为云、Cloudflare等平台的DNS记录解析
- 支持多种解析记录类型：A、AAAA、NS、MX、CNAME、TXT、SRV、CA、REDIRECT_URL、FORWARD_URL


### Terminal CLI
```shell
NAME:
   vdns - This is A tool that supports multi-DNS service provider resolution operations.

USAGE:
   vdns [global options] command [command options] [arguments...]

VERSION:
   v1.0

COMMANDS:
   config             Configure DNS service provider access key pair
   server             Use vdns server (support DDNS)
   resolve            Resolve DNS records
   provider           Support providers
   record             Support record types
   card               Print available network card information
   print-ip-api, pia  Print search ip request api list
   test-ip-api, tia   Test the API for requesting query ip
   request            Request Api (only support get method)
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d    enable debug mode (default: false)
   --help, -h     show help (default: false)
```

- 使用DDNS
> 下载二进制包
```shell
$ wget https://github.com/gngpp/vdns/releases/download/v1.0/linux_amd64_vdns.tar.gz && tar -xzvf linux_amd64_vdns.tar.gz
```
> 直接运行
```shell
$ ./vdns server exec
```
> 后台daemon运行，支持 `Windows XP+`, `Linux/(systemd | Upstart | SysV)`, `OSX/Launchd`
```shell
# 安装运行
./vdns server install
./vdns server start

# 卸载
./vdns server stop
./vdns server uninstall
```
### License

- [MIT License](https://raw.githubusercontent.com/gngpp/vdns/main/LICENSE)

### JetBrains 开源证书支持

> `vdns` 项目一直以来都是在 JetBrains 公司旗下的 GoLand 集成开发环境中进行开发，基于 **free JetBrains Open Source license(s)** 正版免费授权，在此表达我的谢意。

<a href="https://www.jetbrains.com/?from=gnet" target="_blank"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" width="250" align="middle"/></a>
