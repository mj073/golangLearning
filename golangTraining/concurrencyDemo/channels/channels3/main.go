package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for {
			// resource leak
			fmt.Println(<- ch)
			fmt.Println("task")
		}
	}()
	ch <- 5
	close(ch)
	time.Sleep(1 * time.Second)
}
