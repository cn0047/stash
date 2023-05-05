package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Config represents test suite to test Config.
func Test_Config(ts *testing.T) {

	ts.Run("should load ENV name", func(tc *testing.T) {
		// Arrange.
		err := os.Setenv("ENV", "test")
		assert.Nil(tc, err)

		// Act.
		cfg, err := New()
		assert.Nil(tc, err)

		// Assert.
		assert.Equal(tc, "test", cfg.Env)
	})
}
