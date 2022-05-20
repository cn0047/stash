package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//float64Round()
	//float64SortSlice()
	//float64Remainder()
	//float64Parts()
	float64Trap1()
}

func float64Trap1() {
	var x float64 = 0.24
	var y float64 = 0.2
	r := x - y                                               //  0.03999999999999998
	fmt.Printf("[float64Trap1] %+v \n", r)                   // 0.03999999999999998
	fmt.Printf("[float64Trap1] %+v \n", 0.24-0.2)            // 0.04
	fmt.Printf("[float64Trap1] %+v \n", roundTo2AfterDot(r)) // 0.04
}

func float64SortSlice() {
	a := []float64{0.3, 0.1, 0.01, 0.03, 5}
	sort.Float64s(a)
	fmt.Printf("[float64SortSlice] %+v \n", a) // [0.01 0.03 0.1 0.3 5]
}

func float64Remainder() {
	var x float64 = 0.10
	var y float64 = 0.03
	//r := x % y // invalid operation: operator % not defined on x (variable of type float64)
	r1 := math.Mod(x, y)                                     // 0.010000000000000009
	r2 := math.Remainder(x, y)                               // 0.010000000000000009
	fmt.Printf("[remainder 1] %+v \n", roundTo2AfterDot(r1)) // 0.01
	fmt.Printf("[remainder 2] %+v \n", roundTo2AfterDot(r2)) // 0.01
}

func float64Parts() {
	var x float64 = 0.10
	var y float64 = 0.03
	r := x / y                                         // 3.3333333333333335
	valBeforeDot, valAfterDot := math.Modf(float64(r)) // 3, 0.3333333333333335
	valAfterDotRounded := roundTo2AfterDot(valAfterDot)
	fmt.Printf("valBeforeDot: %#v valAfterDotRounded: %#v \n", valBeforeDot, valAfterDotRounded) // 3, 0.33
}

func float64Round() {
	fmt.Printf("[round 1] %+v \n", roundTo2AfterDot(0.031)) // 0.03
	fmt.Printf("[round 2] %+v \n", roundTo2AfterDot(0.034)) // 0.03
	fmt.Printf("[round 3] %+v \n", roundTo2AfterDot(0.035)) // 0.04
	fmt.Printf("[round 4] %+v \n", roundTo2AfterDot(0.039)) // 0.04
}

func roundTo2AfterDot(v float64) float64 {
	return math.Round(v*100) / 100
}
