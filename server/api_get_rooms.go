package main

import (
	"net/http"
	"snakeish/http_utils"
)

func GetRoomsEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

	response := map[string]RoomPreviewStruct{}
	for key, room := range Manager.Rooms {
		response[key] = room.GetPreview()
	}

	http_utils.WriteJSON(response, w)
}
