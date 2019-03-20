package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string
	fmt.Println("Enter any long string:")
	fmt.Scanf("%s",&str)
	countermap:= make(map[string]int)
	for _,c := range str{
		countermap[string(c)] ++
	}
	fmt.Printf("%v\n",countermap)
	fmt.Println("count of y:",strings.Count(str,"y"))
}
