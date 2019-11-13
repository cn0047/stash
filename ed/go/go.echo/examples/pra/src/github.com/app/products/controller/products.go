// Products controller - contains handlers for routes related to products.
// All end-points related to this controller start with prefix "/products".
//
// With purpose to separate concerns here, we just interact with products DAO,
// hence controllers work like user-interface layer,
// all infrastructural problems placed in DAO.

package controller

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/app/products/dao"
)

// GetProduct - find product by id and return product data as payload.
func GetProduct(c echo.Context) error {
	product, exists := dao.GetProductByID(c.Param("id"))

	if exists {
		return c.JSON(http.StatusOK, product)
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

// GetAllProducts - return all known products.
func GetAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, dao.GetProducts())
}
