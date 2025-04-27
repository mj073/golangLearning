/*
Write a function, shortestPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB).
The function should return the length of the shortest path between A and B.
Consider the length as the number of edges in the path, not the number of nodes.
If there is no path between A and B, then return -1. You can assume that A and B exist as nodes in the graph.
*/

package main

import (
	"container/list"
	"fmt"
)

var edges = [][]string {
	{"w", "x"},
	{"x", "y"},
	{"z", "y"},
	{"z", "v"},
	{"w", "v"},
}

func main() {
	buildGraph()
	fmt.Println("shortest path(w,z): ", shortestPath("w", "z"))
}

var graph map[string][]string

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
var visitedNodes = make(map[string]struct{})
var q *list.List
func shortestPath(src, dst string) (distance int) {
	q = list.New()
	type node struct {
		n string
		distance int
	}
	q.PushBack(node{src, 0})
	loop:
		if current := q.Front(); current != nil {
			c := current.Value.(node)
			if _, ok := visitedNodes[c.n]; !ok {
				visitedNodes[c.n] = struct{}{}
			}
			if c.n == dst {
				return c.distance
			}
			for _, v := range graph[c.n] {
				if _, ok := visitedNodes[v]; !ok {
					visitedNodes[v] = struct{}{}
					q.PushBack(node{v, c.distance + 1})
				}
			}
			q.Remove(current)
			goto loop
		}
	return -1
}
