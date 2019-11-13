package main

type LambdaConfig struct {
	UseRealTimeLog bool
}

var (
	config = &LambdaConfig{UseRealTimeLog: true}
)
