/*
Fixing sequence of goroutines using channels.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var taskChan = make(chan int)
	fmt.Println("Hello World")
	startTime := time.Now()
	go processTask(taskChan)
	for task := 0; task < 5; task++ {
		taskChan <- task
		//executeTask(task)
	}
	fmt.Println("time taken:",time.Since(startTime))

}

func processTask(task chan int) {
	for t := range task{
		executeTask(t)
	}
}

func executeTask(task int) {
	fmt.Println("Executing task:",task)
}