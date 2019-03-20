package main

import (
	"github.com/fluent/fluent-logger-golang/fluent"
	"fmt"
)

type Student struct {
	Name string
	Id_no int
	Age int
}

func main() {
	logger,err :=fluent.New(fluent.Config{FluentHost:"192.168.0.118", FluentPort:24224})
	fmt.Println(logger.BufferLimit)
	if err != nil{
		fmt.Println("error creating fluentd instance")
		return
	}

	defer logger.Close()
	tag := "fluentd-log-demo"
	data := Student{"John",1234,25}
	error := logger.Post(tag,data)
	if error != nil{
		panic(error)
	}
}