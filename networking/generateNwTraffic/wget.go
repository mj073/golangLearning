package main

import (
	"time"
	"net/http"
	"fmt"
)

func main(){

	//sleeptime := 50
	//const cap = 20
	nRuns := 20000
	msg := make(chan string)
	closeSignal := make(chan bool)
	k := 0
	go func() {
		for {
			select {
			case x := <- msg :
				fmt.Println(x)
				k++
			default:
				if k == nRuns {
					closeSignal <- true
				}
			}
			time.Sleep(time.Millisecond * 20)
		}
	}()
	for i :=0; i < nRuns;i++{
		go func (j int){
			resp, err := http.Get("http://10.0.1.32/1K")
			if err == nil && resp.Body != nil{
				msg <- fmt.Sprint(">> closing request",j)
				resp.Body.Close()
				resp.Close = true
			}
		}(i)
	}
	<- closeSignal
	//for i := 1; i <= sleeptime; i++{
	//	fmt.Println("i=",i)
	//	time.Sleep(time.Second)
	//}

}
