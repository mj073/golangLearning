package main

import (
	"fmt"
	"github.com/golangLearning/golangTraining/bankApp/common"
	"math/rand"
	"net"
	"net/rpc"
	"os"
)

type TransactionClient struct {
	Id int
	Server string
	Request common.TransactionRequest
	Response common.TransactionResponse
}
var DummyRequests = []common.TransactionRequest {
	common.TransactionRequest{
		Type: common.CreateAccount,
		Details: common.CustomerDetails{
			Name: "Mahesh Jadhav",
			Age: 27,
			Pan: "wjnjwbf",
		},
	},
	common.TransactionRequest{
		Type: common.CreateAccount,
		Details: common.CustomerDetails{
			Name: "Rahul",
			Age: 27,
			Pan: "gewrec",
		},
	},
	common.TransactionRequest{
		Type: common.CreateAccount,
		Details: common.CustomerDetails{
			Name: "Sachin",
			Age: 134,
			Pan: "wegehe",
		},
	},
	common.TransactionRequest{
		Type: common.CheckBalance,
		Details: common.CheckBalanceRequest{
			CustomerId: 1,
		},
	},
	common.TransactionRequest{
		Type: common.CheckBalance,
		Details: common.CheckBalanceRequest{
			CustomerId: 2,
		},
	},
	common.TransactionRequest{
		Type: common.Deposit,
		Details: common.DepositWithdrawDetails{
			CustomerId: 1,
			Amount: 10,
		},
	},
}
var (
	nRequests = len(DummyRequests)
	responseChan = make(chan *TransactionClient, nRequests)
	nResponse int
)
func (c *TransactionClient) Do() {
	conn, err := net.Dial("unix","@"+"bankServer")
	if err != nil{
		fmt.Println("failed to dial unix socket...ERROR:",err)
		os.Exit(1)
	}
	defer conn.Close()
	client := rpc.NewClient(conn)
	defer client.Close()

	err = client.Call("BankTransaction.Transact",&c.Request,&c.Response)
	if err != nil{
		fmt.Println("client call ERROR:",err)
	}
	responseChan <- c
}
func createTransactionClient(req common.TransactionRequest) *TransactionClient{
	id := rand.Intn(100)
	fmt.Println("creating client:",id,req.Type,req.Details)
	return &TransactionClient{
		Id: id,
		Request: req,
	}
}
func waitForAllResponses(done chan bool) {
	for {
		select {
		case r := <- responseChan:
			nResponse++
			fmt.Printf("response from client %d:\n%v\n",r.Id,r.Response)
			if nRequests == nResponse {
				done <- true
			}
		}
	}
}
func main(){
	common.RegisterAllGob()
	for _, req := range DummyRequests {
		go createTransactionClient(req).Do()
	}
	done := make(chan bool)
	go waitForAllResponses(done)
	<- done
}