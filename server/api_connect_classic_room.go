package main

import (
	"net/http"
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

	websocket, err := gosockets.CreateClient(w, r)
	if err != nil {
		http_utils.WriteError(&w, 500, "server-error", "error while creating websocket client")
		return
	}

	connectedClient := ClientsManager.CreateConnectedClient(websocket, room)
	connectedClient.WebSocket.OnDisconnect = append(connectedClient.WebSocket.OnDisconnect, func() {
		ClientsManager.RemoveConnectedClient(websocket.Id)
		println("Disconnected client with id: " + websocket.Id)
	})

	go websocket.ListenForMessages()

	println("Connected websocket. Client id:", websocket.Id)
}
