// @TODO: Implement Held-Karp algorithm.
package main

import (
	"fmt"
)

func main() {
	distance := [][]int{}
	distance = [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	} // min=80
	distance = [][]int{
		{0, 10, 15, 20},
		{5, 0, 9, 10},
		{6, 13, 0, 12},
		{8, 8, 9, 0},
	} // min=35
	distance = [][]int{
		{0, 1, 15, 6},
		{2, 0, 7, 3},
		{9, 6, 0, 12},
		{10, 4, 8, 0},
	} // min=21
	fmt.Printf("res: %d\n", "")
}
4
