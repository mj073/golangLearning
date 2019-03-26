package common

import (
	"sync"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

type TransactionType uint8
const (
	CreateAccount TransactionType = iota
	Deposit
	Withdrawl
	CheckBalance
	CustomerDetails
)
type Balance int
func (b Balance) String() string{
	return fmt.Sprint(b)
}
type CustomerDB struct {
	CustomerInfo map[string]CustomerDetails
	CustomerKey map[uint]string
	BalanceByKey map[string]Balance
}
func (c *CustomerDB) isNewCustomer(d CustomerDetails) (string, bool) {
	key := generateKey([]byte(fmt.Sprint(d.Name," ",d.Age," ",d.Pan)))
	if _, ok := c.CustomerInfo[key]; !ok {
		return key,true
	}
	return "",false
}
func generateKey(b []byte) string{
	hasher := sha1.New()
	hasher.Write(b)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
type BankTransaction struct {
	CustomerDB *CustomerDB
	CustomerCount uint
	Mu *sync.RWMutex
}

type CustomerDetails struct {
	CustomerId uint
	Name	string
	Age 	uint8
	Pan	string
}
type DepositWithdrawDetails struct {
	CustomerId uint
	Amount	uint
}
type CustomerDetailsRequest struct {
	CustomerId uint
}
type TransactionRequest struct {
	Type TransactionType
	Details interface{}
}

type TransactionResponse struct {
	Response interface{}
	Ack bool
}
func (resp TransactionResponse) String() string{
	message := fmt.Sprint("Transaction")
	if resp.Ack{
		message += fmt.Sprintln(" Successful")
		switch r := resp.Response.(type) {
		case CreateAccountResponse:
			message += fmt.Sprint("Your customerId is:",r.customerId)
		case DepositResponse:
			message += fmt.Sprint("")
		case WithdrawResponse:
			message += fmt.Sprintf("Amount %d withdrawn",r.Amount)
		case CheckBalanceResponse:
			message += fmt.Sprintf("Your Current Balance: %s",r.Balance)
		case CustomerDetailsResponse:

		}
	}else {
		message += fmt.Sprintln(" Failed")
	}
	return
}
type CreateAccountResponse struct { customerId uint}
type DepositResponse struct {}
type WithdrawResponse struct { Amount uint }
type CheckBalanceResponse struct { Balance Balance }
type ErrorResponse struct {}
type CustomerDetailsResponse struct {

}

func NewBankTransaction() *BankTransaction{
	return &BankTransaction{
		CustomerDB: &CustomerDB{
			CustomerInfo: make(map[string]CustomerDetails),
			CustomerKey: make(map[uint]string),
			BalanceByKey: make(map[string]Balance),
		},
		Mu: &sync.RWMutex{},
	}
}

func (b *BankTransaction) Transact(req *TransactionRequest, resp *TransactionResponse) error{
	b.Mu.Lock()
	defer b.Mu.Unlock()

	switch req.Type {
	case CreateAccount:
		r := CreateAccountResponse{}
		c, ok := req.Details.(CustomerDetails)
		if !ok {
			resp.Ack = false
			resp.Response = ErrorResponse{}
			return fmt.Errorf("%s","failed to typecast request details")
		}

		fmt.Println("CreateAccount for:",)
		if key, ok := b.CustomerDB.isNewCustomer(c); ok {
			b.CustomerCount++
			c.CustomerId = b.CustomerCount
			b.CustomerDB.CustomerKey[c.CustomerId] = key
			b.CustomerDB.CustomerInfo[key] = c
			b.CustomerDB.BalanceByKey[key] = 0

			resp.Ack = true
			r.customerId = c.CustomerId
		}else {
			resp.Ack = false
		}
		resp.Response = r
	case CustomerDetails:
		r := CustomerDetailsResponse{}
		req.Details.(CustomerDetailsRequest)
	case Deposit:
		d := DepositResponse{}
		resp.Response = d
	case Withdrawl:
		w := WithdrawResponse{}
		resp.Response = w
	case CheckBalance:
		c := CheckBalanceResponse{}
		resp.Response = c
	default:

	}
	return
}

/*
func (r *BankTransaction) CreateAccount(c *CustomerDetails, ack *bool) error{

}

func (r *BankTransaction) CheckBalance(custId uint, balance *Balance) error{

}
func (r *BankTransaction) Deposit(details DepositWithdrawDetails, ack *bool) error{

}
func (r *BankTransaction) Withdraw(details DepositWithdrawDetails, ack *bool) error{

}*/
