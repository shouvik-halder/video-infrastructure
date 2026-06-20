package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteErrorJsonResponse(w http.ResponseWriter, status int, message string) {
	WriteJsonResponse(w, status, message)
}
