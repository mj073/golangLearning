package main

import (
	"time"
	"fmt"
	"sync"
)
var mu sync.RWMutex
var cache = make(map[string]int)
func main() {
	go sendDataRoutine()
	for i:=0; i<1000; i++{
		switch i%2{
		case 0:
			sendData(i,"even")
		case 1:
			sendData(i,"odd")
		}
		time.Sleep(time.Millisecond * 50)
	}
	<- time.After(time.Minute * 10)
}

func sendData(data int, dataType string){
	mu.Lock()
	cache[dataType] = data
	mu.Unlock()
}

func sendDataRoutine() {
	for {
		if len(cache) != 0{
			var temp map[string]int
			temp = cache
			//copy(temp,cache)
			mu.Lock()
			cache = make(map[string]int)
			mu.Unlock()
			for k,v := range temp{
				fmt.Println("k:",k,"v:",v)
			}
		}
		time.Sleep(time.Second * 15)
	}
}
