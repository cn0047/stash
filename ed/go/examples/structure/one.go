package main

type MyType struct {
	N int
}

func (mt MyType) inc1(v int) {
	mt.N = mt.N + v
}

func (mt *MyType) inc2(v int) {
	mt.N = mt.N + v
}

func main() {
	mt := MyType{N: 10}
	mt.inc1(3)
	mt.inc2(5)
	println(mt.N) // 15
}
