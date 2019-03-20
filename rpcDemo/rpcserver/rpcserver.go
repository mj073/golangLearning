package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpcDemo/rpcobject"
	"os"
	"runtime/trace"
	// _ "net/http/pprof"
	"time"
	"fmt"
	"errors"
)

func main() {
	// for pprof
	//go func() {
	//	log.Println(http.ListenAndServe("192.168.56.100:6060", nil))
	//}()

	f, err := os.Create("rpcServerTrace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("starting trace...")
	if err := trace.Start(f); err != nil {
		panic(err)
	}
	defer trace.Stop()

	//register Arith object as a service
	arith := new(rpcobject.Arith)
	err = rpc.Register(arith)
	if err != nil {
		log.Fatalf("Format of service Arith isn't correct. %s", err)
	}
	rpc.HandleHTTP()
	//start listening for messages on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
	}
	log.Println("Serving RPC handler via HTTP..")
	go func() {
		defer func() {
			if err := recover(); err != nil{
				fmt.Println("recovering from panic...")
				fmt.Println("stopping trace...")
				trace.Stop()
				if f != nil{
					f.Close()
				}
			}else {
				fmt.Println("error nil...")
			}

		}()
		if _,ok := <-time.After(time.Second * 40); ok{
			panic(errors.New("forceful panic.."))
		}
		fmt.Println("will this get executed...???")
	}()
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
