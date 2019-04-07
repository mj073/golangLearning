package common

import "fmt"

type TransactionType uint8
const (
	CreateAccount TransactionType = iota
	Deposit
	Withdrawl
	CheckBalance
	CustomerInfo
)

type Balance int

func (b Balance) String() string{
	return fmt.Sprint(int(b))
}

type CustomerDB struct {
	CustomerInfo map[string]CustomerDetails
	CustomerKey map[uint]string
	BalanceByKey map[string]Balance
}
func (c *CustomerDB) isCustomerExist(customerId uint) (info CustomerDetails, ok bool) {
	if key, okk := c.getKey(customerId); okk {
		info, ok = c.CustomerInfo[key]
		if ok {
			return
		}
	}
	return
}
func (c *CustomerDB) getKey(customerId uint) (string, bool) {
	key, ok := c.CustomerKey[customerId]
	return key,ok
}
func (c *CustomerDB) isNewCustomer(d CustomerDetails) (string, bool) {
	key := generateKey([]byte(fmt.Sprint(d.Name," ",d.Age," ",d.Pan)))
	if _, ok := c.CustomerInfo[key]; !ok {
		return key,true
	}
	return "",false
}

type CustomerDetails struct {
	CustomerId uint	`json:"customerId"`
	Name	string	`json:"name"`
	Age 	uint	`json:"age"`
	Pan		string	`json:"pan"`
}
func (c CustomerDetails) String() string{
	return fmt.Sprintf("CustomerId:%v, Name:%s, Age:%d, Pan:%s\n",c.CustomerId,c.Name,c.Age,c.Pan)
}

type DepositWithdrawDetails struct {
	CustomerId uint
	Amount	int
}
type CustomerDetailsRequest struct {
	CustomerId uint
}
type CheckBalanceRequest struct {
	CustomerId uint
}
type CreateAccountResponse struct { CustomerId uint}
type DepositResponse struct {}
type WithdrawResponse struct { Amount int }
type CheckBalanceResponse struct { Balance Balance }
type ErrorResponse struct { Msg string}
type CustomerDetailsResponse struct { Details []CustomerDetails }
