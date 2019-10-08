package main

import (
	"flag"
	"fmt"
	"github.com/golangLearning/golangTraining/bankApp/common"
	"net"
	"net/rpc"
	"os"
)
var (
	username = flag.String("user","admin","user=<username>")
	password = flag.String("pass","","pass=<password>")
	transactionType = flag.String("type","checkbalance","type=[createaccount|deposit|withdraw|checkbalance|customerdetails]")
	amount = flag.Uint("amount",0,"type=[deposit|withdraw] -amount=<amount>")
	customerId = flag.Uint("customerId",0,"customerId=<int>")

)
func main(){
	conn, err := net.Dial(common.ProtoType, common.Address)
	if err != nil{
		fmt.Println("failed to dial unix socket...ERROR:",err)
		os.Exit(1)
	}
	defer conn.Close()
	client := rpc.NewClient(conn)
	defer client.Close()
	flag.Parse()
	common.RegisterAllGob()

	request := common.TransactionRequest{}
	switch *transactionType {
	case "createaccount":
		details := common.CustomerDetails{}
		fmt.Println("Enter your personal info:")
		fmt.Print("Name: ")
		fmt.Scanf("%s",&details.Name)
		fmt.Print("Age: ")
		fmt.Scanf("%d\n",&details.Age)
		fmt.Print("Pan: ")
		fmt.Scanf("%s\n",&details.Pan)
		request.Details = details
		request.Type = common.CreateAccount
	case "customerdetails":
		if !(*username == "admin" && *password == "mahesh123"){
			fmt.Printf("%s\n","invalid username/password")
			os.Exit(1)
		}
		details := common.CustomerDetailsRequest{}
		if *customerId >= 0{
			details.CustomerId = *customerId
		}
		request.Details = details
		request.Type = common.CustomerInfo
	case "deposit":
		details := common.DepositWithdrawDetails{}
		if *customerId > 0{
			details.CustomerId = *customerId
			if *amount > 0 {
				details.Amount = int(*amount)
			}else {
				fmt.Printf("invalid amount\n")
				os.Exit(1)
			}
		}else {
			fmt.Printf("invalid customer id\n")
			os.Exit(1)
		}
		request.Details = details
		request.Type = common.Deposit

	case "withdraw":
		details := common.DepositWithdrawDetails{}
		if *customerId > 0{
			details.CustomerId = *customerId
			if *amount > 0 {
				details.Amount = int(*amount)
			}else {
				fmt.Printf("invalid amount\n")
				os.Exit(1)
			}
		}else {
			fmt.Printf("invalid customer id\n")
			os.Exit(1)
		}
		request.Details = details
		request.Type = common.Withdrawl
	case "checkbalance":
		details := common.CheckBalanceRequest{}
		if *customerId > 0{
			details.CustomerId = *customerId
		}else {
			fmt.Printf("invalid customer id\n")
		}
		request.Details = details
		request.Type = common.CheckBalance
	default:
		fmt.Printf("unkown type: %s",*transactionType)
		os.Exit(1)
	}
	response := common.TransactionResponse{}
	err = client.Call("BankTransaction.Transact",&request,&response)
	if err != nil{
		fmt.Println("ERROR:",err)
	}
	fmt.Println(response)
}