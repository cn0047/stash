// @see: https://leetcode.com/problems/parallel-courses-iii
package main

import (
	"fmt"
)

func main() {
	n, e, t := 0, [][]int{}, []int{}
	n, e, t = 3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 5}                               // 8
	n, e, t = 5, [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}, []int{1, 2, 3, 4, 5} // 12
	r := minimumTime(n, e, t)
	fmt.Printf("\nres: %v", r)
}

// minimumTime returns minimum time required to finish all parallel courses.
// Slice edges contains nodes with values starting from 1 (indexation starts from 1, not form 0).
// Slice time contains value in index 0 for node with value 1 (indexation starts from 0).
func minimumTime(n int, edges [][]int, time []int) int {
	adjacentNodes := make([][]int, n+1)
	inDegree := make([]int, n+1)
	for _, edge := range edges {
		a, b := edge[0], edge[1]                       // edge direction from left to right.
		adjacentNodes[a] = append(adjacentNodes[a], b) // adjacentNodes holds targets for node a from edge.
		inDegree[b]++
	}

	queue := make([]int, 0, n)
	maxTime := make([]int, n)

	for i := 1; i <= n; i++ { // from 1 because indexation starts from 1.
		if inDegree[i] == 0 {
			queue = append(queue, i)
			maxTime[i-1] = time[i-1] // -1 because indexation starts from 0.
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range adjacentNodes[node] {
			maxTime[neighbor-1] = max(maxTime[neighbor-1], maxTime[node-1]+time[neighbor-1])
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	r := 0
	for i := 0; i < n; i++ {
		r = max(r, maxTime[i])
	}

	return r
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
