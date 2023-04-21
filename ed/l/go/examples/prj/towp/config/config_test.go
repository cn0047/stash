package config

import (
	"gotest.tools/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	cfg, err := Load()
	assert.NilError(t, err)
	assert.Equal(t, "test-project", cfg.ProjectID)
}
