package main

import (
	"net/rpc"
	"net"
	"fmt"
	"os"
	"rpcDemo/rpcobject"
)

func main() {
	conn, err := net.Dial("tcp","localhost:5000")
	if err != nil {
		fmt.Errorf("%s","failed to create connection...ERROR:"+err.Error())
		os.Exit(1)
	}
	rpcClient := rpc.NewClient(conn)
	if rpcClient != nil {
		var a interface{}
		err = rpcClient.Call("RpcServer.GetAgent","",a)
		if err != nil {
			fmt.Errorf("%s","failed to call service method...ERROR:"+err.Error())
			os.Exit(1)
		}
		fmt.Println("agent:",a.(*rpcobject.Agent))
	}
}