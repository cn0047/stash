package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	URL = "https://realtimelog.herokuapp.com/ping"
)

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	go bg()
	web()
}

func web() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(map[string]string{"payload": r.RequestURI})
		http.Post(URL, "application/json", bytes.NewBuffer(j))
		w.Write([]byte(`ok`))
	})
	http.ListenAndServe(":8080", nil)
}

func bg() {
	id := random(1, 7000)
	for {
		at := time.Now().UTC()
		j, _ := json.Marshal(map[string]interface{}{"id": id, "at": at})
		http.Post(URL, "application/json", bytes.NewBuffer(j))
		fmt.Printf("Please open: %s to see new message, at: %s \n", URL, at)
		time.Sleep(5 * time.Second)
	}
}
