package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	check("Mahesh")
	check(45)

}
func check(v interface{}){
	// value of v gets assigne to x
	// but we do switching over type of v
	switch x := v.(type){
	case string:
		fmt.Println(x,"is string..")
	case int:
		fmt.Println(x,"is integer..")
	}
}