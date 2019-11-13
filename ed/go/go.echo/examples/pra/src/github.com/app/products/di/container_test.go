package di

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/app/products/config"
)

var (
	pathToArchive = filepath.Join(config.ApplicationRoot, config.ProductsArchive)
)

// TestInit - test case for Init DI.
func TestInit(t *testing.T) {
	actual := Init()

	if actual != nil {
		t.Errorf("Got: %s, want: nil.", actual)
	}
}

// TestInit - test case for Init DI error.
func TestInitErrorCase(t *testing.T) {
	expected := "ProductsArchive not found by path: /app/products.zip"

	// Rename products archive with purpose to get error during DI Init.
	if _, err := os.Stat(pathToArchive); os.IsNotExist(err) {
		t.Errorf("Path to archive not found, got: %s, want: nil.", err)
		return
	}
	err := os.Rename(pathToArchive, pathToArchive+"_")
	if err != nil {
		t.Errorf("Didn't manage to rename archive, got: %s, want: nil.", err)
	}

	er := Init()
	actual := er.Error()

	if actual != expected {
		t.Errorf("Got: %s, want: %s.", actual, expected)
	}

	// Rename products archive back.
	err = os.Rename(pathToArchive+"_", pathToArchive)
	if err != nil {
		t.Errorf("Didn't manage to rename archive back, got: %s, want: nil.", err)
	}
}
