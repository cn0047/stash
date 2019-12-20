package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/airbrake/gobrake"
)

const (
	AirBrakeProjectId  = 0
	AirBrakeProjectKey = ""
)

var (
	airbrake = gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
		ProjectId:   AirBrakeProjectId,
		ProjectKey:  AirBrakeProjectKey,
		Environment: "production",
	})
)

func init() {
	airbrake.AddFilter(func(notice *gobrake.Notice) *gobrake.Notice {
		notice.Params["user"] = map[string]string{
			"id":       "1",
			"username": "johnsmith",
			"name":     "John Smith",
		}
		return notice
	})
	airbrake.AddFilter(func(notice *gobrake.Notice) *gobrake.Notice {
		if notice.Context["environment"] == "development" {
			// Ignore notices in development environment.
			return nil
		}
		return notice
	})
}

func main() {
	defer airbrake.Close()
	defer airbrake.NotifyOnPanic()

	two()
}

func two() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		airbrake.Notify(errors.New("have new error here"), nil)
		if _, err := fmt.Fprintf(w, "ok"); err != nil {
			log.Printf("🟥 error: %#v\n", err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Errorf("error1: %w", err))
	}
}

func one() {
	airbrake.Notify(errors.New("operation failed"), nil)
}
