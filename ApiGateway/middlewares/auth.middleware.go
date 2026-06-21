package middlewares

import (
	"ApiGateway/constants"
	"ApiGateway/helpers"
	"ApiGateway/utils"
	"context"
	"net/http"
	"strconv"
	"strings"
)

func AuthenticateApiKey() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-KEY")
			parts := strings.Split(apiKey, "_")
			if parts[0] != "ak" {
				utils.WriteErrorJsonResponse(w, http.StatusUnauthorized, "invalid apikey format")
				return
			}

			if len(parts[1]) != 16 || len(parts[2]) != 64 {
				utils.WriteErrorJsonResponse(w, http.StatusUnauthorized, "invalid apikey")
				return
			}

			ctx := context.WithValue(r.Context(), constants.ApiKeyContextKey, apiKey)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthenticateAccesstoken() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				utils.WriteErrorJsonResponse(w, http.StatusUnauthorized, "missing authorization header")
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				utils.WriteErrorJsonResponse(w, http.StatusUnauthorized, "invalid authorization header")
				return
			}

			claims, err := helpers.ValidateToken(parts[1])
			if err != nil {
				utils.WriteErrorJsonResponse(w, http.StatusUnauthorized, "invalid token")
				return
			}

			r.Header.Set("X-User-ID", strconv.FormatInt(claims.UserID, 10))
			ctx := context.WithValue(r.Context(), constants.UserIdContextKey, claims.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
