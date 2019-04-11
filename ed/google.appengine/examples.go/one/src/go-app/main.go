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
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/user/oauth", userOAuthHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/config", configHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/template/", templateHandler)
	http.HandleFunc("/config/tag/", configTagHandler)

	appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	linksHTML := `
		<a href="/">home</a>
		<a href="hw">hw</a>
		<a href="goon">goon</a>
		<a href="datastore">datastore</a>
		<a href="user">user</a>
		<a href="user/oauth">user OAuth</a>
		<a href="search">search</a>
		<a href="config">config</a>
		<a href="config/">config/</a>
		<a href="template/">template/</a>
		<a href="config/tag">config/tag</a>
		<hr>
	`

	fmt.Fprintf(w, linksHTML)
	fmt.Fprintf(w, GeTHomeText())
}

func hwHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world!!!))"
	w.Write([]byte(msg))
}
