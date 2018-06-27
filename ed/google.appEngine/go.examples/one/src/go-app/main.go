package go_app

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine" // Required external App Engine library
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hw", hwHandler)
	http.HandleFunc("/goon", goonHandler)
	http.HandleFunc("/datastore", datastoreHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/config", configHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/config/tag/", configTagHandler)

	appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, GeTHomeText())
}

func hwHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world!!!))"
	w.Write([]byte(msg))
}
