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
			//fmt.Printf("addr: %v val: %v ",&ch, ch)
			/*_, ok := <-ch
			if !ok {
				fmt.Println("receive on closed chann")
				break
			}*/
			fmt.Println("received")
		}
	}()
	ch <- 5
	close(ch)
	time.Sleep(1 * time.Second)
}
