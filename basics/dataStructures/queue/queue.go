package queue

import (
	"github.com/golangLearning/basics/dataStructures/linkedlist"
)

type Queue struct {
	linkedlist.List
	Tail *linkedlist.Node
}

func (q *Queue) PushBack(d interface{}) {
	n := &linkedlist.Node{
		Data: d,
	}
	if q.Head == nil {
		q.Head = n
	} else {
		q.Tail.Next = n
	}
	q.Tail = n
	q.Length++
}

func (q *Queue) Front() (h *linkedlist.Node) {
	if q.Head != nil {
		h = q.Head
		n := q.Head.Next
		q.Head.Next = nil
		q.Head = n
	}
	return
}