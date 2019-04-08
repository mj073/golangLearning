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
	}

	done <- true
	fmt.Println("done reading")
}
func writeToChannel(busyChan chan int) {
	for i:=0; i<100; i++{
		busyChan <- i
		fmt.Println("wrote,",i,"to busyChan:")
	}
	done <- true
	fmt.Println("done writing")
}
