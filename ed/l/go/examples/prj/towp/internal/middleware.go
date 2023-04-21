package internal

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/to-com/wp/internal/common"
)

var (
	ErrEmptyEnv   = errors.New("invalid request. Empty X-Env-Type header")
	ErrEmptyToken = errors.New("invalid request. Empty X-Token header")
)

type PackFn func(ctx context.Context, r *http.Request) (context.Context, error)

func WithPackTokenInCtx(ctx context.Context, r *http.Request) (context.Context, error) {
	token := r.Header.Get("X-Token")
	if token == "" {
		return nil, ErrEmptyToken
	}

	ctx = context.WithValue(ctx, common.CtxKeyToken, token)

	return ctx, nil
}

func WithPackEnvInCtx(ctx context.Context, r *http.Request) (context.Context, error) {
	env := r.Header.Get("X-Env-Type")

	if env == "" {
		if strings.Contains(r.Host, "dev.nonprod.") {
			env = "dev"
		} else if strings.Contains(r.Host, "qai.nonprod.") {
			env = "qai"
		} else if strings.Contains(r.Host, ".uat.") {
			env = "uat"
		} else if strings.Contains(r.Host, ".prod.") {
			env = "prod"
		}
	}

	if env == "" {
		return nil, ErrEmptyEnv
	}

	ctx = context.WithValue(ctx, common.CtxKeyEnv, env)

	return ctx, nil
}

func PackInCtx(ctx context.Context, r *http.Request, packFns ...PackFn) (context.Context, error) {
	for _, fn := range packFns {
		c, err := fn(ctx, r)
		if err != nil {
			return nil, err
		}
		ctx = c
	}

	return ctx, nil
}
