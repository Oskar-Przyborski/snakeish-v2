package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"os"
	"snakeish/core"
)

var ClientsManager = CreateConnectedClientsManager()
var Core = core.CreateCore()

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	api := router.Group("api")

	api.GET("get-rooms", GetRoomsEndpoint)
	api.GET("get-suggested-rooms", GetSuggestedRoomsEndpoint)
	api.GET("get-room", GetRoomEndpoint)
	api.GET("room-name-available", RoomNameAvailableEndpoint)
	api.POST("create-room", CreateRoomEndpoint)
	api.GET("ws-connect-classic-room", ConnectClassicRoomEndpoint)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("Listening on port: " + port)
	router.Run(":" + port)
}
