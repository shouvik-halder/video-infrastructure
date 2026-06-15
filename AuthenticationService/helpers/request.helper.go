package helpers

import (
	"AuthenticationService/constants"
	"context"
)

func GetPayload[T any](cxt context.Context) (*T, bool) {
	payload, ok := cxt.Value(constants.ValidatorContextKey).(*T)
	return payload, ok
}

func GetUserId(ctx context.Context) (int64, bool) {
    userID, ok := ctx.Value(constants.UserIdContextKey).(int64)
    return userID, ok
}