package middlewares

import (
	"ApiGateway/constants"
	"ApiGateway/utils"
	"context"
	"net/http"
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
