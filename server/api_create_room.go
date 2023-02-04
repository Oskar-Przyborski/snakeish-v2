package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"snakeish/http_utils"
)

type CreateRoomEndpointData struct {
	RoomName   string `json:"roomName"`
	ConfigName string `json:"configName"`
}

// TODO add PIN code support
func CreateRoomEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

	println("Create room request")
	if r.Header["Content-Type"] == nil || r.Header["Content-Type"][0] != "application/json" {
		http_utils.WriteError(&w, 400, "invalid-content-type", "header 'Content-Type' should be 'application/json'")
		return
	}

	var reqData CreateRoomEndpointData
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http_utils.WriteError(&w, 400, "body-incomplete", "error while parsing request body\nrequest body should contain fields:\nroomName <string>\nconfigName: <string>")
		return
	}

	for _, room := range Core.GetRooms() {
		if room.GetName() == reqData.RoomName {
			http_utils.WriteError(&w, 400, "name-exists", fmt.Sprintf("room with name '%s' already exists", reqData.RoomName))
			return
		}
	}

	room, err := Core.CreateRoom(reqData.RoomName, reqData.ConfigName)
	if err != nil {
		http_utils.WriteError(&w, 500, "room-create-error", "error while creating room")
	}

	println("Created room. Id:", room.GetId())
	http_utils.WriteJSON(w, room.GetPreview())
}
