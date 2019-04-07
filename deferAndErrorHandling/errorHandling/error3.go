package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("in main")
	err := throwCustomError()
	fmt.Printf("%T %v\n",err,err)

	err = throwNewError()
	fmt.Printf("%T %v\n",err,err)
}
func throwCustomError() error{
	return fmt.Errorf("custom error")
}
func throwNewError() error{
	return errors.New("new error")
}