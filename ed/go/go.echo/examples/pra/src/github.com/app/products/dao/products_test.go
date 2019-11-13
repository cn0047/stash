// Integration test suite for products DAO.

package dao

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/app/products/config"
)

var (
	pathToArchive = filepath.Join(config.ApplicationRoot, config.ProductsArchive)

	testPathToArchive         = "/tmp/test.products.zip"
	testPathToStorage         = "/tmp/test.data"
	testPathToProductsStorage = filepath.Join(testPathToStorage, config.ProductsStorageDirectory)
)

// TestInit - test case for Init.
func TestInit(t *testing.T) {
	actual := Init()

	if actual != nil {
		t.Errorf("Got: %s, want: nil.", actual)
	}
}

// TestInitProductsStorageArchiveNotExist - test case for InitProductsStorage, error: products archive not found.
// This test case important for InitProductsStorage coverage.
func TestInitProductsStorageArchiveNotExist(t *testing.T) {
	expected := "ProductsArchive not found by path: {error}"

	err := InitProductsStorage("{error}", "", "")
	actual := err.Error()

	if actual != expected {
		t.Errorf("Got: %s, want: %s.", actual, expected)
	}
}

// TestInitProductsStorageProductsStorageDirectoryNotExist - test case for InitProductsStorage,
// error: products storage directory not found.
func TestInitProductsStorageProductsStorageDirectoryNotExist(t *testing.T) {
	expected := "ProductsStorageDirectory not found by path: {error}"

	err := InitProductsStorage(pathToArchive, testPathToStorage, "{error}")
	actual := err.Error()

	if actual != expected {
		t.Errorf("Got: %s, want: %s.", actual, expected)
	}
}

// TestUnpackArchiveArchiveNotExist - test case for UnpackArchive, error: archive not found.
func TestUnpackArchiveArchiveNotExist(t *testing.T) {
	expected := "ProductsArchive not found by path: {error}"

	err := UnpackArchive("{error}", "")
	actual := err.Error()

	if actual != expected {
		t.Errorf("Got: %s, want: %s.", actual, expected)
	}
}

// TestUnpackArchiveRemoveDataStorageDirectory - test case for UnpackArchive,
// case when data storage directory must be deleted.
func TestUnpackArchiveRemoveDataStorageDirectory(t *testing.T) {
	expected := "zip: not a valid zip file"

	// Create stubs for products archive and data storage directory.
	os.OpenFile(testPathToArchive, os.O_RDONLY|os.O_CREATE, 0666)
	os.Mkdir(testPathToStorage, 0666)

	err := UnpackArchive(testPathToArchive, testPathToStorage)
	actual := err.Error()

	if actual != expected {
		t.Errorf("Got: %s, want: %s.", actual, expected)
	}
	if _, er := os.Stat(testPathToStorage); er == nil {
		t.Errorf("Storage directory: %s wasn't removed.", testPathToStorage)
	}

	// Clear stub for products archive.
	err = os.Remove(testPathToArchive)
	if err != nil {
		t.Errorf("Stub for products archive: %s wasn't removed.", testPathToArchive)
	}
}

// TestLoadProductsProductsStorageDirectoryNotExist - test case for LoadProducts,
// error: products storage directory not found.
func TestLoadProductsProductsStorageDirectoryNotExist(t *testing.T) {
	expected := "ProductsStorageDirectory not found by path: {error}"

	err := LoadProducts("{error}")
	actual := err.Error()

	if actual != expected {
		t.Errorf("Got: %s, want: %s.", actual, expected)
	}
}

// TestLoadProductsFailedParseFile - test case for LoadProducts, error: failed parse file.
func TestLoadProductsFailedParseFile(t *testing.T) {
	expected := "failed parse file: /tmp/test.data/products/error.json"

	// Create stubs for data storage and products storage directories.
	os.Mkdir(testPathToStorage, 0666)
	os.Mkdir(testPathToProductsStorage, 0666)

	// Create stub for corrupted JSON file.
	pathToCorruptedProduct := filepath.Join(testPathToProductsStorage, "error.json")
	err := ioutil.WriteFile(pathToCorruptedProduct, []byte("error"), 0666)
	if err != nil {
		t.Errorf("Failed to create stub for corrupted JSON file: %s", testPathToProductsStorage)
	}

	er := LoadProducts(testPathToProductsStorage)
	actual := er.Error()

	if actual != expected {
		t.Errorf("Got: %s, want: %s.", actual, expected)
	}

	// Clear stubs for data storage directory, products storage directory and for corrupted JSON file.
	err = os.RemoveAll(testPathToStorage)
	if err != nil {
		t.Errorf("Stub for data storage: %s wasn't removed.", testPathToStorage)
	}
}

// TestGetProductById - test case for GetProductById.
func TestGetProductById(t *testing.T) {
	_, exists := GetProductByID("{id}")

	if exists {
		t.Errorf("Got: %t, want: %t.", true, false)
	}
}

// TestGetProducts - test case for GetProducts.
func TestGetProducts(t *testing.T) {
	products := GetProducts()

	if products == nil {
		t.Errorf("Got: nil, want: map.")
	}
}
