package experiment

import (
	"fmt"
	"net/http"
)

func HTTPError500Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "Error 500.")
}
