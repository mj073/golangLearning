package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("in main")
	defer fmt.Println("defer statement in main")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic in main")
		}
	}()
	go panicRoutine3()
	//panic("forceful panic")
	fmt.Println("sleeping for 1 sec")
	time.Sleep(1 * time.Second)
	fmt.Println("exiting main")
}
func panicRoutine3(){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic in panicRoutine3")
		}
	}()
	defer fmt.Println("defer statement in panicRoutine2 before panic()")
	var a []byte
	fmt.Println(a[0])
	defer fmt.Println("defer statement in panicRoutine2 after panic()")

}