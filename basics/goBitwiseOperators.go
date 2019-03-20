package main

import (
	"fmt"
)
func main() {
	fmt.Println("Hello, playground")
	x := 2
	y := 4
	fmt.Printf("x = %b, y = %b, ^y = %b\n", x, y, ^(y - 1))
	fmt.Printf("x & ^y = %b\n", x & ^y)
}