package main

import (
	"errors"
	"fmt"
)

func main() {
	//one()
	// two()
	f3()
}

func f3() {
	e := fmt.Errorf("err1")
	println("[f3]", e.Error()) // [f3] err1
}

func two() {
	//ErrorCustom1 := fmt.Errorf("err1")
	//ErrorCustom2 := fmt.Errorf("err2")
	////ErrorCustom3 := fmt.Errorf("err3")
	//err := ErrorCustom1
	//if verr, ok := err.(*okjwt.ValidationError); ok {
	//if err.Er & (ErrorCustom1 | ErrorCustom2) {
	//}
}

func one() {
	fmt.Println(sqr(-1))
	fmt.Println(sqr(2))
	fmt.Println(service(-2))
}

func service(n int) (int, error) {
	r, err := sqr(n)
	if err != nil {
		return 0, err
	}

	return r, nil
}

func sqr(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("n - must be positive integer")
	}

	return n * n, nil
}

/*
0 n - must be positive integer
4 <nil>
*/
