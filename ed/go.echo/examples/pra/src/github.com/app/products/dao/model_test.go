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
