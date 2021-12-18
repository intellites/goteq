package provider

import (
	"encoding/json"
	"net/http"
)

// Success makes the response with payload as json format
func Success(w http.ResponseWriter, status int, payload interface{}) {
	// Set content type as application/json
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(status)
	w.Write([]byte(response))
}

// Error makes the error response with payload as json format
func Error(w http.ResponseWriter, code int, message string) {
	Success(w, code, map[string]string{"error": message})
}
