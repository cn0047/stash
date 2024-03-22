package main

import (
	"fmt"
)

func main() {
	input := []int{}
	// input = []int{10, 15, 20} // 15
	// input = []int{10, 15, 20, 25} // 30
	// input = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1} // 6
	input = []int{1, 2, 3, 4, 100, 1, 2, 3, 100} // 10

	// r1 := climbStairsMinCost(input)
	// fmt.Printf("Result: %v \n", r1)

	r2, p := climbStairsMinCostPath(input)
	fmt.Printf("Result: %v, Path: %v \n", r2, p)
}

// climbStairsMinCost returns min cost to climb to the top.
// @see: https://leetcode.com/problems/min-cost-climbing-stairs
// @see: https://monosnap.com/file/naSjIx1Rg3blc9onRoPINWiCIsFvEG
func climbStairsMinCost(cost []int) int {
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}

	return min(cost[0], cost[1])
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func climbStairsMinCostPath(cost []int) (minCost int, path []int) {
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}

	minCost = min(cost[0], cost[1])

	for i := len(cost)-1; i>0; i-- {
		if cost[i-1] <= cost[i] {
			path = append([]int{i-1}, path...)
			i--
		} else {
			path = append([]int{i}, path...)
		}
	}

	return minCost, path
}
