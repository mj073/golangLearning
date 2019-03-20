package main

import (
	"fmt"
)
type slice struct{
	x int
}
type intSlice []slice
func main() {
	fmt.Println("Hello, playground")
	arr := &intSlice{{x:1},{x:3},{x:4},{x:5}}
	fmt.Println(*arr)
	fmt.Println((*arr)[0:])
	fmt.Println((*arr)[4:])
	fmt.Println((*arr)[:len(*arr)])
	fmt.Println((*arr)[:len(*arr)-3])
	arr.resize(15)
	fmt.Println("len(arr):",len(*arr))
	fmt.Println("cap(arr):",cap(*arr))
	fmt.Println("arr:",arr)
}

func (arr *intSlice) resize(n int){
	current_cap := cap(*arr)
	current_len := len(*arr)
	fmt.Println("current cap:",current_cap)
	fmt.Println("current len:",current_len)
	q := make(intSlice,n,current_cap*2+n)
	copy(q,*arr)
	fmt.Println(q)
	fmt.Println(*arr)
	*arr = q[:n]
}