package main

import "fmt"

type VO struct {
    Value int
}

func main() {
    var vo VO
    f(&vo)
    fmt.Printf("%+v\n", vo)
}

func f(vo *VO) {
    vo.Value = 200
}
