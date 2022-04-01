# vdns
`vdns`支持多平台`DNS`解析操作, 同时以服务方式支持`DDNS`，支持多种解析记录类型：A、AAAA、NS、MX、CNAME、TXT、SRV、CA、REDIRECT_URL、FORWARD_URL


### Terminal CLI
```shell
❯ go run main.go
NAME:
   vdns - A tool that supports multi-DNS service provider resolution operations

USAGE:
   vdns [global options] command [command options] [arguments...]

VERSION:
   v1.0

COMMANDS:
   show, s     Show vdns information
   config, c   Configure dns service provider access key pair
   resolve, r  Resolving dns records
   server      Use vdns server (support DDNS)
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```