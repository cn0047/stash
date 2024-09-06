// @see: https://leetcode.com/problems/course-schedule-ii
// @see: https://gist.github.com/cn007b/44261e504172e91dd65b81dfbe2830bc
package main

import (
	"fmt"
)

func main() {
	edges, n := [][]int{}, 0
	n, edges = 4, [][]int{{0, 1}, {1, 2}, {3, 1}, {3, 2}}                 // [2 1 3 0]
	n, edges = 6, [][]int{{2, 3}, {3, 1}, {4, 0}, {4, 1}, {5, 0}, {5, 2}} // [1 3 2 0 5 4], [0 1 3 2 4 5]
	n, edges = 2, [][]int{{1, 0}, {0, 1}}                                 // []
	n, edges = 1, [][]int{{1, 0}}                                         // [0 1] //
	n, edges = 1, [][]int{{0, 1}}                                         // [0] //
	n, edges = 0, [][]int{}                                               // []
	n, edges = 2, [][]int{}                                               // [1 0], [0 1]
	n, edges = 3, [][]int{}                                               // [2 1 0], [0 1 2]
	n, edges = 3, [][]int{{1, 0}}                                         // [2 0 1], [0 1 2]
	n, edges = 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}                 // [0 1 2 3]
	r := courseSchedule(edges, n)
	fmt.Printf("res: %v \n", r)
}

func courseSchedule(edges [][]int, n int) []int {
	n = adjustN(edges, n)
	hasCycle := kahn(n, edges)
	if hasCycle {
		return nil
	}

	return topologicalSort(edges, n)
}

// adjustN adjusts n value, just for leetcode.
func adjustN(edges [][]int, n int) int {
	max := 0
	for i := 0; i < len(edges); i++ {
		e := edges[i]
		if e[0] > max {
			max = e[0]
		}
		if e[1] > max {
			max = e[1]
		}
	}

	if max > n {
		return max
	}

	return n
}

// topologicalSort represents function to perform topological sort.
// @uses: recursion topologicalSubSort.
func topologicalSort(edges [][]int, n int) []int {
	adjacentNodes := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]                       // direction b before a.
		adjacentNodes[a] = append(adjacentNodes[a], b)
	}

	stack := []int{}
	visited := make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		_, ok := visited[i]
		if !ok {
			topologicalSubSort(i, adjacentNodes, visited, &stack)
		}
	}

	return stack
}

func topologicalSubSort(node int, adjacentNodes [][]int, visited map[int]struct{}, stack *[]int) {
	visited[node] = struct{}{}

	for _, neighbor := range adjacentNodes[node] {
		_, ok := visited[neighbor]
		if !ok {
			topologicalSubSort(neighbor, adjacentNodes, visited, stack)
		}
	}

	*stack = append(*stack, node)
}

// @see: ed/ad/graph/kahnAlgorithm.go
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
