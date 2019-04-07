package main

import (
	"fmt"
)

var done = make(chan bool, 2)
func main(){
	var busyChan = make(chan int,5)
	go readFromChannel(busyChan)
	go writeToChannel(busyChan)

	<- done
	<- done
}
func readFromChannel(busyChan chan int) {
	for i:=0; i<100; i++{
		fmt.Println("reading from busyChan:",<-busyChan)
		//time.Sleep(time.Microsecond * 10)
	}

	done <- true
	fmt.Println("done reading")
}
func writeToChannel(busyChan chan int) {
	for i:=0; i<100; i++{
		fmt.Println("writing to busyChan:",i)
		busyChan <- i
		//time.Sleep(time.Microsecond * 10)
	}
	done <- true
	fmt.Println("done writing")
}
