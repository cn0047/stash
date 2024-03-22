// @TODO: Return count of all possible coins variations.
// @TODO: Return count of unique coins variations.
// @TODO: Return min coins count with limited coins count.
package main

import (
	"fmt"
)

func main() {
	var r interface{}

	r = getMinCoinsCountChange([]int{1, 2, 5}, 11) // 3
	//r = getMinCoinsCountChange([]int{2}, 3)                    // -1
	//r = getMinCoinsCountChange([]int{1}, 0)                    // 0
	//r = getMinCoinsCountChange([]int{1, 2, 5, 10, 25, 50}, 73) // 5
	//r = getMinCoinsCountChange([]int{1, 2, 5, 10, 25, 50}, 87) // 4

	fmt.Printf("result: %+v \n", r)
}

// getMinCoinsCountChange returns fewest number of coins need to make up the amount.
// @see: https://leetcode.com/problems/coin-change
func getMinCoinsCountChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	// Initialize dp with value greater than amount.
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for c := 0; c < len(coins); c++ {
		coin := coins[c]
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

// min returns min value for a and b.
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
