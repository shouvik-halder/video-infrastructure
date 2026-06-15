package validator

import (
	"AuthenticationService/constants"
	"AuthenticationService/utils"
	"context"
	"net/http"
)

func ValidateRequest[T any]() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var payload T

			if err := utils.ReadRequestJson(r, &payload); err != nil {
				utils.WriteErrorResponseJson(w, http.StatusBadRequest, err.Error())
				return
			}

			if err := utils.Validator.Struct(payload); err != nil {
				utils.WriteErrorResponseJson(w, http.StatusUnprocessableEntity, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), constants.ValidatorContextKey, &payload)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
