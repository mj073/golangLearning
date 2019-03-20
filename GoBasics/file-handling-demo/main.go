package main

import (
	"os"
	"fmt"
)

func main(){
	_, err := os.Open("D:\\randomfile.txt")
	if err != nil {
		fmt.Println("error opening file ERROR: ",err)
	}
}
