package main

import (
	"net"
	"fmt"
	"os"
	"net/rpc"
	"github.com/golangLearning/bankApp/common"
)

func main(){
	listener, err := net.Listen("unix","@"+"bankServer")
	if err != nil {
		fmt.Errorf("%s","failed to listen on unix socket..ERROR:",err)
		os.Exit(1)
	}
	bankTransact := common.NewBankTransaction()
	rpc.Register(bankTransact)
	go rpc.Accept(listener)
}
