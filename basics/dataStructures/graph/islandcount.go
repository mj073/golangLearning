/*
Write a function, islandCount, that takes in a grid containing Ws and Ls.
W represents water and L represents land.
The function should return the number of islands on the grid.
An island is a vertically or horizontally connected region of land.
*/
package main

import "fmt"

var grids = [][][]rune {
	{
		{'W', 'L', 'W', 'W', 'W'},
		{'W', 'L', 'W', 'W', 'W'},
		{'W', 'W', 'W', 'L', 'W'},
		{'W', 'W', 'L', 'L', 'W'},
		{'L', 'W', 'W', 'L', 'L'},
		{'L', 'L', 'W', 'W', 'W'},
	},
	{
		{'L', 'W', 'W', 'L', 'W'},
		{'L', 'W', 'W', 'L', 'L'},
		{'W', 'L', 'W', 'L', 'W'},
		{'W', 'W', 'W', 'W', 'W'},
		{'W', 'W', 'L', 'L', 'L'},
	},
	{
		{'L', 'L', 'L'},
		{'L', 'L', 'L'},
		{'L', 'L', 'L'},
	},
	{
		{'W', 'W'},
		{'W', 'W'},
		{'W', 'W'},
	},
}

var visitedNodes = make(map[string]struct{})

func main() {
	for i, grid := range grids {
		fmt.Printf("number of islands on the grid[%d]: %d\n", i+1, isLandCount(grid))
	}
}

func isLandCount(grid [][]rune) (count int) {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if explore(grid, r, c) {
				count++
			}
		}
	}
	return
}

func explore(grid [][]rune, r, c int) bool {
	rBound := r >=0 && r < len(grid)
	cBound := c >=0 && c < len(grid[0])
	if !rBound || !cBound {
		return false
	}
	if grid[r][c] == 'W' {
		return false
	}
	key := fmt.Sprintf("%b,%b", r,c)
	if _, ok := visitedNodes[key]; ok {
		return false
	}
	visitedNodes[key] = struct{}{}

	explore(grid, r-1, c)
	explore(grid, r+1, c)
	explore(grid, r, c-1)
	explore(grid, r, c+1)
	return true
}