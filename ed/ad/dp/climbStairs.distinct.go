package main

import (
	"fmt"
)

func main() {
	r := 0
	// r = climbStairsDistinctCount(2) // expected 2
	// r = climbStairsDistinctCount(3) // expected 3
	// r = climbStairsDistinctCount(4) // expected 5
	r = climbStairsDistinctCount(5) // expected 8
	fmt.Printf("%v \n", r)
}

// climbStairs returns count of distinct ways to climb to the top.
// @see: https://leetcode.com/problems/climbing-stairs
func climbStairsDistinctCount(n int) int {
	return climbStairsDistinctCount0(n)
	// return climbStairsDistinctCount1(n)
	// return climbStairsDistinctCount2(n)
}

func climbStairsDistinctCount2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-2] + dp[i-1] // OPTIMIZATION: don't store whole array, just 2 elements
	}
	return dp[n]
}

// @see: https://monosnap.com/file/LcQPQGS9b5LrdvExSHRO1C0Y5DPrF5
func climbStairsDistinctCount1(n int) int {
	dp := make([]int, n+1)
	dp[n] = 1
	dp[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		dp[i] = dp[i+1] + dp[i+2] // OPTIMIZATION: don't store whole array, just 2 elements
	}

	return dp[0]
}

// climbStairsDistinctCount0 - OPTIMIZED solution.
func climbStairsDistinctCount0(n int) int {
	v1 := 1
	v2 := 1
	for i := n - 2; i >= 0; i-- {
		v := v1 + v2
		v1, v2 = v, v1
	}

	return v1
}
