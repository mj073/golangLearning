package main

import "fmt"

func main(){
	fmt.Println("in main...")
	defer fmt.Println("defer statement 1")
	defer fmt.Println("defer statement 2")
	defer fmt.Println("defer statement 3")
	defer fmt.Println("defer statement 4")
}