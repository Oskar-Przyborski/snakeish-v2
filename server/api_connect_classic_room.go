package main

import (
	"encoding/json"
	"net/http"
	classic_room "snakeish/core/room/classic"
	"snakeish/core/utils"
	"snakeish/gosockets"
	"snakeish/http_utils"
)

func ConnectClassicRoomEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

	roomId := r.URL.Query().Get("id")
	if roomId == "" {
		http_utils.WriteError(&w, 400, "missing-parameter", "url should contain 'id' parameter")
		return
	}

	room, found := Core.GetRoomById(roomId)
	if !found {
		http_utils.WriteError(&w, 400, "room-not-found", "room with given id doesn't exist")
		return
	}

	if room.GetModeTag() != "classic" {
		http_utils.WriteError(&w, 400, "room-wrong-mode-tag", "room should be classic")
		return
	}

	classicRoom := room.(*classic_room.ClassicRoom)

	websocket, err := gosockets.CreateClient(w, r)
	if err != nil {
		http_utils.WriteError(&w, 500, "server-error", "error while creating websocket client")
		return
	}

	connectedClient := ClientsManager.CreateConnectedClient(websocket, room)
	connectedClient.WebSocket.OnDisconnect = append(connectedClient.WebSocket.OnDisconnect, func() {
		classicRoom.RemovePlayer(connectedClient.PlayerId)
		ClientsManager.RemoveConnectedClient(websocket.Id)

		println("Disconnected client with id: " + websocket.Id)
	})

	type joinRequestType struct {
		Color string `json:"color"`
		Name  string `json:"name"`
	}

	connectedClient.WebSocket.AddListener("request-join", func(s string) {
		data := joinRequestType{}
		if err := json.Unmarshal([]byte(s), &data); err != nil {
			return
		}

		player := classicRoom.AddPlayer(data.Name, data.Color)

		connectedClient.IsPlayer = true
		connectedClient.PlayerId = player.Id
	})

	connectedClient.WebSocket.AddListener("request-leave", func(s string) {
		classicRoom.RemovePlayer(connectedClient.PlayerId)
		connectedClient.IsPlayer = false
		connectedClient.PlayerId = ""
	})

	connectedClient.WebSocket.AddListener("change-direction", func(s string) {
		player := classicRoom.GetPlayerById(connectedClient.PlayerId)
		type ChangeDirectionPayload struct {
			Direction string `json:"direction"`
		}
		payload := ChangeDirectionPayload{}
		if err := json.Unmarshal([]byte(s), &payload); err != nil {
			println("Error while parsing json:", err.Error())
			return
		}
		direction, err := utils.DirectionToVector(payload.Direction)
		if err != nil {
			return
		}
		player.ChangeDirection(direction)
	})

	go websocket.ListenForMessages()

	println("Connected websocket. Client id:", websocket.Id)
}
