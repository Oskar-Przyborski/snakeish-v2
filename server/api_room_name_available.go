package main

import (
	"encoding/json"
	"net/http"
	"snakeish/golang/http_utils"
)

func RoomNameAvailableEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http_utils.WriteError(&w, 400, "missing-parameter", "url should contain 'name' parameter")
		return
	}

	roomNameAvailable := true
	for _, room := range Manager.Rooms {
		if room.RoomName == name {
			roomNameAvailable = false
			break
		}
	}

	response := struct {
		Available bool `json:"available"`
	}{
		Available: roomNameAvailable,
	}

	json, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
