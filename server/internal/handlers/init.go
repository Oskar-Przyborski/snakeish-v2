package handlers

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("rooms", GetRoomsEndpoint)
	api.GET("rooms/suggested", GetSuggestedRoomsEndpoint)
	api.GET("rooms/name/:name", RoomNameAvailableEndpoint)
	api.POST("rooms/create", CreateRoomEndpoint)

	api.GET("room/:id", GetRoomEndpoint)

	api.GET("connect/classic/:id", ConnectClassicRoomEndpoint)
	api.GET("connect/battle-royale/:id", ConnectBattleroyaleEndpoint)
}
