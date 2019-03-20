package main

import (
	"net"
	"fmt"
	"os"
	"net/rpc"
	"rpcDemo/rpcobject"
)

func main() {
	l, err := net.Listen("tcp","localhost:5000")
	if err != nil {
		fmt.Errorf("%s","failed to create listener..ERROR:"+err.Error())
		os.Exit(1)
	}
	s := rpc.NewServer()
	agent := rpcobject.NewAgent()
	err = s.Register(agent.RpcServer)
	if err != nil {
		fmt.Errorf("%s","failed to register with rpc...ERROR:"+err.Error())
		os.Exit(1)
	}
	s.Accept(l)
}
