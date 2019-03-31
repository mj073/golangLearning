/*
To let goroutine complete using time.Sleep(). Not Recommended
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	go helloGoroutines()
	time.Sleep(1 * time.Second)
}

func helloGoroutines() {
	fmt.Println("Hello Goroutines")
}
