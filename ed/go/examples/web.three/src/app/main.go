package main

import (
	"net/http"

	"app/di"
	"app/http/controller"
	//"app/http/middleware"
)

func main() {
	di.Init()
	controller.Startup()
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe(":8080", new(middleware.RequestInfo))
	//http.ListenAndServe(":8080", &middleware.RequestInfo{new(middleware.OsInfo)})
}
