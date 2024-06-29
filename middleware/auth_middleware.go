package middleware

import (
	"context"
	"net/http"
	"strings"
	"recruitment-system/utils"
)

var UserContextKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header missing", http.StatusUnauthorized)
				return
			}

			token := strings.Split(authHeader, "Bearer ")[1]
			user, err := utils.ValidateJWT(token, secretKey)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
