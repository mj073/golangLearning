package main

import (
	"fmt"
	//"time"
)

func main() {
	i := 0
	ch := make(chan int)
	go func() {
		i += 2
		ch <- i
	}()
	go func() {
		//time.Sleep(1 * time.Second)
		i++
		ch <- i
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}