package main

import (
	"fmt"
	"net/http"
	"os"
	"snakeish/core"
)

var Manager = CreateAppManager()
var Core = core.CreateCore()

func main() {
	http.HandleFunc("/api/create-room", CreateRoomEndpoint)
	http.HandleFunc("/api/get-room", GetRoomEndpoint)
	http.HandleFunc("/api/get-suggested-rooms", GetSuggestedRoomsEndpoint)
	http.HandleFunc("/api/get-rooms", GetRoomsEndpoint)
	http.HandleFunc("/api/room-name-available", RoomNameAvailableEndpoint)
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	println("Listening on port: " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		println(err.Error())
	}
}
