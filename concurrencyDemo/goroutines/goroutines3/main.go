/*
Fixing sequence of goroutines using time.Sleep(). Not Recommended
 */
package main

import (
	"fmt"
	"time"
)
type Task struct {
	taskId int
	data int
}
func main() {
	fmt.Println("Hello World")
	data := 10
	task1 := Task{data: data}
	for i:=0; i<5; i++ {
		task1.taskId = i
		go executeAddTask(task1)
		//time.Sleep(time.Millisecond * 1)
	}
	time.Sleep(1 * time.Second)
}

func executeAddTask(task Task) {
	fmt.Printf("Executing add task %d: %v",task.taskId,task.data+task.taskId)
}
