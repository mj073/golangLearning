package main

import (
	"fmt"
	"runtime/debug"
	//"errors"
	//"time"
	"time"
)

func main(){
	fmt.Println("in main()..")
	var chan1 = make(chan bool)
	i := 0
	//for {
	routine:
		go func() {
			x := 10
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(string(debug.Stack()))
				}
				fmt.Println("ending defer func()..")
				chan1 <- true
			}()
			for {
				func1()
				fmt.Println("i=", i)
				fmt.Println("x=",x)
				i++
				x++
			}
		}()
		//<- time.After(time.Second * 2)
		if ok := <- chan1;ok{
			time.Sleep(time.Millisecond * 50)
			goto routine
		}
	//}
	fmt.Println("ending main()...")
}

func func1(){
	fmt.Println("in func1()...")
	//panic(errors.New("forceful panic"))
}
