package lib

import (
	"errors"
)

func Go()  {
	go func() {
		my()
	}()
}

func my() {
	panic(errors.New("some error"))
}
