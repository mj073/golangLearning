package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go routine(ch1, ch2)

	fmt.Println(<- ch1)
	fmt.Println(<- ch2)
	fmt.Println("exiting main")
}
func routine(ch1, ch2 chan int) {

	ch1 <- 1
	ch2 <- 2
}
