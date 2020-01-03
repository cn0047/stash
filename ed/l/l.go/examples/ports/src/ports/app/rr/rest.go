package rr

import (
	"encoding/json"
	"net/http"
	"strings"
)

func StartREST() {
	http.HandleFunc("/ports/", getPortByID)

	http.ListenAndServe(RESTPort, nil)
}

func getPortByID(w http.ResponseWriter, r *http.Request) {
	// todo: discuss way to get ID from URL and id validation.
	path := r.URL.Path
	pos := strings.LastIndex(path, "/") + 1
	id := path[pos:]

	e := json.NewEncoder(w)
	port, err := GetPort(id)
	if err != nil {
		e.Encode(map[string]string{"error": "failed to get port, error:" + err.Error(), "data": ""})
		return
	}

	e.Encode(map[string]interface{}{"error": "", "data": port})
}
