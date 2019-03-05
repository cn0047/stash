package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n[%s] New request", time.Now())
		_, busy := r.URL.Query()["busy"]
		if busy {
			time.Sleep(time.Second * 9)
		}
		w.Write([]byte("[go] It works!"))
		fmt.Printf("\n[%s] \t Set response", time.Now())
	})

	http.ListenAndServe(":8080", nil)
}

// curl "http://localhost:8080/?busy"
