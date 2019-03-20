package main

import (
	"fmt"
	"net"
)
//run as root then only it works
func main() {
	protocol := "tcp"
	netaddr, _ := net.ResolveIPAddr("ip4", "127.0.0.1")
	//for {
		conn, _ := net.ListenIP("ip4:" + protocol, netaddr)
		buf := make([]byte, 1024)
		numRead, _, _ := conn.ReadFrom(buf)
		fmt.Printf("% X\n", buf[:numRead])
		//conn.Close()
	//}
}

