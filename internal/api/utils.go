package api

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	// FIX: We can only return either statusCode or data.
	// w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
