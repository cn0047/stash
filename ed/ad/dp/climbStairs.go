// @see: https://leetcode.com/problems/climbing-stairs/
package main

import (
	"fmt"
)

func main() {
	r := 0
	// r = climbStairs(2) // expected 2
	// r = climbStairs(3) // expected 3
	// r = climbStairs(4) // expected 5
	r = climbStairs(5) // expected 8
	fmt.Printf("%v \n", r)
}

// climbStairs returns count of distinct ways to climb to the top.
func climbStairs(n int) int {
	// return climbStairs1(n)
	return climbStairs2(n)
}

func climbStairs2(n int) int {
	cache := make([]int, n+1)
	cache[0] = 1
	cache[1] = 1
	for i := 2; i < n+1; i++ {
		cache[i] = cache[i-2] + cache[i-1] // OPTIMIZATION: don't store whole array, just 2 elements
	}
	return cache[n]
}

// @see: https://monosnap.com/file/LcQPQGS9b5LrdvExSHRO1C0Y5DPrF5
func climbStairs1(n int) int {
	cache := make([]int, n+1)
	cache[n] = 1
	cache[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		cache[i] = cache[i+1] + cache[i+2] // OPTIMIZATION: don't store whole array, just 2 elements
	}

	return cache[0]
}
