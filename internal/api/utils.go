package api

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, data map[string]string) {
	// w.Header().Set("Content-Type", "application/json")

	// // FIX: We can only return either statusCode or data.
	// // w.WriteHeader(statusCode)
	// json.NewEncoder(w).Encode(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
