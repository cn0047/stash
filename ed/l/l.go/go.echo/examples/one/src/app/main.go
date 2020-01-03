package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Product struct {
	Name string `json:"name"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/ok", Ok, MyMiddleware("ok"))
	e.GET("/nook", Ok, MyMiddleware("nook"))
	e.GET("/products", Products)
	e.GET("/products/:id", ProductById)

	e.Logger.Fatal(e.Start(":8080"))
}

func Ok(c echo.Context) error {
	return c.String(http.StatusOK, "OK!\n"+c.Get("myKey").(string))
}

func MyMiddleware(param string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			if param != "ok" {
				return ctx.JSON(http.StatusOK, "MyMiddleware param isn't ok, actual value: "+param)
			}
			ctx.Set("myKey", "[mySecretValue]")
			return next(ctx)
		}
	}
}

func Products(c echo.Context) error {
	p := []Product{
		{Name: "iphone"},
		{Name: "ipad"},
	}
	return c.JSON(http.StatusOK, p)
}

func ProductById(c echo.Context) error {
	p := Product{Name: c.Param("id")}
	return c.JSON(http.StatusOK, p)
}
