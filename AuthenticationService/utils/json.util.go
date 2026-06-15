package utils

import (
	"AuthenticationService/models"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = newValidator()
}

func newValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func ReadRequestJson(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(result)
}

func WriteResponseJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteErrorResponseJson(w http.ResponseWriter, status int, message string) {
	WriteResponseJson(w, status, models.ErrorResponse{Message: message})
}
