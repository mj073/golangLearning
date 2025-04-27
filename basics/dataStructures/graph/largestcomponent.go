/*
Write a function, largestComponent, that takes in the adjacency list of an undirected graph.
The function should return the size of the largest connected component in the graph.
*/

package main

import "fmt"

var graph = map[int][]int{
	0: {8, 1, 5},
	1: {0},
	5: {0, 8},
	8: {0, 5},
	2: {3, 4},
	3: {2, 4},
	4: {3, 2},
}

var visitedNodes = make(map[int]struct{})

func main() {
	fmt.Println("size of the largest connected component in the graph: ", largestComponent())
}

func largestComponent() (max int) {
	for k, _ := range graph {
		if s := traverseDft(k); s > max {
			max = s
		}
	}
	return
}

func traverseDft(node int) (size int) {
	if _, ok := visitedNodes[node]; ok {
		return
	} else {
		visitedNodes[node] = struct{}{}
	}
	size = 1
	for _, v := range graph[node] {
		size += traverseDft(v)
	}
	return
}