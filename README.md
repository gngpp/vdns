## vdns
`vdns`支持多平台域名解析, 同时以服务方式支持DDNS
并且支持多种解析记录类型：A、AAAA、NS、MX、CNAME、TXT、SRV、CA、REDIRECT_URL、FORWARD_URL


### 以`Terminal CLI`方式
```shell
❯ go run cli.go
NAME:
   vdns - vdns is a tool that supports multi-DNS service provider resolution operations.

USAGE:
   vdns [global options] command [command options] [arguments...]

COMMANDS:
   show, s     Show vdns information.
   config, c   Configure dns service provider access key pair.
   resolve, r  Resolving DNS records.
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

- 显示vdns相关信息
```shell
❯ go run cli.go show
NAME:
   vdns show - Show vdns information.

USAGE:
   vdns show command [command options] [arguments...]

COMMANDS:
   provider, p  support providers.
   record, r    supports record types.
   help, h      Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```

- DNS服务商配置
```shell
❯ go run cli.go config --help
NAME:
   vdns config - Configure dns service provider access key pair.

USAGE:
   vdns config command [command options] [arguments...]

COMMANDS:
   alidns      Configure alidns access key pair.
   dnspod      Configure dnspod access key pair.
   huaweidns   Configure huaweidns access key pair.
   cloudflare  Configure cloudflare access key pair.
   cat         Print all dns configuration.
   help, h     Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
   
❯ go run cli.go config alidns --help
NAME:
   vdns config alidns - Configure alidns access key pair.

USAGE:
   vdns config alidns command [command options] [arguments...]

COMMANDS:
   cat      Print dns provider configuration.
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --ak value     api accessKey.
   --sk value     api secretKey.
   --token value  api token.
   --help, -h     show help (default: false)
```

- 域名解析操作
```shell
❯ go run cli.go resolve alidns
NAME:
   vdns resolve alidns - resolve AliDNS DNS records.

USAGE:
   vdns resolve alidns command [command options] [arguments...]

COMMANDS:
   search, s  describe AliDNS DNS records.
   create, c  create AliDNS DNS record.
   update, u  update AliDNS DNS record.
   delete, d  delete AliDNS DNS record.
   help, h    Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
   
❯ go run cli.go resolve alidns search --help
NAME:
   vdns resolve alidns search - describe AliDNS DNS records.

USAGE:
   vdns resolve alidns search [command options] [arguments...]

OPTIONS:
   --ps value      page size. (default: 5)
   --pn value      page number. (default: 1)
   --domain value  record domain.
   --type value    record type.
   --rk value      the keywords recorded by the host, (fuzzy matching before and after) pattern search, are not case-sensitive.
   --vk value      the record value keyword (fuzzy match before and after) pattern search, not case-sensitive.
   --help, -h      show help (default: false)
```