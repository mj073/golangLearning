package main

import (
	"net/rpc"
	"net"
	"fmt"
	"os"
	"flag"
	"github.com/golangLearning/bankApp/common"
)
var (
	username = flag.String("user","admin","<username>")
	password = flag.String("pass","","<password>")
	transactionType = flag.String("type","checkbalance","type [createaccount|deposit|withdraw|checkbalance|customerdetails]")
	customerId = flag.Uint("customerId",0,"customerId <int>")

)
func main(){
	conn, err := net.Dial("unix","@"+"bankServer")
	if err != nil{
		fmt.Println("failed to dial unix socket...ERROR:",err)
		os.Exit(1)
	}
	client := rpc.NewClient(conn)
	flag.Parse()
	switch t := *transactionType {
	case "createaccount":
		custDetails := common.CustomerDetails{}
		fmt.Println("Enter your personal info:")
		fmt.Print("Name: ")
		fmt.Scanf("%s",custDetails.Name)
		fmt.Print("Age: ")
		fmt.Scanf("%d\n",custDetails.Age)
		fmt.Print("Pan: ")
		fmt.Scanf("%s\n",custDetails.Pan)

		request := common.TransactionRequest{Type: common.CreateAccount, Details: custDetails}
		response := common.TransactionResponse{}
		err = client.Call("BankTransaction.Transact",&request,&response)
		if err != nil{
			fmt.Println("RPC server call failed...ERROR:",err)
		}
		//presentResponse(response)
		fmt.Println(response)
	case "customerdetails":
		if !(username == "admin" && password == "mahesh123"){
			fmt.Errorf("%s","invalid username/password")
			return
		}
		details := common.CustomerDetailsRequest{}
		if customerId != 0{
			details.CustomerId = customerId
		}
		request := common.TransactionRequest{Type: common.CustomerDetails, Details: details}
		response := common.TransactionResponse{}
		err = client.Call("BankTransaction.Transact",&request,&response)
		if err != nil{
			fmt.Println("RPC server call failed...ERROR:",err)
		}
		fmt.Println(response)
	case "deposit":
		cache := &common.CacheItem{}
		err = client.Call("RPCCacheService.Get",os.Args[2],cache)
		if err != nil{
			fmt.Println("RPC server call failed...ERROR:",err)
		}else {
			fmt.Println(cache.Value)
		}
	case "withdraw":
		var ack bool
		err = client.Call("RPCCacheService.Delete",os.Args[2],&ack)
		if err != nil{
			fmt.Println("RPC server call failed...ERROR:",err)
		}
		fmt.Println(ack)
	case "checkbalance":
		var ack bool
		err = client.Call("RPCCacheService.Clear",true,&ack)
		if err != nil{
			fmt.Println("RPC server call failed...ERROR:",err)
		}
		fmt.Println(ack)
	default:
		fmt.Errorf("unkown type: %s",t)
		os.Exit(1)
	}


}