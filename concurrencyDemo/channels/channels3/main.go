package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for {
			<- ch
			fmt.Println("received")
		}
	}()
	ch <- 5
	close(ch)
	time.Sleep(1 * time.Second)
}
