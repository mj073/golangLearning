package main

import "fmt"

func main() {
	fmt.Println("in main")
	defer fmt.Println("defer called")
	fmt.Println("exiting main")
}
