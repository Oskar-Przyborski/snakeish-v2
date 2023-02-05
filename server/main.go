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

	api.GET("rooms", GetRoomsEndpoint)
	api.GET("rooms/suggested", GetSuggestedRoomsEndpoint)
	api.GET("rooms/name/:name", RoomNameAvailableEndpoint)
	api.POST("rooms/create", CreateRoomEndpoint)

	api.GET("room/:id", GetRoomEndpoint)

	api.GET("connect/classic/:id", ConnectClassicRoomEndpoint)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("Listening on port: " + port)
	router.Run(":" + port)
}
