/*
Write a function, connectedComponentsCount, that takes in the adjacency list of an undirected graph.
The function should return the number of connected components within the graph.
*/

package main

import "fmt"

var graph = map[int][]int {
	0: {8,1,5},
	1: {0},
	5: {0,8},
	8: {0,5},
	2: {3,4},
	3: {2,4},
	4: {3,2},
	//7: {},        // with this added, output should be 3
}

func main() {
	fmt.Println("number of connected components within the graph: ", connectedComponentsCount())
}

var visitedNodes = make(map[int]struct{})

func connectedComponentsCount() (count int) {
	for k, _ := range graph {
		if traverseDft(k) {
			count++
		}
	}
	return
}

func traverseDft(node int) bool {
	if _, ok := visitedNodes[node]; ok {
		return false
	}
	visitedNodes[node] = struct{}{}
	for _,v := range graph[node] {
		traverseDft(v)
	}
	return true
}
