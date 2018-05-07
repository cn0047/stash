package main

import "fmt"

type CarInterface interface {
	Drive() error
}

type Builder interface {
	Build() CarInterface
}

type Director struct {
	Builder Builder
}

func (d Director) Construct() CarInterface {
	return d.Builder.Build()
}

type FerrariBuilder struct {
}

func (fb FerrariBuilder) Build() CarInterface {
	f := Ferrari{Speed: "349 km/h", Wheels: "sport type"}
	// Create engine and init into car
	// Create tyres, etc...
	return f
}

type Ferrari struct {
	Speed  string
	Wheels string
}

func (f Ferrari) Drive() error {
	fmt.Printf("\nThis is Ferrari with sppeed: %s", f.Speed)
	return nil
}

func main() {
	d := Director{FerrariBuilder{}}
	c := d.Construct()
	c.Drive()
}
