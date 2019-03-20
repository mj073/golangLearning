package main

import (
	"net"
	"log"
)

func main(){
	service := "192.168.56.100:8888"
	TCPAddr, err := net.ResolveTCPAddr("tcp4", service)
	l, err := net.ListenTCP("tcp4",TCPAddr)
	//l, err := net.ListenTCP("tcp4",&net.TCPAddr{IP: net.IPv4(172,17,222,89), Port: 8888,Zone: ""})
	if err != nil {
		log.Fatal("failed to listen TCP...ERROR:",err)
	}
	for {
		log.Println("listening for TCP")
		b := make([]byte, 1000)
		con, err := l.Accept()
		if err != nil {
			log.Fatal("failed to accept the connection..ERROR:", err)
		}
		_, err = con.Read(b)
		if err != nil {
			log.Fatal("failed to read from conn..ERROR:", err)
		}
		log.Print("bytes read:", string(b))
		con.Close()
	}
}


