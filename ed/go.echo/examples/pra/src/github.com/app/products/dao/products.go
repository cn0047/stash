// DAO - Data Access Object.
// This layer aware of project infrastructure,
// and interacts with products storage, which is based on products.zip archive.

package dao

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/app/lib"
	"github.com/app/products/config"
)

var (
	// Storage which contains all products.
	products = make(map[string]ProductModel)
)

// Init - initialize products DAO.
// Prepare dependencies and run InitProductsStorage.
// If you need initialize something else related to products DAO - this method the best place for that.
func Init() (err error) {
	pathToArchive := filepath.Join(config.ApplicationRoot, config.ProductsArchive)
	pathToStorage := filepath.Join(config.ApplicationRoot, config.DataStorageDirectory)
	pathToProductsStorage := filepath.Join(pathToStorage, config.ProductsStorageDirectory)

	return InitProductsStorage(pathToArchive, pathToStorage, pathToProductsStorage)
}

// InitProductsStorage - load products from archive into products DAO memory.
func InitProductsStorage(pathToArchive string, pathToStorage string, pathToProductsStorage string) (err error) {
	err = UnpackArchive(pathToArchive, pathToStorage)
	if err != nil {
		return
	}

	err = LoadProducts(pathToProductsStorage)
	if err != nil {
		return
	}

	return
}

// UnpackArchive - Unpack products archive into target storage directory.
func UnpackArchive(pathToArchive string, pathToStorage string) (err error) {
	if _, er := os.Stat(pathToArchive); os.IsNotExist(er) {
		return errors.New("ProductsArchive not found by path: " + pathToArchive)
	}

	if _, er := os.Stat(pathToStorage); er == nil {
		e := os.RemoveAll(pathToStorage)
		if e != nil {
			return e
		}
	}

	err = lib.Unzip(pathToArchive, pathToStorage)

	return
}

// LoadProducts - load products from unpacked archive.
func LoadProducts(pathToProductsStorage string) (err error) {
	if _, er := os.Stat(pathToProductsStorage); os.IsNotExist(er) {
		return errors.New("ProductsStorageDirectory not found by path: " + pathToProductsStorage)
	}

	files, er := filepath.Glob(filepath.Join(pathToProductsStorage, "*.json"))
	if er != nil {
		return er
	}

	for _, file := range files {
		jsonProduct, er := ioutil.ReadFile(file)
		if er != nil {
			return er
		}

		var product ProductModel
		er = json.Unmarshal(jsonProduct, &product)
		if er != nil {
			return errors.New("failed parse file: " + file)
		}

		products[product.ID] = product
	}

	return
}

// GetProductByID - get product by id.
func GetProductByID(id string) (ProductModel, bool) {
	p, exists := products[id]

	return p, exists
}

// GetProducts - get all product as array.
func GetProducts() (productsList []ProductModel) {
	// The only way to return array of products.
	for _, p := range products {
		productsList = append(productsList, p)
	}

	return
}
