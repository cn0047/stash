package app

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"

	"quizengine/app/payload"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Token")
		token = strings.TrimSpace(token)
		if token == "" {
			sendError(w, 400, "X-Token header is missing.")
			return
		}

		claims := &payload.AuthenticationToken{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			log.Errorf("authentication failed, err: %v", err)
			sendError(w, 403, "Authentication failed.")
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, payload.ContextKeyUserID, claims.UserID)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
