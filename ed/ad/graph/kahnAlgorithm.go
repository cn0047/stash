package main

import (
	"fmt"
)

func main() {
	edges := [][]int{}
	edges = [][]int{{2, 3}, {3, 1}, {4, 0}, {4, 1}, {5, 0}, {5, 2}} // true
	edges = [][]int{{0, 1}, {1, 2}, {3, 2}, {3, 4}}                 // true
	edges = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}                 // false
	edges = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {4, 2}, {4, 3}} // false
	edges = [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}                 // true
	edges = [][]int{{1, 0}}                                         // true
	edges = [][]int{{1, 0}, {0, 1}}                                 // false
	r := courseSchedule(edges)
	fmt.Printf("res: %v \n", r)
}

// @see: https://leetcode.com/problems/course-schedule
func courseSchedule(edges [][]int) bool {
	edges = swapElementsInSubSlices(edges)

	hasCycle := kahn(edges)

	return !hasCycle
}

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
