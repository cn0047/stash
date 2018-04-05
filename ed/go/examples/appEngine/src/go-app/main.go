package go_app

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine" // Required external App Engine library
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/goon", goonHandler)
	http.HandleFunc("/config", configHandler)

	appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, GeTHomeText())
}
