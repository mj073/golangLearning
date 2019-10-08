package main

import (
	"fmt"
	"github.com/golangLearning/golangTraining/bankApp/common"
	"net"
	"net/rpc"
	"os"
)

func main(){
	listener, err := net.Listen(common.ProtoType, common.Address)
	if err != nil {
		fmt.Errorf("%s","failed to listen on unix socket..ERROR:",err)
		os.Exit(1)
	}
	common.RegisterAllGob()
	bankTransact := common.NewBankTransaction()
	err = rpc.Register(bankTransact)
	if err != nil {
		fmt.Errorf("%s:%v","failed to register rpc object..error",err)
		os.Exit(1)
	}
	rpc.Accept(listener)
}
