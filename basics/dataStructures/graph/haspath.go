package main

import (
	"container/list"
	"fmt"
)

var graph = map[string][]string {
	"f": {"g", "i"},
	"g": {"h"},
	"h": {},
	"i": {"g", "k"},
	"j": {"i"},
	"k": {},
}

func main() {
	fmt.Printf("hasPathDft[f, k]: %v\n", hasPathDft("f", "k"))
	fmt.Printf("hasPathDft[j, f]: %v\n", hasPathDft("j", "f"))
	fmt.Printf("hasPathBft[f, k]: %v\n", hasPathBft("f", "k"))
	fmt.Printf("hasPathBft[j, f]: %v\n", hasPathBft("j", "f"))
}

func hasPathDft(src, dst string) bool {
	if src == dst {
		 return true
	}
	for _, v := range graph[src] {
		if hasPathDft(v, dst) {
			return true
		}
	}
	return false
}

var q *list.List
func hasPathBft(src, dst string) bool {
	if src == dst {
		return true
	}
	q = list.New()
	q.PushBack(src)
	loop:
		if current := q.Front(); current != nil {
			c := current.Value.(string)
			if c == dst {
				return true
			}
			for _,v := range graph[c] {
				q.PushBack(v)
			}
			q.Remove(current)
			goto loop
		}
	return false
}
