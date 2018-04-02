package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/app/products/controller"
	"github.com/app/products/di"
)

var (
	e = echo.New()
)

// Init DI.
func init() {
	err := di.Init()
	if err != nil {
		e.Logger.Fatal(err)
	}
}

// With purpose to separate concerns here, we have only routes.
// All handlers placed in controller directory.
func main() {
	e.Use(middleware.Logger())

	e.GET("/products", controller.GetAllProducts)
	e.GET("/products/:id", controller.GetProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
