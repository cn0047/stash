package main

import (
	"net/http"

	"webapp/controller"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8000", nil)
}
