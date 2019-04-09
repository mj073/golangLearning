package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("in main")
	go panicRoutine()
	fmt.Println("sleeping for 1 sec")
	time.Sleep(1 * time.Second)
	fmt.Println("exiting main")
}
func panicRoutine(){
	panic("forceful panic")
	/*var a []byte
	fmt.Println(a[0])*/
}