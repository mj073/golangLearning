package common

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/gob"
)

const (
	/*
	ProtoType = "unix"
	Address = "@bankServer"
	*/
	ProtoType = "tcp"
	Address = ":6666"
)
func generateKey(b []byte) string{
	hasher := sha1.New()
	hasher.Write(b)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func RegisterAllGob(){
	gob.Register(CreateAccountResponse{})
	gob.Register(ErrorResponse{})
	gob.Register(CheckBalanceResponse{})
	gob.Register(CustomerDetailsResponse{})
	gob.Register(DepositResponse{})
	gob.Register(WithdrawResponse{})
	gob.Register(DepositWithdrawDetails{})
	gob.Register(CustomerDetails{})
	gob.Register(CustomerDetailsRequest{})
	gob.Register(CheckBalanceRequest{})
}