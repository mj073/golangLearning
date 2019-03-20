package main

import "fmt"

func main(){
	fmt.Println("Hello main")
	defer fmt.Println("defered statement1...")
	defer func(){
		fmt.Println("defered statement2...")
	}()
	defer func(){
		fmt.Println("defered statement3...")
		panic("forceful panic")
	}()
}
