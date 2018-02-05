package main

import "fmt"

type Saiyan struct {
    Name string
    Power int
}

func main() {
    goku := Saiyan{"Goku", 9000}
    Super(goku)
    fmt.Println(goku.Power)

    goku2 := &Saiyan{"Goku", 8000}
    Super2(goku2)
    fmt.Println(goku2.Power)

    goku3 := &Saiyan{"Goku", 7000}
    Super3(goku3)
    fmt.Println(goku3.Power)
}

func Super(s Saiyan) {
    s.Power += 10000
}

func Super2(s *Saiyan) {
    s.Power += 10000
}

func Super3(s *Saiyan) {
    s = &Saiyan{"Gohan", 1000}
}
