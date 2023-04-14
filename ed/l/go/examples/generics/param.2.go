package main

func min[P interface{ ~int64 | ~float64 }](x, y P) P {}

func min[P ~int64 | ~float64](x, y P) P {}
