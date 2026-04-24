package middleware

import (
	"context"
	"net/http"
	"strings"

	"simple-clothes-store/internal/utils"
)

type contextKey string

const ContextUsernameKey contextKey = "username"


func AdminAuth(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, `{"error": "missing authorization header"}`, http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				http.Error(w, `{"error": "invalid authorization header format"}`, http.StatusUnauthorized)
				return
			}

			claims, err := utils.ValidateToken(parts[1], jwtSecret)
			if err != nil {
				http.Error(w, `{"error": "invalid or expired token"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ContextUsernameKey, claims.Username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
