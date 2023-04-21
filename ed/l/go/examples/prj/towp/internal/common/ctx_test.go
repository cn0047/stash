package common

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromCtx(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyEnv, "test-env")
	ctx = context.WithValue(ctx, CtxKeyToken, "test-token")

	assert.Equal(t, "test-env", fromCtx(ctx, CtxKeyEnv))
	assert.Equal(t, "test-token", fromCtx(ctx, CtxKeyToken))
	assert.Equal(t, "test-env", GetCtxEnv(ctx))
	assert.Equal(t, "test-token", GetCtxToken(ctx))
}
