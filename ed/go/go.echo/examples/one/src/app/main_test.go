package main

import (
	"github.com/labstack/echo"
	"net/http/httptest"
	"testing"
)

func TestProducts(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/products", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	Products(c)

	actual := rec.Body.String()
	expected := `[{"name":"iphone"},{"name":"ipad"}]`
	if actual != expected {
		t.Errorf("FAIL-1 Got %v want %v", actual, expected)
	}
}

func TestProductById(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/products/macbook", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("macbook")
	ProductById(c)

	actual := rec.Body.String()
	expected := `{"name":"macbook"}`
	if actual != expected {
		t.Errorf("FAIL-2 Got %v want %v", actual, expected)
	}
}
