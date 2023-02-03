package main

import (
	"net/http"
	"snakeish/core/room"
	"snakeish/http_utils"
)

func GetRoomsEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

	response := []room.RoomPreview{}
	for _, room := range Core.GetRooms() {
		response = append(response, room.GetPreview())
	}

	http_utils.WriteJSON(w, response)
}
