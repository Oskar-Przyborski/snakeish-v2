package http_utils

import "net/http"

func CheckCors(w *http.ResponseWriter, r *http.Request) bool {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")

	if r.Method == "OPTIONS" {
		(*w).WriteHeader(200)
		return true
	}
	return false
}
