package main

import (
	"fmt"
	"github.com/golangLearning/rpcDemo/rpcobject"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/rpc"
	"os"
)

func main() {
	// for pprof
	go func() {
		log.Println(http.ListenAndServe("localhost:5555", nil))
	}()
	c := rpcobject.NewRPCCache()
	rpc.Register(c)

	l, err := net.Listen("tcp", "192.168.1.105:2000")
	if err != nil {
		fmt.Println("failed to listen on localhost:2000..>ERROR:", err)
		os.Exit(1)
	}
	/*	for {
		co, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR:",err)
			continue
		}
		rpc.ServeConn(co)
	}*/
	rpc.Accept(l)

}
