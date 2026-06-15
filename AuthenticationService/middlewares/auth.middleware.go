package middlewares

import (
	"AuthenticationService/constants"
	"AuthenticationService/helpers"
	"AuthenticationService/utils"
	"context"
	"net/http"
	"strings"
)

func ValidateToken() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				utils.WriteErrorResponseJson(w, http.StatusUnauthorized, "missing authorization header")
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				utils.WriteErrorResponseJson(w, http.StatusUnauthorized, "invalid authorization header")
				return
			}

			claims, err := helpers.ValidateToken(parts[1])
			if err != nil {
				utils.WriteErrorResponseJson(w, http.StatusUnauthorized, "invalid token")
				return
			}

			ctx := context.WithValue(r.Context(), constants.UserIdContextKey, claims.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
