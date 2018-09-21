package route

import (
	"net/http"

	"go-app/controller/home"
)

func Home() {
	http.HandleFunc("/", home.IndexHandler)
	http.HandleFunc("/index", home.IndexHandler)
	http.HandleFunc("/home", home.IndexHandler)
}
