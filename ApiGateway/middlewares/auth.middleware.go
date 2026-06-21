package middlewares

import (
	"ApiGateway/constants"
	"ApiGateway/helpers"
	"ApiGateway/models"
	"ApiGateway/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type AuthMiddleware struct {
	authServiceURL string
	httpClient     *http.Client
}

func NewAuthMiddleware(authServiceURL string, httpClient *http.Client) *AuthMiddleware {
	return &AuthMiddleware{
		authServiceURL: authServiceURL,
		httpClient:     httpClient,
	}
}

func (m *AuthMiddleware) AuthenticateApiKey() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-KEY")

			if apiKey == "" {
				utils.WriteErrorJsonResponse(w, http.StatusUnauthorized, "missing X-API-Key header")
				return
			}

			parts := strings.Split(apiKey, "_")
			if len(parts) != 3 || parts[0] != "ak" {
				utils.WriteErrorJsonResponse(w, http.StatusUnauthorized, "invalid api key format")
				return
			}

			if len(parts[1]) != 16 || len(parts[2]) != 64 {
				utils.WriteErrorJsonResponse(w, http.StatusUnauthorized, "invalid apikey")
				return
			}
			_, err := m.callValidate(r.Context(), apiKey)
			if err != nil {
				utils.WriteErrorJsonResponse(w, http.StatusUnauthorized, "invalid or revoked api key")
				return
			}

			ctx := context.WithValue(r.Context(), constants.ApiKeyContextKey, apiKey)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (m *AuthMiddleware) callValidate(ctx context.Context, apiKey string) (*models.Principal, error) {
	url := fmt.Sprintf("%s/api/v1/api-key/verify", m.authServiceURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("X-User-ID", strconv.FormatInt(ctx.Value(constants.UserIdContextKey).(int64), 10))
	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call auth service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("auth service returned %d", resp.StatusCode)
	}

	var result models.Principal
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &result, nil
}

func (m *AuthMiddleware) AuthenticateAccesstoken() func(http.Handler) http.Handler {
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
			ctx := context.WithValue(r.Context(), constants.UserIdContextKey, claims.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
