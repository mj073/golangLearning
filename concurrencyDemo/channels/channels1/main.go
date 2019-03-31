package main

import "fmt"

func main() {
	var done = make(chan bool)
	fmt.Println("Hello World")
	go helloGoroutines(done)
	<- done
}
func helloGoroutines(done chan bool) {
	fmt.Println("Hello Goroutines")
	done <- true
}
