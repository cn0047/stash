package main

type Saiyan struct {
    Name string
    Power int
}

func NewSaiyan(name string, power int) *Saiyan {
    return &Saiyan{
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
    goku.Name = "goku"
    goku.Power = 9001
    println(goku.Power)

    goku2 := &Saiyan {
        Name: "goku",
        Power: 9000,
    }
    println(goku2.Power)
}
