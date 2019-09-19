package main

import (
	"errors"
	"fmt"
)

var (
	ERR_2 = fmt.Errorf("f2 error.")
)

func main() {
	err := f1()
	fmt.Printf("%v \n", err)
	fmt.Printf("error is ERR_2: %v \n", errors.Is(err, ERR_2))
	fmt.Printf("%v \n", errors.Unwrap(err))
}

func f1() error {
	err := f2()
	if err != nil {
		return fmt.Errorf("f1 error to call next func: %w", err)
	}
	return nil
}

func f2() error {
	return ERR_2
}
