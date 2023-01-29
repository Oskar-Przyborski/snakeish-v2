package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"snakeish/golang/http_utils"

	"github.com/google/uuid"
)

type CreateClassicRoomEndpointData struct {
	RoomName   string `json:"roomName"`
	ConfigName string `json:"configName"`
}

func CreateClassicRoomEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

	println("Create room request")
	if r.Header["Content-Type"] == nil || r.Header["Content-Type"][0] != "application/json" {
		http_utils.WriteError(&w, 400, "invalid-content-type", "header 'Content-Type' should be 'application/json'")
		return
	}

	var reqData CreateClassicRoomEndpointData
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http_utils.WriteError(&w, 400, "body-incomplete", "error while parsing request body\nrequest body should contain fields:\nroomName <string>\nconfigName: <string>")
		return
	}

	for _, room := range Manager.Rooms {
		if room.RoomName == reqData.RoomName {
			http_utils.WriteError(&w, 400, "name-exists", fmt.Sprintf("room with name '%s' already exists", reqData.RoomName))
			return
		}
	}

	room := Room{
		RoomName: reqData.RoomName,
		Id:       "rm-" + uuid.NewString(),
		Users:    []*User{},
	}

	room.SetConfig(reqData.ConfigName)

	Manager.Rooms[room.Id] = &room
	go room.StartRoom()

	println("Created room. Id:", room.Id)
	w.WriteHeader(200)
}

func (room *Room) SetConfig(configName string) {
	switch configName {
	default:
		room.ApplesQuantity = 3
		room.FrameTime = 250
		room.GridSize = 9
		room.ModeTag = "classic"
		room.ModeName = "Casual"
		room.MaxPlayers = 5
	}
}
