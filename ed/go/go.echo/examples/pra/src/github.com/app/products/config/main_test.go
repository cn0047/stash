package config

import (
	"testing"
)

// TestMainConfig - pretty stupid test, but anyway it checks some configs values.
func TestMainConfig(t *testing.T) {
	if len(ApplicationRoot) == 0 {
		t.Error("Const ApplicationRoot can't be empty.")
	}

	if len(ProductsArchive) == 0 {
		t.Error("Const ProductsArchive can't be empty.")
	}
	if len(DataStorageDirectory) == 0 {
		t.Error("Const DataStorageDirectory can't be empty.")
	}
	if len(ProductsStorageDirectory) == 0 {
		t.Error("Const ProductsStorageDirectory can't be empty.")
	}
}
