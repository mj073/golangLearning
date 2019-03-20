package main

import (
	"fmt"
	"time"
)

func main(){
	var CloseListener = make(map[string]chan bool)
	//CloseListener["127.0.0.1:6379"] = make(chan bool)
	go func(){
		for {
			select {
			case x := <-CloseListener["127.0.0.1:6379"]:
				fmt.Println("closing listener for addr:","127.0.0.1:6379")
				fmt.Println(x)
				break
			}
		}
	}()
	if _, ok := CloseListener["127.0.0.1:6379"]; !ok{
		CloseListener["127.0.0.1:6379"] = make(chan bool)
	}
	CloseListener["127.0.0.1:6379"] <- true
	time.Sleep(time.Second * 1)
}