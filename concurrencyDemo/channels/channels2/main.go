package main

import (
	"fmt"
)

func main() {
	var taskChan = make(chan int)
	fmt.Println("Hello World")
	go processTask(taskChan)
	for task := 0; task < 5; task++ {
		taskChan <- task
		//executeTask(task)
	}
}

func processTask(task chan int) {
	for t := range task{
		executeTask(t)
	}
}

func executeTask(task int) {
	fmt.Println("Executing task:",task)
}