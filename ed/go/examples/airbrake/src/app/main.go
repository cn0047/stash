package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/airbrake/gobrake"
)

var (
	airbrake = gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
		ProjectId:   0,
		ProjectKey:  "",
		Environment: "production",
	})
)

func init() {
	return
	airbrake.AddFilter(func(notice *gobrake.Notice) *gobrake.Notice {
		notice.Params["user"] = map[string]string{
			"id":       "1",
			"username": "johnsmith",
			"name":     "John Smith",
		}
		return notice
	})
}

func main() {
	two()
}

func two() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		airbrake.Notify(errors.New("have new error here"), nil)
		if _, err := fmt.Fprintf(w, "ok"); err != nil {
			log.Printf("🟥 error: %#v\n", err)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func one() {
	defer airbrake.Close()
	defer airbrake.NotifyOnPanic()

	airbrake.Notify(errors.New("operation failed"), nil)
}
