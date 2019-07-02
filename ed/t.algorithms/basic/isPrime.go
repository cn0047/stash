package main

import (
	"math"
	"strconv"
)

// Prime number - a number that is divisible only by itself and 1.
func isPrime(s string) string {
	x, _ := strconv.ParseInt(s, 10, 64)
	if x == 1 {
		return "Not prime"
	}
	y := math.Sqrt(float64(x))
	for i := int64(2); i <= int64(y); i++ {
		if x%i == 0 {
			return "Not prime"
		}
	}
	return "Prime"
}
