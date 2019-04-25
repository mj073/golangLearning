package main

import (
	"fmt"
	"github.com/golangLearning/golangTraining/rpcDemo/rpcobject"
	"net"
	"net/rpc"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		fmt.Println("failed to dial localhost:2000...ERROR:", err)
		os.Exit(1)
	}
	defer conn.Close()
	client := rpc.NewClient(conn)
	defer client.Close()

	fmt.Println("args:", os.Args)
	switch os.Args[1] {
	case "set":
		var ack bool
		err = client.Call("RPCCacheService.Put", &rpcobject.CacheItem{os.Args[2], os.Args[3]}, &ack)
		if err != nil {
			fmt.Println("RPC server call failed...ERROR:", err)
		}
		fmt.Println(ack)
	case "get":
		cache := &rpcobject.CacheItem{}
		err = client.Call("RPCCacheService.Get", os.Args[2], cache)
		if err != nil {
			fmt.Println("RPC server call failed...ERROR:", err)
		} else {
			fmt.Println(cache.Value)
		}
	case "del":
		var ack bool
		err = client.Call("RPCCacheService.Delete", os.Args[2], &ack)
		if err != nil {
			fmt.Println("RPC server call failed...ERROR:", err)
		}
		fmt.Println(ack)
	case "clear":
		var ack bool
		err = client.Call("RPCCacheService.Clear", true, &ack)
		if err != nil {
			fmt.Println("RPC server call failed...ERROR:", err)
		}
		fmt.Println(ack)
	default:
		fmt.Errorf("unkown argument: %s", os.Args[1])
		os.Exit(1)
	}

}
