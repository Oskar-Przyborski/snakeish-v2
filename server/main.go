package main

import (
	"net/http"
	"os"
	"snakeish/core"
)

var ClientsManager = CreateConnectedClientsManager()
var Core = core.CreateCore()

func main() {
	http.HandleFunc("/api/create-room", CreateRoomEndpoint)
	http.HandleFunc("/api/get-room", GetRoomEndpoint)
	http.HandleFunc("/api/get-suggested-rooms", GetSuggestedRoomsEndpoint)
	http.HandleFunc("/api/get-rooms", GetRoomsEndpoint)
	http.HandleFunc("/api/room-name-available", RoomNameAvailableEndpoint)
	http.HandleFunc("/api/ws-connect-classic-room", ConnectClassicRoomEndpoint)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	println("Listening on port: " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		println(err.Error())
	}
}
