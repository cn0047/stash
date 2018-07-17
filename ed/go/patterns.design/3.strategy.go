package main

import "fmt"

type ConcatenateStrategy interface {
	Add(a interface{}, b interface{}) interface{}
}

type String struct {
}

func (s String) Add(a interface{}, b interface{}) interface{} {
	return string(a.(string) + b.(string))
}

type Number struct {
}

func (n Number) Add(a interface{}, b interface{}) interface{} {
	return a.(int) + b.(int)
}

type Calculator struct {
	strategy ConcatenateStrategy
}

func NewCalculator(c ConcatenateStrategy) Calculator {
	return Calculator{strategy: c}
}

func (c Calculator) Concatenate(a interface{}, b interface{}) interface{} {
	return c.strategy.Add(a, b)
}

func main() {
	c1 := NewCalculator(String{})
	fmt.Printf("String: %+v \n", c1.Concatenate("1", "2"))

	c2 := NewCalculator(Number{})
	fmt.Printf("Number: %+v \n", c2.Concatenate(1, 2))
}
