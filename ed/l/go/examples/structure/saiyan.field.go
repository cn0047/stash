package main

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func main() {
	gohan := &Saiyan{
		Name:  "Gohan",
		Power: 1000,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  9001,
			Father: nil,
		},
	}
	println(gohan.Name, gohan.Father.Name)
}
