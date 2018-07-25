package _interface

type FooInterface interface {
	Bar()
}

type MyFoo struct {
}

func (m MyFoo) Bar() {
}

type MyBar struct {
}

func main() {
	var _ FooInterface = (*MyFoo)(nil)
	//var _ FooInterface = (*MyBar)(nil) // cannot use (*MyBar)(nil) (type *MyBar) as type FooInterface
}
