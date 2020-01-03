package main

import (
	"net/http"
)

func main() {
	trgt := "http://realtimelog.herokuapp.com/"

	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, trgt, http.StatusSeeOther)
	})
	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", trgt)
		w.WriteHeader(http.StatusFound)
	})

	http.ListenAndServe(":8080", nil)
}
