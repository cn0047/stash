package main

import (
	"fmt"
	"sort"
)

type Item struct {
	value  float64
	weight float64
}

// fractionalKnapsack represents fractional knapsack problem algorithm.
func fractionalKnapsack(cost []float64, weight []float64, capacity float64) float64 {
	n := len(cost)

	// k holds itemse sorted by cost per 1 unit of wight.
	k := make([]Item, n)
	for i := 0; i < n; i++ {
		k[i] = Item{value: cost[i], weight: weight[i]}
	}
	sort.Slice(k, func(i, j int) bool {
		a := k[i].value / k[i].weight
		b := k[j].value / k[j].weight
		return a > b
	})

	res := 0.0
	currentCapacity := capacity

	for i := 0; i < n; i++ {
		if k[i].weight <= currentCapacity {
			res += k[i].value
			currentCapacity -= k[i].weight
		} else {
			fraction := currentCapacity / k[i].weight
			res += k[i].value * fraction
			break // Knapsack is full.
		}
	}

	return res
}

func main() {
	capacity, cost, weight := 0.0, []float64{}, []float64{}
	capacity, cost, weight = 50.0, []float64{60, 100, 120}, []float64{10, 20, 30}                    // 240
	capacity, cost, weight = 15.0, []float64{10, 5, 15, 7, 6, 18, 3}, []float64{2, 3, 5, 7, 1, 4, 1} //

	fmt.Printf("res: %.2f \n", fractionalKnapsack(cost, weight, capacity))
}
