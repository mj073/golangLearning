/*
Fixing sequence of goroutines using channels.
*/
package main

import (
	"fmt"
	"time"
)

var count = 0
var done = make(chan bool)

func main() {
	var taskChan = make(chan int, 5)
	fmt.Println("Hello World")
	startTime := time.Now()
	go processTask(taskChan)
	for task := 0; task < 5; task++ {
		taskChan <- task
	}
	<-done
	fmt.Println("time taken:", time.Since(startTime))
}

func processTask(task chan int) {
	for t := range task {
		executeTask(t)
	}
}
func executeTask(task int) {
	fmt.Println("Executing task:", task)
	count++
	if count == 5 {
		done <- true
	}
}
