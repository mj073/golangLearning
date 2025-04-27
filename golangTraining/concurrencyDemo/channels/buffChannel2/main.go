package main

import (
	"fmt"
	"runtime"
)

var done = make(chan bool, 2)
var num = 100
func main(){
	/*runtime.GOMAXPROCS(1)
	var busyChan = make(chan int,num)*/
	var busyChan = make(chan int,5)
	go readFromChannel(busyChan)
	go writeToChannel(busyChan)

	<- done
	<- done
}
func readFromChannel(busyChan chan int) {
	for i:=0; i<num; i++{
		fmt.Println("reading from busyChan:",<-busyChan)
	}

	done <- true
	fmt.Println("done reading")
}
func writeToChannel(busyChan chan int) {
	for i:=0; i<num; i++{
		busyChan <- i
		fmt.Println("wrote,",i,"to busyChan:")
	}
	done <- true
	fmt.Println("done writing")
}
