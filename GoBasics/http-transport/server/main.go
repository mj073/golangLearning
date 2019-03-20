package main

import (

	"net"
	"fmt"
	"bytes"
)

func main(){
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		// handle error
	}
	//conn, err := ln.Accept()
	//fmt.Printf("%T",conn)
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Printf("conn: %v\n",conn)
		go handleConnection(conn)
	}
}

func handleConnection(con net.Conn){
	fmt.Println("remote address: ",con.RemoteAddr().String())
	read := make([]byte,2048)
	//buff := bytes.NewBuffer(read)
	_,err := con.Read(read)
	if err != nil{
		fmt.Println("error while reading: ",err)
	}else{
		fmt.Println("reading: ",string(read))
	}

}

