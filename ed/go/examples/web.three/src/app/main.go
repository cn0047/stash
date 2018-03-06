package main

import (
	"net/http"

	"app/http/controller"
	"app/http/middleware"
)

func main() {
	controller.Startup()
	//http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", new(middleware.RequestInfo))
	//http.ListenAndServe(":8080", &middleware.RequestInfo{new(middleware.OsInfo)})
}
