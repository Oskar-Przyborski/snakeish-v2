package main

import (
	"net/http"
	"snakeish/http_utils"
)

type RoomPreviewStruct struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Players      int    `json:"players"`
	MaxPlayers   int    `json:"maxPlayers"`
	GameModeName string `json:"gameModeName"`
	GameModeTag  string `json:"gameModeTag"`
}

func GetRoomEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http_utils.WriteError(&w, 400, "missing-parameter", "url should contain 'id' parameter")
		return
	}

	room, _ := Core.GetRoomById(id) // Get room
	if room == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
		return
	}

	http_utils.WriteJSON(w, room.GetPreview())
}
