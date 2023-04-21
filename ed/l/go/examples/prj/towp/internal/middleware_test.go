package internal

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/to-com/wp/internal/common"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPackInCtxNoTokenHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	_, err := PackInCtx(context.Background(), request, WithPackTokenInCtx)

	assert.ErrorIs(t, err, ErrEmptyToken)
}

func TestPackInCtxNoEnvHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	_, err := PackInCtx(context.Background(), request, WithPackEnvInCtx)

	assert.ErrorIs(t, err, ErrEmptyEnv)
}

func TestPackInCtxOK(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	request.Header.Add("X-Token", "123")
	request.Header.Add("X-Env-Type", "dev")

	ctx, err := PackInCtx(context.Background(), request, WithPackTokenInCtx, WithPackEnvInCtx)

	assert.Nil(t, err)
	assert.Equal(t, ctx.Value(common.CtxKeyToken), "123")
	assert.Equal(t, ctx.Value(common.CtxKeyEnv), "dev")
}
