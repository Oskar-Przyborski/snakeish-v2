package main

import (
	"encoding/json"
	"net/http"
	"snakeish/golang/http_utils"
)

type RoomPreviewStruct struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	PlayersCount int    `json:"playersCount"`
	GameModeName string `json:"gameModeName"`
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

	room := Manager.GetRoomById(id) // Get room
	if room == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
		return
	}

	// Convert room to json string
	json, err := json.Marshal(RoomPreviewStruct{
		Id:           room.Id,
		Name:         room.RoomName,
		PlayersCount: room.getPlayersCount(),
		GameModeName: "classic", //TODO change gamemodename if you add gamemodes
	})
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
