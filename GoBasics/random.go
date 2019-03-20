package main

import (
	"fmt"
	"time"
)
var random = make(chan int)
var random2 = make(chan int)
func main() {
	go loop()
	fmt.Println("in main")
	time.Sleep(time.Second * 5)
	random<-1
	random2<-2
	time.Sleep(time.Minute *5)
}

func loop(){

	var k int;
	for{
		k++;
		select{
		case i:=<-random :
			fmt.Println("read on random",i)
		case j:= <- random2:
			fmt.Println("read on random2",j)
		default:
			fmt.Println("both are blocked")
		}
		fmt.Println("Sleeping")
		time.Sleep(5*time.Second)
	}
}