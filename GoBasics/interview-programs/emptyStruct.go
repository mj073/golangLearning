package main

import "reflect"

func worker(ch chan struct{}) {
	// Receive a message from the main program.
	<-ch
	println("roger from worker")

	// Send a message to the main program.
	close(ch)
}
type hello struct{
	x []int
	y []int
}
func main() {
	ch := make(chan struct{})
	go worker(ch)

	// Send a message to a worker.
	ch <- struct{}{}

	// Receive a message from the worker.
	<-ch
	println("roger from main")
	// Output:
	// roger
	// roger


	X:= hello{x: []int{1,2},y: []int{3,4}}
	Y:= hello{x: []int{1,2},y: []int{3,4}}

	//println(X == Y) //compilation error
	println(reflect.DeepEqual(X,Y))
}
