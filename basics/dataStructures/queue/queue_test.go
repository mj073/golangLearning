package queue

import (
	"fmt"
	"testing"
)

var q = &Queue{}
func TestQueue_PushBack(t *testing.T) {
	for i := range 11 {
		q.PushBack(i)
	}
	fmt.Println(q)
}

func TestQueue_PushBack1(t *testing.T) {
	for {
		var i int
		fmt.Print("enter: ")
		fmt.Scan(&i)
		q.PushBack(i)
		if i == 10 {
			break
		}
	}
	fmt.Println(q)
}

func TestQueue_Front(t *testing.T) {
	TestQueue_PushBack(t)
	fmt.Println(q.Front(), q)
}