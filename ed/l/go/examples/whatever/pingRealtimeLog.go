package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	URL = "https://realtimelog.herokuapp.com/ping"
)

var (
	HostName = os.Getenv("HOSTNAME")
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
		rtl(map[string]string{"pod": HostName, "payload": r.RequestURI})
		_, err := w.Write([]byte(`ok`))
		if err != nil {
		}
	})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		sig := <-signals
		rtl(map[string]string{"got OS signal": fmt.Sprintf("%#v", sig)})
		panic(fmt.Errorf("have to exit"))
	}()

	http.ListenAndServe(":8080", nil)
}

func bg() {
	id := random(1, 7000)
	for {
		at := time.Now().UTC()
		rtl(map[string]interface{}{"pod": HostName, "id": id, "at": at})
		fmt.Printf("Please open: %s to see new message, at: %s \n", URL, at)
		time.Sleep(10 * time.Second)
	}
}

func rtl(data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
	}
	_, err = http.Post(URL, "application/json", bytes.NewBuffer(j))
	if err != nil {
	}
}
