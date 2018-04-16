package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := Saiyan{"Goku", 9000}
	Super(goku)
	fmt.Println(goku.Power) // 9000

	goku2 := &Saiyan{"Goku", 8000}
	Super2(goku2)
	fmt.Println(goku2.Power) // 18000

	goku3 := &Saiyan{"Goku", 7000}
	Super3(goku3)
	fmt.Println(goku3.Power) // 7000

	goku4 := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	fmt.Println("4:", goku4) // 4: {Goku 9000}

	goku5 := Saiyan{}
	fmt.Println("5:", goku5) // 5: { 0}
	goku5.Power = 5
	fmt.Println("5.2:", goku5) // 5.2: { 5}
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
