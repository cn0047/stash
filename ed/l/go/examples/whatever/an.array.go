package main

import (
	"fmt"
)

func main() {
	//one()
	//one2()
	//two()
	updateArray()
}

func updateArray() {
	a := [3]int{0, 1, 2}
	incEach(a)
	fmt.Printf("[updateArray] %+v \n", a)           // [updateArray] [0 1 2]
	fmt.Printf("[updateArray] %+v \n", incEach2(a)) // [updateArray] [0 2 4]
}

func incEach(a [3]int) {
	for k, v := range a {
		a[k] = v + v
	}
}

func incEach2(a [3]int) [3]int {
	for k, v := range a {
		a[k] = v + v
	}
	return a
}

func one() {
	var a1 = [5]string{2: "Bob", 3: "Chris", 4: "Ron"}
	fmt.Printf("\n a1 = %#v", a1) // a1 = [5]string{"", "", "Bob", "Chris", "Ron"}
}

func one2() {
	// Delete item at index i:
	i := 2
	a2 := []int{1, 2, 3, 4, 5}
	a2 = append(a2[:i], a2[i+1:]...)
	fmt.Printf("\n a2 = %#v", a2) // a2 = []int{1, 2, 4, 5}
}

func two() {
	var scores0 [10]int
	scores0[0] = 339

	scores1 := [4]int{9001, 9333, 212, 33}

	// size from elements in {}
	scores2 := [...]int{9001, 9333, 212, 33} // this is array
	//scores2 = append(scores2, 5) // ERROR: first argument to append must be slice; have [4]int

	fmt.Printf("\n scores0 = %#v", scores0) // scores0 = [10]int{339, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Printf("\n scores1 = %#v", scores1) // scores1 = [4]int{9001, 9333, 212, 33}
	fmt.Printf("\n scores2 = %#v", scores2) // scores2 = [4]int{9001, 9333, 212, 33}
}
