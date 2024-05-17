// @see: https://leetcode.com/problems/course-schedule
package main

import (
	"fmt"
)

func main() {
	n, edges := 0, [][]int{}
	n, edges = 6, [][]int{{2, 3}, {3, 1}, {4, 0}, {4, 1}, {5, 0}, {5, 2}} // true
	n, edges = 5, [][]int{{0, 1}, {1, 2}, {3, 2}, {3, 4}}                 // true
	n, edges = 5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}                 // false
	n, edges = 5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {4, 2}, {4, 3}} // false
	n, edges = 5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}                 // true
	n, edges = 2, [][]int{{1, 0}}                                         // true
	n, edges = 2, [][]int{{1, 0}, {0, 1}}                                 // false
	r := courseSchedule(n, edges)
	fmt.Printf("res: %v \n", r)
}

func courseSchedule(n int, edges [][]int) bool {
	hasCycle := kahn(n, edges)

	return !hasCycle
}

func kahn(n int, edges [][]int) (hasCycle bool) {
	adjacentNodes := make([][]int, n)
	inDegree := make([]int, n)

	for _, edge := range edges {
		a, b := edge[0], edge[1] // direction right to left
		adjacentNodes[b] = append(adjacentNodes[b], a) // adjacentNodes holds targets for node b from edge.
		inDegree[a]++
	}

	queue := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	checkedVertices := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		checkedVertices++
		for _, neighbor := range adjacentNodes[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if checkedVertices != n {
		return true
	}

	return false
}
