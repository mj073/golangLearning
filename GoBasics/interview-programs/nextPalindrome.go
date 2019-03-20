package main

import (
	"fmt"
)

func main(){
	var num string
	for {
		fmt.Print("enter any number: ")
		fmt.Scanf("%v", &num)
		fmt.Println("num=", num)
		NextPalindrome(num)
	}
}

func NextPalindrome(num string){
	b := []byte(num)
	len := len(b)
	fmt.Println("len:",len)
	var prevIndex,nextIndex int
	var odd bool
	new := make([]byte,len)

	for i := 0;i < len / 2 - 1;i++{
		new[i] = b[i]
		new[len - 1 - i] = b[i]
	}
	fmt.Println("num: ",b)
	if len % 2 == 0{
		prevIndex,nextIndex = len / 2 - 1, len / 2
	}else {
		prevIndex,nextIndex = len / 2 - 1, len / 2 + 1
		odd = true
	}
	if b[prevIndex] < b[nextIndex]{
		if odd{
			new[nextIndex] = b[prevIndex]
			new[prevIndex + 1] = b[prevIndex + 1] + 1
			new[prevIndex] = b[prevIndex]
		}else{
			new[prevIndex] = b[nextIndex]
			new[nextIndex] = b[nextIndex]
		}
	}else if b[prevIndex] >= b[nextIndex]{
		new[nextIndex] = b[prevIndex]
		new[prevIndex] = b[prevIndex]
	}
	fmt.Println("new: ",fmt.Sprintf("%v",new))
	fmt.Printf("%s\n",new)
	//nextPal, _ := strconv.Atoi(string(new))
	//fmt.Println("next palindrome is:",nextPal)
}