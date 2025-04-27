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
	go processTask(taskChan)
	for task := 0; task < 5; task++ {
		fmt.Println("sending task",task)
		taskChan <- task
	}
	for task := 0; task < 5; task++ {
		fmt.Println("Main task:", task)
		time.Sleep(time.Millisecond * 100)
	}
	<-done
}

func processTask(task chan int) {
	for t := range task {
		executeTask(t)
	}
}
func executeTask(task int) {
	time.Sleep(5 * time.Second)
	fmt.Println("Executing task:", task)
	count++
	if count == 5 {
		done <- true
	}
}
