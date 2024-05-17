// @see: https://leetcode.com/problems/course-schedule-ii
package main

import (
	"fmt"
)

func main() {
	edges, n := [][]int{}, 0
	edges, n = [][]int{{0, 1}, {1, 2}, {3, 1}, {3, 2}}, 4                 // [2 1 3 0]
	edges, n = [][]int{{2, 3}, {3, 1}, {4, 0}, {4, 1}, {5, 0}, {5, 2}}, 6 // [1 3 2 0 5 4]
	edges, n = [][]int{{1, 0}, {0, 1}}, 2                                 // []
	edges, n = [][]int{{1, 0}}, 1                                         // [0 1]
	edges, n = [][]int{{0, 1}}, 1                                         // [0]
	edges, n = [][]int{}, 0                                               // []
	edges, n = [][]int{}, 2                                               // [1 0]
	edges, n = [][]int{}, 3                                               // [2 1 0]
	edges, n = [][]int{{1, 0}}, 3                                         // [2 0 1]
	r := courseSchedule(edges, n)
	fmt.Printf("res: %v \n", r)
}

func courseSchedule(edges [][]int, n int) []int {
	edges = swapElementsInSubSlices(edges)
	n = adjustN(edges, n)

	hasCycle := kahn(edges)
	if hasCycle {
		return nil
	}

	return topologicalSort(groupEdgesDirectionFromLeftToRight(edges), n)
}

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
func topologicalSort(edges map[int][]int, n int) []int {
	stack := []int{}
	visited := make(map[int]struct{}, n)
	res := []int{}

	for i := 0; i < n; i++ {
		_, ok := visited[i]
		if !ok {
			topologicalSubSort(i, edges, visited, &stack)
		}
	}

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, v)
	}

	return res
}

func topologicalSubSort(v int, edges map[int][]int, visited map[int]struct{}, stack *[]int) {
	visited[v] = struct{}{}

	for _, i := range edges[v] {
		_, ok := visited[i]
		if !ok {
			topologicalSubSort(i, edges, visited, stack)
		}
	}

	*stack = append(*stack, v)
}

// ============================================================================

// kahn represents Kahn's algorithm,
// and returns true in case when graph represented by edges has cycle.
func kahn(edges [][]int) (hasCycle bool) {
	edgesCount := len(edges)
	vertices := getVertices(edges)
	verticesCount := len(vertices)

	// inDegree holds map where key is vertice and value is verrice's incoming degree.
	inDegree := make(map[int]int, verticesCount)
	for i := 0; i < edgesCount; i++ {
		vertice := edges[i][1] // edge direction from left to right.
		_, ok := inDegree[vertice]
		if !ok {
			inDegree[vertice] = 0
		}
		inDegree[vertice]++
	}

	// Push into queue vertices without inDegree.
	queue := make([]int, 0, verticesCount)
	for v, _ := range vertices {
		val, ok := inDegree[v]
		if !ok || val == 0 {
			queue = append(queue, v)
		}
	}

	groupedEdges := groupEdgesDirectionFromLeftToRight(edges)
	checkedVertices := make([]int, 0, verticesCount)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // dequeue.
		checkedVertices = append(checkedVertices, node)

		adjacentNodes := groupedEdges[node]
		for _, adjacent := range adjacentNodes {
			_, ok := inDegree[adjacent]
			if ok {
				inDegree[adjacent]--
				if inDegree[adjacent] == 0 {
					queue = append(queue, adjacent)
				}
			}
		}
	}

	if verticesCount != len(checkedVertices) {
		hasCycle = true
	}

	return hasCycle
}

// groupEdgesDirectionFromLeftToRight converts [][]int{{0, 1}, {1, 2}, {3, 1}, {3, 2}} into map[0:[1] 1:[2] 3:[1 2]]
// where key is vertice 1 (1st value from edge pair),
// and value is all vertices 2 (2nd value from edge pair) of all edges for vertice 1.
func groupEdgesDirectionFromLeftToRight(edges [][]int) map[int][]int {
	res := make(map[int][]int, len(edges))
	for _, edge := range edges {
		_, ok := res[edge[0]]
		if !ok {
			res[edge[0]] = make([]int, 0, 0)
		}
		res[edge[0]] = append(res[edge[0]], edge[1])
	}

	return res
}

func getVertices(edges [][]int) map[int]struct{} {
	res := make(map[int]struct{}, len(edges))
	for _, edge := range edges {
		res[edge[0]] = struct{}{}
		res[edge[1]] = struct{}{}
	}

	return res
}

func swapElementsInSubSlices(edges [][]int) [][]int {
	res := make([][]int, len(edges))
	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		res[i] = []int{edge[1], edge[0]}
	}

	return res
}
