package main

import (
	"rpcDemo/rpcobject"
	"net/rpc"
	"net"
	"fmt"
	"os"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main(){
	// for pprof
	go func() {
		log.Println(http.ListenAndServe("192.168.56.100:5555", nil))
	}()
	c := rpcobject.NewRPCCache()
	rpc.Register(c)

	l, err := net.Listen("tcp","192.168.56.100:2000")
	if err != nil{
		fmt.Println("failed to listen on localhost:2000..>ERROR:",err)
		os.Exit(1)
	}
	rpc.Accept(l)

}
