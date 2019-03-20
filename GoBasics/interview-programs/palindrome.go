package main

import (
	"fmt"
)

func main() {
	var num string
	for {
		fmt.Print("enter any string: ")
		fmt.Scanf("%v", &num)
		fmt.Println("num=", num)
		TestPalindrome(num)
	}
}

func TestPalindrome(num string){
	b := []byte(num)
	fmt.Println("num in bytes:",b)
	len := len(b)
	for i:=0;i< len -1 ;i++{
		if(b[i] == b[len-1-i]){
			continue
		}else{
			fmt.Println("is not a palindrome")
			return
		}
	}
	fmt.Println("is a palindrome")
}