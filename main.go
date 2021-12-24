package main

import (
	"container/list"
	"fmt"
	"vdns/lib/common"
)

func main() {
	println(common.Record.OfType("A"))
	isOK, recordType := common.Record.OfType("SB")
	if isOK {
		fmt.Println(recordType)
	}
	fmt.Println(recordType == "")
	list := list.New()
	list.PushBack("A")
	list.PushBack("B")

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
