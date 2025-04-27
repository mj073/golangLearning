package main

import (
	"container/list"
	"fmt"
)

// Given Edges
var edges = [][]string {
	{"i", "j"},
	{"k", "i"},
	{"m", "k"},
	{"k", "l"},
	{"o", "n"},
}

// Convert Edges into Adjacency List / Graph
var graph map[string][]string

var visitedNodes = make(map[string]struct{})

func main() {
	buildGraph()
	fmt.Printf("undirectedPathDft[i, l]: %v\n", undirectedPathDft("i", "l"))
	visitedNodes = map[string]struct{}{}
	fmt.Printf("undirectedPathDft[i, n]: %v\n", undirectedPathDft("i", "n"))
	visitedNodes = map[string]struct{}{}
	fmt.Printf("undirectedPathBft[i, l]: %v\n", undirectedPathBft("i", "l"))
	visitedNodes = map[string]struct{}{}
	fmt.Printf("undirectedPathBft[i, n]: %v\n", undirectedPathBft("i", "n"))
}

func buildGraph() {
	for _, v := range edges {
		if graph == nil {
			graph = make(map[string][]string)
		}
		if _, ok := graph[v[0]]; ok {
			graph[v[0]] = append(graph[v[0]], v[1])
		} else {
			graph[v[0]] = []string{v[1]}
		}
		if _, ok := graph[v[1]]; ok {
			graph[v[1]] = append(graph[v[1]], v[0])
		} else {
			graph[v[1]] = []string{v[0]}
		}
	}
	fmt.Println(graph)
}

func undirectedPathDft(src, dst string) bool {
	if src == dst {
		return true
	}
	if _, ok := visitedNodes[src]; !ok {
		visitedNodes[src] = struct{}{}
	} else {
		return false
	}
	for _, v := range graph[src] {
		if undirectedPathDft(v, dst) {
			return true
		}
	}
	return false
}

var q *list.List
func undirectedPathBft(src, dst string) bool {
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
			if _, ok := visitedNodes[c]; !ok {
				visitedNodes[c] = struct{}{}
			}
			for _, v := range graph[c] {
				if _, ok := visitedNodes[v]; !ok {
					q.PushBack(v)
				}
			}
			q.Remove(current)
			goto loop
		}
	return false
}

