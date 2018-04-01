package main

import "fmt"

func main() {
    one()
    two()
}

func one() {

}

func two() {
    var scores0 [10]int
    scores0[0] = 339

    scores1 := [4]int{9001, 9333, 212, 33}

    // size from elements in {}
    scores2 := [...]int{9001, 9333, 212, 33}
    //scores2 = append(scores2, 5) // ERROR: first argument to append must be slice; have [4]int

    fmt.Printf("\n scores0 = %#v", scores0)
    fmt.Printf("\n scores1 = %#v", scores1)
    fmt.Printf("\n scores2 = %#v", scores2)

    // foreach
    /*
    powers := make([]int, len(saiyans))
    for index, saiyan := range saiyans {
        powers[index] = saiyan.Power
    }
    */
}

/*
 scores0 = [10]int{339, 0, 0, 0, 0, 0, 0, 0, 0, 0}
 scores1 = [4]int{9001, 9333, 212, 33}
 scores2 = [4]int{9001, 9333, 212, 33}
*/
