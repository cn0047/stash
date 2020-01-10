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
	AppVersion = "1.1.3"
	URL        = "https://realtimelog.herokuapp.com/ping"
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

func getData() map[string]interface{} {
	return map[string]interface{}{"v": AppVersion, "host": HostName}
}

func web() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m := getData()
		m["payload"] = r.RequestURI
		j := rtl(m)
		_, err := w.Write(j)
		if err != nil {
			// @TODO: Add error handling.
		}
	})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		sig := <-signals
		s := "@see: https://github.com/golang/go/blob/master/src/syscall/zerrors_darwin_amd64.go#L1391"
		rtl(map[string]string{"got OS signal": fmt.Sprintf("%#v %s", sig, s)})
		panic(fmt.Errorf("have to exit"))
	}()

	http.ListenAndServe(":8080", nil)
}

func bg() {
	id := random(1, 7000)
	for {
		at := time.Now().UTC()
		m := getData()
		m["id"] = id
		m["at"] = at
		rtl(m)
		fmt.Printf("Please open: %s to see new message, at: %s \n", URL, at)
		time.Sleep(10 * time.Second)
	}
}

func rtl(data interface{}) []byte {
	j, err := json.Marshal(data)
	if err != nil {
		// @TODO: Add error handling.
	}
	_, err = http.Post(URL, "application/json", bytes.NewBuffer(j))
	if err != nil {
		// @TODO: Add error handling.
	}

	return j
}
