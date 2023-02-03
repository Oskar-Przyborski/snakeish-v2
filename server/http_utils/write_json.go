package http_utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, payload interface{}) bool {
	json, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		return true
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	return false
}
