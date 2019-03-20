package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	go func() {
		for {
			intf, _ := net.InterfaceByName("eth6")
			addrs, _ := intf.Addrs()
			if addrs[0].String() != "192.168.16.103" {
				fmt.Println("address changed...:", addrs[0])
			}
			time.Sleep(time.Second * 2)
		}
	}()
	fmt.Println(http.ListenAndServe("192.168.16.103:5000", nil))
}
