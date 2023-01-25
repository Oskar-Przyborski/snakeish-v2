package main

import (
	"net/http"
	"snakeish/golang/gosockets"
	"snakeish/golang/http_utils"
)

func WsConnectRoomEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

	roomId := r.URL.Query().Get("id")
	if roomId == "" {
		http_utils.WriteError(&w, 400, "missing-parameter", "url should contain 'id' parameter")
		return
	}

	room := Manager.GetRoomById(roomId)
	if room == nil {
		http_utils.WriteError(&w, 400, "room-not-found", "room with given id doesn't exist")
		return
	}

	wsClient, err := gosockets.CreateClient(w, r)
	if err != nil {
		http_utils.WriteError(&w, 500, "server-error", "error while creating websocket client")
		return
	}

	room.AddUser(wsClient)

	go wsClient.ListenForMessages()

	println("Connected websocket. Client id:", wsClient.Id)
}
