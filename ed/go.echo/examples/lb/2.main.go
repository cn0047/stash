package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/labstack/echo"
)

type Req struct {
	c    echo.Context
	done chan error
}

var (
	requests = make(chan *Req)
)

func main() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i)
	}

	e := echo.New()

	e.GET("/:n", func(c echo.Context) error {
		done := make(chan error)
		requests <- &Req{c: c, done: done}
		return <-done
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func worker(n int) {
	for {
		select {
		case req := <-requests:
			println("worker #", n)
			req.done <- p2(req.c)
		}
	}
}

func p2(c echo.Context) error {
	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error()+"\n")
	}

	v := n + n
	return c.String(http.StatusOK, fmt.Sprintf("%d\n", v))
}
