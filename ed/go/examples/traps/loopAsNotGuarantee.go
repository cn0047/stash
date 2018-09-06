package main

type T struct {
	msg string
}

var (
	g *T
)

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

func main() {
	go setup()
	for g == nil {
	}
	print(g.msg) // there is no guarantee that it will observe the initialized value for g.msg.
}
