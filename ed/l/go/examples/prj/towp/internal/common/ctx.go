package common

import (
	"context"
)

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

var (
	CtxKeyEnv   = ContextKey("env")
	CtxKeyToken = ContextKey("token")
)

func fromCtx(ctx context.Context, key ContextKey) string {
	val, ok := ctx.Value(key).(string)
	if ok {
		return val
	}

	return ""
}

func GetCtxEnv(ctx context.Context) string {
	return fromCtx(ctx, CtxKeyEnv)
}

func GetCtxToken(ctx context.Context) string {
	return fromCtx(ctx, CtxKeyToken)
}
