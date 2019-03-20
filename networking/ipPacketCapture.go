package main

import (
	"net"
	"log"
)

func main(){
	service := "192.168.56.100:8888"
	IPAddr, err := net.ResolveIPAddr("ip4", service)
	//l, err := net.ListenIP("ip:tcp",&net.IPAddr{IP: net.IPv4(172,17,222,89)})
	l, err := net.ListenIP("ip:tcp",IPAddr)
	if err != nil {
		log.Fatal("failed to listen TCP...ERROR:",err)
	}
	for {
		log.Println("listening for TCP")
		b := make([]byte, 1000)
		//con, err := l.Accept()
		//if err != nil {
		//	log.Fatal("failed to accept the connection..ERROR:", err)
		//}
		_, err = l.Read(b)
		if err != nil {
			log.Fatal("failed to read from conn..ERROR:", err)
		}
		log.Print("bytes read:", string(b))
	}
}
