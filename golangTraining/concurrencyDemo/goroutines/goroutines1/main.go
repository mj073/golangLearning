package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// helloGoroutines()
	go helloGoroutines()
}
func helloGoroutines() {
	fmt.Println("Hello Goroutines")
}
