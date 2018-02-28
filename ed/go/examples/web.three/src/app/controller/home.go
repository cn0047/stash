package controller

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type home struct {
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	var query query
	err := dec.Decode(&query)
	if err != nil {
		 fmt.Println("Error: ", err)
	}
	fmt.Printf("%#v\n", query)

	w.Header().Add("Content-Type", "application/json");
	res := response{"OK", 200}
	e := json.NewEncoder(w)
	err2 := e.Encode(res)
	if err2 != nil {
		fmt.Println("Error: ", err2)
	}
}
