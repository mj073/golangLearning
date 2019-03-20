package main

import (
	"net"
	"fmt"
)

func main() {
	pktCon,err := net.ListenPacket("udp",":9090")
	if err != nil{
		fmt.Println("error while listening..ERROR:",err)
		return
	}
	fmt.Println("local address:",pktCon.LocalAddr().String())

	//pktCon.SetDeadline(time.Now().Add(50 * time.Second))
	buf := make([]byte,2048)
	for {

		numread, _, err := pktCon.ReadFrom(buf)
		if err != nil {
			fmt.Println("error while reading from packet..ERROR:", err)
			return
		}
		fmt.Printf("% X\n", buf[:numread])
	}
}
