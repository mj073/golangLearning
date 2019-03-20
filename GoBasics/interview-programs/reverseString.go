package main

import "fmt"

func main(){
	var str,reverseStr string
	fmt.Print("enter any string: ")
	fmt.Scanf("%s",&str)
	for _,c := range str{
		reverseStr = string(c) + reverseStr
	}
	fmt.Printf("reverse string: %s",reverseStr)
}
