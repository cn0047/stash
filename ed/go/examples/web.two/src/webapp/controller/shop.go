package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"webapp/model"
)

type shop struct {
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
	http.HandleFunc("/shop/", s.handleShop)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`)
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1])
		s.handleCategory(w, r, categoryID)
	} else {
		categories := model.GetCategories()
		w.Write([]byte("Categories list page. Data printed to terminal."))
		fmt.Println(categories)
	}
}

func (s shop) handleCategory(w http.ResponseWriter, r *http.Request, categoryID int) {
	products := model.GetProductsForCategory(categoryID)
	w.Write([]byte("Category page. Data printed to terminal."))
	fmt.Println(products)
}
