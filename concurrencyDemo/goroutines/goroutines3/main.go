/*
Fixing sequence of goroutines using time.Sleep(). Not Recommended
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	for i:=0; i<5; i++ {
		go executeTask(i)
		//time.Sleep(time.Millisecond * 1)
	}
	time.Sleep(1 * time.Second)
}

func executeTask(task int) {
	fmt.Println("Executing task:",task)
}