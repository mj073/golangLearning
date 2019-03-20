package main

import (
	"log"
	"runtime/debug"
)

func main() {

	for i:=0;i<5 ;i++{
		log.Println("i=",i)
		panicGenerator()
	}
}
func panicGenerator(){
	lp := "panicGenerator"
	defer func() {
		if r := recover(); r != nil {
			log.Println(lp,"Recovered from crash..Details:", r,string(debug.Stack()))
			//panic(debug.Stack())
		}

	}()
	arr := []string{}
	lp = lp + "_1"
	log.Println(arr[0])
}