package main

import (
	"net/http"
	"snakeish/http_utils"
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
	for _, room := range Core.GetRooms() {
		if room.GetName() == name {
			roomNameAvailable = false
			break
		}
	}

	response := struct {
		Available bool `json:"available"`
	}{
		Available: roomNameAvailable,
	}

	http_utils.WriteJSON(w, response)
}
