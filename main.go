package main

import "vdns/lib/common"

func main() {
	println(common.Record.OfType("A"))
	println(common.Record.OfType("SB"))
}
