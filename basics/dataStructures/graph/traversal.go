package main

import (
	"container/list"
	"fmt"
	"time"
)

// Adjacency List
var graph = map[string][]string{
	"a": {"b","c"},
	"b": {"d"},
	"c": {"e"},
	"d": {"f"},
	"e": {},
	"f": {},
}

func main() {
	dft("a")		// Uses Stack
	dft2("a")	// Using Recursive
	fmt.Println()
	bft("a")		// Uses Queue
}

var stack []string
func dft(source string) {
	stack = append(stack, source)
	loop:
		if current, ok := pop(); ok {
			fmt.Print(current,",")
			for _, v := range graph[current] {
				stack = append(stack, v)
			}
			time.Sleep(time.Second)
			goto loop
		}
	fmt.Println()
}

func pop() (s string, ok bool) {
	if l := len(stack); l != 0 {
		s = stack[l-1]
		stack = stack[:l-1]
		//fmt.Println("stack after pop", stack)
		ok = true
	}
	return
}

func dft2(source string) {
	fmt.Print(source, ",")
	for _,v := range graph[source] {
		time.Sleep(time.Second)
		dft2(v)
	}
}

var q *list.List
func bft(source string) {
	q = list.New()
	q.PushBack(source)
	loop:
	if e := q.Front(); e != nil {
		fmt.Print(e.Value, ",")
		for _, v := range graph[e.Value.(string)] {
			q.PushBack(v)
		}
		time.Sleep(time.Second)
		q.Remove(e)
		goto loop
	}
	fmt.Println()
}
