package main

import (
	"fmt"
	//"time"
)
var random = make(chan bool)
func main() {
	defer doProcess()
	fmt.Println("in main")

	//time.Sleep(time.Minute * 10)

//	<-random
}

func doProcess(){
	fmt.Println("in doProcess()")
}

