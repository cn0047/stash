package main

import (
	"fmt"
	"math"
)

func main() {
	var r interface{}

	r = getChange(18, 17.83) // change: 0.17 r: [2 1 1 0 0 0]
	r = getChange(5, 0.61) // change: 4.39 r: [4 0 1 1 0 4]
	r = getChange(5, 0.51) // change: 4.49 r: [4 0 2 1 0 4]

	fmt.Printf("result: %+v \n", r)
}

// getChange returns coins array with change calculated by formula: money-price.
// Given coins array:   1c,   5c,  10c,  25c, 50c, 1$.
// Numbers for coins: 0.01, 0.05, 0.10, 0.25, 0.5,  1.
func getChange(money float32, price float32) []int32 {
	res := make([]int32, 6)

	diff := money - price
	dollars, coins := math.Modf(float64(diff))
	c := round(coins)

	res[5] = int32(dollars)

	if c-0.5 >= 0 {
		res[4] = 1
		c = round(c - 0.5)
	}

	if c-0.25 >= 0 {
		res[3] = 1
		c = round(c - 0.25)
	}

	if c-0.2 >= 0 {
		res[2] = 2 // 2 coins 10c
		c = round(c - 0.2)
	}

	if c-0.1 >= 0 {
		res[2] = 1
		c = round(c - 0.1)
	}

	if c-0.05 >= 0 {
		res[1] = 1
		c = round(c - 0.05)
	}

	res[0] = int32(c / 0.01)

	return res
}

// round performs math round.
func round(x float64) float64 {
	return math.Round(x*100) / 100
}
