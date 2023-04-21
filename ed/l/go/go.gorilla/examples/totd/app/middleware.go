package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/to-com/poc-td/app/payload"
)

func headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f := `{"errors":[{"error":"%s","ref":"%s","type":"bad request"}]}`

		token := r.Header.Get("X-Token")
		if token == "" {
			res := fmt.Sprintf(f, "X-Token header is missing.", "header")
			_, _ = w.Write([]byte(res))
			return
		}

		retailer := r.Header.Get("X-Retailer-Id")
		if retailer == "" {
			res := fmt.Sprintf(f, "X-Retailer-Id header is missing.", "header")
			_, _ = w.Write([]byte(res))
			return
		}

		env := r.Header.Get("X-Env-Type")
		if env == "" {
			res := fmt.Sprintf(f, "X-Env-Type header is missing.", "header")
			_, _ = w.Write([]byte(res))
			return
		}

		mfc := r.Header.Get("X-Mfc")
		if mfc == "" {
			res := fmt.Sprintf(f, "X-Mfc header is missing.", "header")
			_, _ = w.Write([]byte(res))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, payload.ContextKeyToken, token)
		ctx = context.WithValue(ctx, payload.ContextKeyRetailer, retailer)
		ctx = context.WithValue(ctx, payload.ContextKeyEnv, env)
		ctx = context.WithValue(ctx, payload.ContextKeyMfc, mfc)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(payload.ContextKeyToken)
		if token == "" {
			res := fmt.Sprintf(
				`{"errors":[{"error":"%s","ref":"%s","type":"bad request"}]}`,
				"X-Token header is missing.", "header",
			)
			_, _ = w.Write([]byte(res))
			return
		}

		// @TODO: check permissions.

		next.ServeHTTP(w, r)
	})
}
