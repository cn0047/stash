package main

import (
	"fmt"
)

func main() {
	threshold, cost, weight := 0, []int{}, []int{}
	threshold, cost, weight = 8, []int{1, 2, 5, 6}, []int{2, 3, 4, 5}  // 8
	threshold, cost, weight = 30, []int{70, 20, 50}, []int{11, 12, 13} // 120

	res := tabulationKnapSack01(threshold, weight, cost)
	fmt.Printf("max knapSack: %v \n", res)
}

// tabulationKnapSack01 represents 0-1 Knapsack problem algorithm based on tabulation/table approach.
// Plan: define 2d-table with values calculated by formula, pick max value from table.
func tabulationKnapSack01(threshold int, weight []int, cost []int) int {
	n := len(cost)

	// k holds table data with calculated values.
	k := make([][]int, n+1)
	for i := range k {
		k[i] = make([]int, threshold+1)
	}
	for i := 0; i <= n; i++ {
		for w := 0; w <= threshold; w++ {
			if i == 0 || w == 0 {
				k[i][w] = 0
			} else if weight[i-1] <= w {
				k[i][w] = max(k[i-1][w], cost[i-1]+k[i-1][w-weight[i-1]]) // main formula.
			} else {
				k[i][w] = k[i-1][w]
			}
		}
	}

	return k[n][threshold]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
