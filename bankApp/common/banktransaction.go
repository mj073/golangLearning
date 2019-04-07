package common

import (
	"fmt"
	"sync"
)

type BankTransaction struct {
	CustomerDB *CustomerDB
	CustomerCount uint
	Mu *sync.RWMutex
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
	message := fmt.Sprintln("Transaction Successful")

	switch r := resp.Response.(type) {
	case CreateAccountResponse:
		message += fmt.Sprint("Your customerId is:",r.CustomerId)
	case DepositResponse:
		message += fmt.Sprint("")
	case WithdrawResponse:
		message += fmt.Sprintf("Amount %d withdrawn",r.Amount)
	case CheckBalanceResponse:
		message += fmt.Sprintf("Balance: %s",r.Balance)
	case CustomerDetailsResponse:
		message += fmt.Sprintf("Customer Details:\n")
		for _, c := range r.Details {
			message += fmt.Sprint(c)
		}
	case ErrorResponse:
		message = fmt.Sprintln("Transaction Failed")
		message += "ERROR:"+r.Msg
	}
	return message
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

func (b *BankTransaction) Transact(req *TransactionRequest, resp *TransactionResponse) (err error){
	b.Mu.Lock()
	defer b.Mu.Unlock()

	switch req.Type {
	case CreateAccount:
		r := CreateAccountResponse{}
		resp.Ack = true
		c, ok := req.Details.(CustomerDetails)
		if !ok {
			resp.Ack = false
			resp.Response = ErrorResponse{ Msg: fmt.Sprintf("%s","failed to typecast request details")}
			return
		}
		if key, ok := b.CustomerDB.isNewCustomer(c); ok {
			b.CustomerCount++
			c.CustomerId = b.CustomerCount
			b.CustomerDB.CustomerKey[c.CustomerId] = key
			b.CustomerDB.CustomerInfo[key] = c
			b.CustomerDB.BalanceByKey[key] = 0

			r.CustomerId = c.CustomerId
		}else {
			resp.Ack = false
			resp.Response = ErrorResponse{ Msg: fmt.Sprintf("%s","customer already exists") }
			return
		}
		resp.Response = r
	case CustomerInfo:
		r := CustomerDetailsResponse{}
		resp.Ack = true
		c, ok := req.Details.(CustomerDetailsRequest)
		if !ok {
			resp.Ack = false
			resp.Response = ErrorResponse{ Msg: fmt.Sprintf("%s","failed to typecast request details")}
			return
		}
		if c.CustomerId == 0{
			for _,v := range b.CustomerDB.CustomerInfo {
				r.Details = append(r.Details, v)
			}
		}else if info, ok := b.CustomerDB.isCustomerExist(c.CustomerId); ok{
			r.Details = append(r.Details,info)
		}else {
			resp.Ack = false
			resp.Response = ErrorResponse{ Msg: fmt.Sprintf("%s","customer does not exists") }
			return
		}
		resp.Response = r

	case Deposit:
		d := DepositResponse{}
		resp.Ack = true
		r := req.Details.(DepositWithdrawDetails)
		if  key, ok := b.CustomerDB.getKey(r.CustomerId); ok{
			currBal, _ := b.CustomerDB.BalanceByKey[key]
			b.CustomerDB.BalanceByKey[key] = Balance(int(currBal) + r.Amount)
		}else {
			resp.Ack = false
			resp.Response = ErrorResponse{ Msg: fmt.Sprintf("%s","customer does not exists") }
			return
		}
		resp.Response = d
	case Withdrawl:
		w := WithdrawResponse{}
		resp.Ack = true
		r := req.Details.(DepositWithdrawDetails)
		if  key, ok := b.CustomerDB.getKey(r.CustomerId); ok{
			currBal, _ := b.CustomerDB.BalanceByKey[key]
			if int(currBal) < r.Amount {
				resp.Ack = false
				resp.Response = ErrorResponse{ Msg: fmt.Sprintf("%s","insuffient balance")}
				return
			}
			b.CustomerDB.BalanceByKey[key] = Balance(int(currBal) - r.Amount)
		}else {
			resp.Ack = false
			resp.Response = ErrorResponse{ Msg: fmt.Sprintf("%s","customer does not exists") }
			return
		}
		w.Amount = r.Amount
		resp.Response = w
	case CheckBalance:
		c := CheckBalanceResponse{}
		resp.Ack = true
		d := req.Details.(CheckBalanceRequest)
		if  key, ok := b.CustomerDB.getKey(d.CustomerId); ok{
			c.Balance, _ = b.CustomerDB.BalanceByKey[key]
		}else {
			resp.Ack = false
			resp.Response = ErrorResponse{ Msg: fmt.Sprintf("%s","customer does not exists") }
			return
		}
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
