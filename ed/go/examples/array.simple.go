package main

import "fmt"

func main() {
    scores := make([]int, 10) // slice
    scores2 := make([]int, 0, 10); // slice of length 0 and capacity 10.
    
    scores3 := make([]int, 0, 10)
    //scores3[7] = 903 // panic: runtime error: index out of range
    
    fmt.Printf("%+v \n", scores)
    fmt.Printf("%+v \n", scores2)
    fmt.Printf("%+v \n", scores3)
}

/*
[0 0 0 0 0 0 0 0 0 0]
[]
[]
*/
