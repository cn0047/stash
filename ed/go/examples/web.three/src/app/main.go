package main

import (
	"net/http"

	"app/controller"
	"app/middleware"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8000", &middleware.My{new(middleware.My)})
}
