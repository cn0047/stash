package main

import (
	"fmt"
)

func main() {
	edges := [][]int{}
	edges = [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}} // cost:26, path:[0 3 4 5]
	p, c := dijkstra(edges)
	fmt.Printf("\n===\ncost:%v, path:%v \n", c, p)
}

// dijkstra represents implementation for dijkstra algorith.
func dijkstra(edges [][]int) (path []int, spentCost int) {
	n := len(edges)

	adjacent := make([]map[int]int, len(edges))
	for i := 0; i < n; i++ {
		e := edges[i]
		a, b, c := e[0], e[1], e[2]
		if len(adjacent[a]) == 0 {
			adjacent[a] = make(map[int]int, n)
		}
		adjacent[a][b] = c
	}

	a := 0
	path = []int{a}

	for len(adjacent[a]) > 0 {
		minB, minCost := first(adjacent[a])
		for b, c := range adjacent[a] {
			if c < minCost {
				minCost = c
				minB = b
			}
		}
		spentCost += minCost
		a = minB
		path = append(path, minB)
	}

	return path, spentCost
}

func first(m map[int]int) (k int, v int) {
	for k, v := range m {
		return k, v
	}
	return 0, 0
}
