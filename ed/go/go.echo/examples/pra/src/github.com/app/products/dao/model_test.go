package dao

import (
	"testing"
)

// TestProductModel - silly test, but it works as a sanity test for product model.
func TestProductModel(t *testing.T) {
	model := ProductModel{}

	if model.ID != "" {
		t.Errorf("Got: %s, want: '' (empty string).", model.ID)
	}
}

// TestSetName - test case for method SetName.
func TestSetName(t *testing.T) {
	expected := "test"

	model := ProductModel{}
	model.SetName(expected)

	if model.Name != expected {
		t.Errorf("Got: %s, want: %s.", model.Name, expected)
	}
}
