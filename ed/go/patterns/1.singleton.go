package main

var (
	instance = &singleton{}
)

type singleton struct {
}

func GetInstance() *singleton { // only pointer can be exported outside.
	return instance
}
