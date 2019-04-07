package main

import (
	"encoding/json"
	"fmt"
	"github.com/golangLearning/bankApp/common"
	"io/ioutil"
	"net"
	"net/rpc"
	"os"
	"time"
)

type TransactionClient struct {
	Id int
	Server string
	Request common.TransactionRequest
	Response common.TransactionResponse
}
var (
	DummyRequests = createTransactionRequest()
	nRequests = len(DummyRequests)
	responseChan = make(chan *TransactionClient, nRequests)
	nResponse int
)
func createTransactionRequest() []common.TransactionRequest{
	tr := []common.TransactionRequest{}
	details := []common.CustomerDetails{}
	b, err := ioutil.ReadFile("/home/mahesh/Downloads/MOCK_DATA.json")
	if err != nil {
		fmt.Println("failed to read file...ERROR:",err)
		os.Exit(1)
	}
	err = json.Unmarshal(b,&details)
	if err != nil {
		fmt.Println("failed to unmarshal...ERROR:",err)
		os.Exit(1)
	}
	for _, d := range details {
		tr = append(tr, common.TransactionRequest{Type: common.CreateAccount, Details: d})
		/*tr[i].Type = common.CreateAccount
		tr[i].Details = d*/
	}
	return tr
}
func (c *TransactionClient) Do() {
	MAXRETRY := 5
	var conn net.Conn
	var err error
	for i:=0; i<MAXRETRY; i++ {
		conn, err = net.Dial("unix","@"+"bankServer")
		if err != nil{
			if i != MAXRETRY {
				time.Sleep(100 *time.Millisecond)
				continue
			}else {
				fmt.Println("failed to dial unix socket...ERROR:", err)
				os.Exit(1)
			}
		}else {
			break
		}
	}
	if conn != nil {
		defer conn.Close()
		client := rpc.NewClient(conn)
		defer client.Close()
		err = client.Call("BankTransaction.Transact",&c.Request,&c.Response)
		if err != nil{
			fmt.Println("client call ERROR:",err)
		}
		responseChan <- c
	}

}
func createTransactionClient(req common.TransactionRequest,id int) *TransactionClient{

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
	for id, req := range DummyRequests {
		go createTransactionClient(req,id).Do()
	}
	done := make(chan bool)
	go waitForAllResponses(done)
	fmt.Println("waiting for done..")
	<- done
}