package main

import (
	"fmt"
	"net/http"
)

var Manager = CreateAppManager()

func main() {
	http.HandleFunc("/api/create-classic-room", CreateClassicRoomEndpoint)
	http.HandleFunc("/api/get-room", GetRoomEndpoint)
	http.HandleFunc("/api/get-rooms", GetRoomsEndpoint)
	http.HandleFunc("/api/ws-connect-room", WsConnectRoomEndpoint)

	http.HandleFunc("/api/check", func(w http.ResponseWriter, r *http.Request) {
		msg := ""

		msg += "Game Rooms\n"
		for id, room := range Manager.Rooms {
			msg += id + "\n"
			msg += "	Players\n"
			for idx, player := range room.Users {
				msg += fmt.Sprintf("	%d. %s \n", idx, player.Websocket.Id)
			}
		}

		w.Write([]byte(msg))
	})

	println("Listening on localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		println(err.Error())
	}
}
