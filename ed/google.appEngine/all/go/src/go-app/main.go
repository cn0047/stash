package go_app

import (
	"net/http"

	"google.golang.org/appengine" // Required external App Engine library
)

func init() {
	http.HandleFunc("/", indexHandler)

	appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	msg := "ago - ok"
	w.Write([]byte(msg))
}
