package middlewares

import (
	"AuthenticationService/constants"
	"AuthenticationService/utils"
	"context"
	"net/http"
	"strconv"
)

func ExtractUserId() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := r.Header.Get("X-User-ID")
			if userID == "" {
				utils.WriteErrorResponseJson(w, http.StatusUnauthorized, "missing X-User-ID header")
				return
			}

			userIDInt, err := strconv.ParseInt(userID, 10, 64)
			if err != nil {
				utils.WriteErrorResponseJson(w, http.StatusUnauthorized, "invalid X-User-ID header")
				return
			}

			ctx := context.WithValue(r.Context(), constants.UserIdContextKey, userIDInt)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ExtractApiKey() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")
			ctx := context.WithValue(r.Context(), constants.ApiKeyContextKey, apiKey)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
