package main

import (
    "net/http"
    "github.com/labstack/echo"
)

type Product struct {
    Name string `json:"name"`
}

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!\n")
    })
    e.GET("/products", Products)
    e.GET("/products/:id", ProductById)

    e.Logger.Fatal(e.Start(":8080"))
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
