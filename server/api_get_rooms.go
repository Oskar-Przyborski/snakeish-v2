package main

import (
	"encoding/json"
	"net/http"
	"snakeish/golang/http_utils"
)

func GetRoomsEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

	response := map[string]RoomPreviewStruct{}
	for key, room := range Manager.Rooms {
		response[key] = RoomPreviewStruct{
			Id:           room.Id,
			Name:         room.RoomName,
			PlayersCount: len(room.Users),
			GameModeName: "classic", //TODO change gamemodename if you add gamemodes
		}
	}

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
