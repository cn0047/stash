package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/:n", p1)

	e.Logger.Fatal(e.Start(":8080"))
}

func p1(c echo.Context) error {
	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error()+"\n")
	}

	v := n + n
	return c.String(http.StatusOK, fmt.Sprintf("%d\n", v))
}
