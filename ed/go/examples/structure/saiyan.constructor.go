package main

import "fmt"

type Saiyan struct {
    Type string
    Name string
    Power int
}

func NewSaiyan(name string, power int) *Saiyan {
    return &Saiyan{
        Type: "default", // won't workß
        Name: name,
        Power: power,
    }
}

// or

// func NewSaiyan(name string, power int) Saiyan {
//     return Saiyan{
//         Name: name,
//         Power: power,
//     }
// }

func main() {
    goku := new(Saiyan)
    goku.Name = "goku1"
    goku.Power = 9001
    fmt.Println(goku.Power, *goku)

    goku2 := &Saiyan {
        Name: "goku2",
        Power: 9000,
    }
    fmt.Println(goku2.Power, *goku2)
}
