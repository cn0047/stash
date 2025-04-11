package main

import (
	"fmt"
	"sort"
)

var (
	/*
			  1
			 / \
			/   2
		6    / \
		|   7   3
		\  / \  /
			5  | /
			 \ |/
			   4
	*/
	edges1 = [][]int{
		{1, 2, 28},
		{1, 6, 10},
		{2, 7, 14},
		{2, 3, 16},
		{3, 4, 12},
		{4, 7, 18},
		{4, 5, 22},
		{5, 6, 25},
		{5, 7, 24},
	} // MST: 1-6-5-4-3-2-7

	/*
		    3 4
		    |/
		    1
		   /| 5
		  / |/
		10--2--6--9
		    |\
		    | 7
		    8
		    |
		    |
		    11
	*/
	edges2 = [][]int{
		{1, 2, 0},
		{1, 3, 0},
		{1, 4, 0},
		{1, 10, 0},
		{2, 5, 0},
		{2, 6, 0},
		{2, 7, 0},
		{2, 8, 0},
		{2, 10, 0},
		{6, 9, 0},
		{8, 11, 0},
	}

	edges3 = [][]int{
		{1, 2, 2},
		{2, 3, 3},
		{3, 1, 10},
	} // MST: 1-2-3

	edges4 = [][]int{
		{0, 1, 2},
		{1, 2, 3},
		{0, 3, 6},
		{1, 3, 8},
		{1, 4, 5},
		{2, 4, 7},
	} // MST: 0-1-2-4-3

	edges5 = [][]int{
		{1, 2, 4}, {1, 3, 13}, {1, 4, 7}, {1, 5, 7},
		{2, 3, 9}, {2, 4, 3}, {2, 5, 7},
		{3, 4, 10}, {3, 5, 14},
		{4, 5, 4},
	} // MST: 2-4-1-5-3
	edges6 = [][]int{
		{1, 2, 23}, {1, 3, 36}, {1, 4, 64}, {1, 5, 36},
		{2, 3, 51}, {2, 4, 49}, {2, 5, 35},
		{3, 4, 28}, {3, 5, 16},
		{4, 5, 28},
	} // MST: 3-5-4-2-1, MinCost: 102
)

func main() {
	p, c := prim(edges1)
	fmt.Printf("\n===\nres:%v | %v\n", p, c)
}

// prim represents Prim's algorithm to find the Minimum Spanning Tree.
func prim(edges [][]int) (path []int, spentCost int) {
	n := len(edges)
	if n == 0 {
		return nil, 0
	}

	adjacent, currNode := prepareAdjacent(edges)
	visitedNodes := map[int]struct{}{currNode: {}}
	minSpanningTree := []int{currNode}
	workingSet := map[int]struct{}{}

	for {
		_, ok := adjacent[currNode]
		if ok {
			// Add currNode to workingSet.
			workingSet[currNode] = struct{}{}
		} else {
			// Delete currNode from workingSet.
			delete(workingSet, currNode)
		}

		minNode, minCost := -1, -1
		// Find actual non-visited min cost node.
		for node, _ := range workingSet {
			adjacentNodes := adjacent[node]
			for i := 0; i < len(adjacentNodes); i++ {
				_, visited := visitedNodes[adjacentNodes[i][0]]
				if !visited && (minCost == -1 || adjacentNodes[i][1] < minCost) {
					currNode = node
					minNode = adjacentNodes[i][0]
					minCost = adjacentNodes[i][1]
				}
			}
		}

		// Finish.
		if minCost == -1 {
			break
		}

		// Add currNode to minSpanningTree.
		visitedNodes[minNode] = struct{}{}
		minSpanningTree = append(minSpanningTree, minNode)
		spentCost += minCost
		delAdjacent(currNode, minNode, adjacent)

		// Next iteration.
		currNode = minNode
	}

	return minSpanningTree, spentCost
}

func prepareAdjacent(edges [][]int) (adjacent map[int][][]int, minNode int) {
	n := len(edges)

	// Prepare adjacent nodes.
	adjacent = make(map[int][][]int, n)
	for i := 0; i < n; i++ {
		e := edges[i]
		a, b, cost := e[0], e[1], e[2]

		// Edge a->b.
		_, ok := adjacent[a]
		if !ok {
			adjacent[a] = [][]int{}
		}
		adjacent[a] = append(adjacent[a], []int{b, cost})

		// Edge b->a.
		_, ok = adjacent[b]
		if !ok {
			adjacent[b] = [][]int{}
		}
		adjacent[b] = append(adjacent[b], []int{a, cost})
	}

	// Any first valid min cost node.
	first := getFirstKey(adjacent)
	minNode = adjacent[first][0][0]
	minCost := adjacent[first][0][1]

	// Sort adjacent nodes by cost and find actual min cost node.
	for _, edges := range adjacent {
		sort.Slice(edges, func(i, j int) bool {
			return edges[i][1] < edges[j][1]
		})
		// Find actual min cost node.
		if edges[0][1] < minCost {
			minNode = edges[0][0]
			minCost = edges[0][1]
		} else if edges[0][1] == minCost {
			// When cost same, then use node with smaller value.
			minNode = min(minNode, edges[0][0])
		}
	}

	return adjacent, minNode
}

func delAdjacent(a int, b int, adjacent map[int][][]int) {
	defer func() {
		if len(adjacent[a]) == 0 {
			delete(adjacent, a)
		}
		if len(adjacent[b]) == 0 {
			delete(adjacent, b)
		}
	}()
	// Edge a->b.
	_, ok := adjacent[a]
	if ok {
		for i := 0; i < len(adjacent[a]); i++ {
			if adjacent[a][i][0] == b {
				adjacent[a] = append(adjacent[a][:i], adjacent[a][i+1:]...)
				break
			}
		}
	}
	// Edge b->a.
	_, ok = adjacent[b]
	if ok {
		for i := 0; i < len(adjacent[b]); i++ {
			if adjacent[b][i][0] == a {
				adjacent[b] = append(adjacent[b][:i], adjacent[b][i+1:]...)
				break
			}
		}
	}
}

func getFirstKey(m map[int][][]int) int {
	for k, _ := range m {
		return k
	}
	panic("ERR_0001")
}
