package main

import (
	"container/list"
	"fmt"
	"vdns/lib/standard"
)

func main() {
	println(standard.Record.OfType("A"))
	isOK, recordType := standard.Record.OfType("SB")
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
