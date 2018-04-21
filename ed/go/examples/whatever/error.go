package main

import (
	"errors"
	"fmt"
)

func main() {
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
