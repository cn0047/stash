package main

var (
	a    string
	done bool
)

func setup1() {
	a = "hello, world"
	done = true
}

func main() {
	go setup1()
	for !done {
	}
	print(a)
}
