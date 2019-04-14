package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/fib", fibonacciHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	msg := "go - ok."
	w.Write([]byte(msg))
}

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	n := r.URL.Query().Get("n")
	i, _ := strconv.Atoi(n)
	result := fibonacci(i)
	fmt.Fprintf(w, "fibonacci %d = %d", i, result)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
