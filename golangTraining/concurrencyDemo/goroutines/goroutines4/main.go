/*
Listing Goroutines
 */
package main

import (
	"fmt"
	"net/http"
	"time"
	_ "net/http/pprof"
)

func main() {
	fmt.Println("Hello World")
	for task:=0; task<5; task++ {
		go executeTask(task)
		time.Sleep(time.Millisecond * 1)
	}
	fmt.Println(http.ListenAndServe("localhost:9090",nil))
}

func executeTask(task int) {
	fmt.Println("Executing task:",task)
	time.Sleep(time.Second * 20)
	fmt.Printf("Task %d Done\n",task)
}
