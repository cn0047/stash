package main

import (
	"fmt"
)

func main() {
	r := 0

	// 1
	// r = uniquePaths(3, 2) // 3
	// r = uniquePaths(3, 7) // 28

	// 2
	// grid := [][]int{}
	// grid = [][]int{{0, 1}, {0, 0}}                  // 1
	// grid = [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}} // 2
	// r = uniquePathsWithObstacles(grid)

	// 3
	grid := [][]int{}
	grid = [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}} // 2
	grid = [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}  // 4
	r = uniquePathsIII(grid)

	fmt.Printf("%v\n", r)
}

// @see: https://leetcode.com/problems/unique-paths
func uniquePaths(m int, n int) int {
	cache := make([]int, n)
	cache[n-1] = 1

	for r := m - 1; r >= 0; r-- {
		for c := n - 1; c >= 0; c-- {
			if c+1 < n {
				cache[c] += cache[c+1]
			}
		}
	}

	return cache[0]
}

// @see: https://leetcode.com/problems/unique-paths-ii
func uniquePathsWithObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cache := make([]int, n)
	cache[n-1] = 1

	for r := m - 1; r >= 0; r-- {
		for c := n - 1; c >= 0; c-- {
			if grid[r][c] == 1 {
				cache[c] = 0
			} else if c+1 < n {
				cache[c] += cache[c+1]
			}
		}
	}

	return cache[0]
}

// @see: https://leetcode.com/problems/unique-paths-iii
func uniquePathsIII(grid [][]int) int {
	shouldBeVisitedCelsCount := 0
	startRow, startCol := 0, 0

	for r := 0; r < len(grid); r++ { // row
		for c := 0; c < len(grid[0]); c++ { // column
			if grid[r][c] == 0 {
				shouldBeVisitedCelsCount++
			} else if grid[r][c] == 1 {
				startRow = r
				startCol = c
			}
		}
	}

	return dfs(grid, startRow, startCol, shouldBeVisitedCelsCount)
}

// dfs represents DFS with additional counting logic.
// @category leetCodeSpecificSolution
func dfs(grid [][]int, x, y, shouldBeVisitedCelsCount int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == -1 {
		return 0
	}

	// Check whether end is reached.
	if grid[x][y] == 2 { // 2 - means end of path.
		if shouldBeVisitedCelsCount == -1 {
			return 1
		}
		return 0
	}

	grid[x][y] = -1 // mark as visited.
	shouldBeVisitedCelsCount--

	totalPaths := dfs(grid, x+1, y, shouldBeVisitedCelsCount) + // top
		dfs(grid, x, y+1, shouldBeVisitedCelsCount) + // right
		dfs(grid, x-1, y, shouldBeVisitedCelsCount) + // down
		dfs(grid, x, y-1, shouldBeVisitedCelsCount) // left

	grid[x][y] = 0
	shouldBeVisitedCelsCount++

	return totalPaths
}
