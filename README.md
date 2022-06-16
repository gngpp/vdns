<p align="center">
	<a target="_blank" href="https://github.com/gngpp/vdns/blob/main/LICENSE">
		<img src="https://img.shields.io/badge/license-MIT-blue.svg"/>
	</a>
	<a target="_blank" href="https://go.dev/">
		<img src="https://img.shields.io/github/go-mod/go-version/gngpp/vdns"/>
	</a>
	<a target="_blank" href="https://github.com/gngpp/vdns/actions">
		<img src="https://github.com/gngpp/vdns/actions/workflows/sync.yaml/badge.svg"/>
	</a>
<!-- 	<a target="_blank" href="https://github.com/gngpp/vdns/releases/latest">
		<img alt="GitHub Release" src="https://img.shields.io/github/v/release/gngpp/vdns.svg?logo=github">
	</a> -->
</p>

# vdns
`vdns`支持多平台`DNS`解析操作, 同时以服务方式支持`DDNS`，支持多种解析记录类型：A、AAAA、NS、MX、CNAME、TXT、SRV、CA、REDIRECT_URL、FORWARD_URL


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
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### License

- [MIT License](https://raw.githubusercontent.com/gngpp/vdns/main/LICENSE)

### JetBrains 开源证书支持

> `vdns` 项目一直以来都是在 JetBrains 公司旗下的 GoLand 集成开发环境中进行开发，基于 **free JetBrains Open Source license(s)** 正版免费授权，在此表达我的谢意。

<a href="https://www.jetbrains.com/?from=gnet" target="_blank"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" width="250" align="middle"/></a>
