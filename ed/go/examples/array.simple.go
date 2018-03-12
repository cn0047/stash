package main

import "fmt"

func main() {
    scores := make([]int, 10)
    scores2 := make([]int, 0, 10)
    
    scores3 := make([]int, 0, 10)
    // scores3[7] = 903
    
    fmt.Printf("%+v \n", scores)
    fmt.Printf("%+v \n", scores2)
    fmt.Printf("%+v \n", scores3)
}
