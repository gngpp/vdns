<p align="center">
	<a target="_blank" href="https://github.com/zf1976/vdns/blob/main/LICENSE">
		<img src="https://img.shields.io/badge/license-MIT-blue.svg"/>
	</a>
	<a target="_blank" href="https://go.dev/">
		<img src="https://img.shields.io/github/go-mod/go-version/gngpp/vdns"/>
	</a>
	<a target="_blank" href="https://github.com/zf1976/vdns/actions">
		<img src="https://github.com/zf1976/vdns/actions/workflows/sync.yaml/badge.svg"/>
	</a>
<!-- 	<a target="_blank" href="https://github.com/zf1976/vdns/releases/latest">
		<img alt="GitHub Release" src="https://img.shields.io/github/v/release/zf1976/vdns.svg?logo=github">
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
   config             Configure dns service provider access key pair
   resolve            Resolving dns records
   server             Use vdns server (support DDNS)
   provider           Support providers
   record             Support record types
   print-ip-api, pia  Print search ip request api list
   test-ip-api, tia   Test the API for requesting query ip
   request            Request Api (only support get method)
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

```

## License

`vdns`是在 MIT 许可下发布的，
提供许可证副本在 [LICENSE](./LICENSE) 文件中。
