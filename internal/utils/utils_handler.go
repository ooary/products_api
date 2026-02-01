package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("content-type", "Application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ReadJson(r *http.Request, data any) {
	json.NewDecoder(r.Body).Decode(data)
}
