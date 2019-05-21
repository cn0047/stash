package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSwaps([]int{2, 1, 3, 4}) == 1)
	fmt.Println(minimumSwaps([]int{2, 1, 4, 3}) == 2)
	fmt.Println(minimumSwaps([]int{4, 3, 1, 2}) == 3)
	fmt.Println(minimumSwaps([]int{2, 3, 4, 1, 5}) == 3)
	fmt.Println(minimumSwaps([]int{1, 3, 5, 2, 4, 6, 7}) == 3)
}

// minimumSwaps returns minimum required swaps to sort array.
func minimumSwaps(a []int) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			position := a[i] - 1
			// swap
			v := a[i]
			a[i] = a[position]
			a[position] = v
			// update values
			count++
			i--
		}
	}
	return count
}
