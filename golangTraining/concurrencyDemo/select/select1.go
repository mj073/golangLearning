package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int,1)
	c2 := make(chan int,1)
	c3 := make(chan int,1)
	c1 <- 1
	c2 <- 2
	//loop:
	//for {
		select {
		case v1 := <-c1:
			fmt.Printf("received %v from c1\n", v1)
		case v2 := <-c2:
			fmt.Printf("received %v from c2\n", v2)
		case c3 <- 23:
			fmt.Printf("sent %v to c3\n", 23)
		default:

			fmt.Printf("no one was ready to communicate\n")
			//break loop
		}
	//}
	fmt.Println(<-c3)
}
