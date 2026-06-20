package constants

type contextKey string

const (
	ValidatorContextKey contextKey = "ValidatedPayload"
	UserIdContextKey    contextKey = "UserId"
	ApiKeyContextKey    contextKey = "ApiKey"
)
