module github.com/vdns/cli

go 1.17

replace (
	github.com/zf1976/vdns/lib => ../lib
	github.com/zf1976/vdns/testing => ../testing
)

require github.com/zf1976/vdns/lib v0.0.0-00010101000000-000000000000
