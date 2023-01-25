package http_utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseStruct struct {
	Name        string `json:"errorName"`
	Description string `json:"errorDescription"`
}

func WriteError(w *http.ResponseWriter, statusCode int, errorName string, errorDescription string) {
	(*w).WriteHeader(statusCode)
	(*w).Header().Add("Content-Type", "application/json")
	json, err := json.Marshal(ErrorResponseStruct{
		Name:        errorName,
		Description: errorDescription,
	})
	if err != nil {
		(*w).Write([]byte("ah, error while converting error message response"))
		return
	}
	(*w).Write(json)
}
