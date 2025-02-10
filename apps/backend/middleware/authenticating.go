package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ei-sugimoto/adpay/apps/backend/utils"
)

func AuthenticatingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "authorization token must start with 'Bearer '", http.StatusUnauthorized)
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		if tokenString == "" {
			http.Error(w, "authorization token is required", http.StatusUnauthorized)
			return
		}

		userID, err := utils.GetUserIDFromToken(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
