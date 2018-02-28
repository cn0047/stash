package controller

import (
	"net/http"

	"webapp/model"
)

type home struct {
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Home page."))
    w.Write([]byte("\n\nProducts: \n"))

    products := model.GetProducts()
	for _, p := range products {
		w.Write([]byte(p.Name))
	}

    w.Write([]byte("\n\nCategories: \n"))
	categories := model.GetCategories()
	for _, c := range categories {
		w.Write([]byte(c.Title))
	}
}
