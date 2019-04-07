package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 2
	}()
	v, ok := <- ch
	if ok {
		fmt.Println("v:",v)
	}
}
