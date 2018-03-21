package main

import "fmt"

func main() {
    one()
    two()
}

func one() {
    scores := make([]int, 10) // slice
    scores2 := make([]int, 0, 10); // slice of length 0 and capacity 10.
    
    scores3 := make([]int, 0, 10)
    //scores3[7] = 903 // panic: runtime error: index out of range
    
    fmt.Printf("%+v \n", scores)
    fmt.Printf("%+v \n", scores2)
    fmt.Printf("%+v \n", scores3)
}

func two() {
    var scores [10]int
    scores[0] = 339
    scores := [4]int{9001, 9333, 212, 33}
    // size from elements in {}
    scores := [...]int{9001, 9333, 212, 33}

    scores = append(scores, 5)

    // foreach
    powers := make([]int, len(saiyans))
    for index, saiyan := range saiyans {
        powers[index] = saiyan.Power
    }
}

/*
[0 0 0 0 0 0 0 0 0 0]
[]
[]
*/
