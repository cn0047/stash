package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := runtime.NumGoroutine()
		log.Printf("🔴 %+v", n)
		time.Sleep(5 * time.Second)
		err := json.NewEncoder(w).Encode(map[string]int{"NumGoroutine": n})
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
