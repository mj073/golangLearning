package main

import (
	"log"
	"net/rpc"
	"rpcDemo/rpcobject"
)

func main() {
	//make connection to rpc server
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	//make arguments object
	args := &rpcobject.Args{
		A: 2,
		B: 3,
	}
	//this will store returned result
	var result rpcobject.Result
	//call remote procedure with args
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	//we got our result in result
	log.Printf("%d*%d=%d\n", args.A, args.B, result)

	err = client.Call("Arith.Add", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	//we got our result in result
	log.Printf("%d+%d=%d\n", args.A, args.B, result)
}
