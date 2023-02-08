package main

import (
	"snakeish/core/room"

	"github.com/gin-gonic/gin"
)

func GetRoomsEndpoint(c *gin.Context) {
	response := []room.RoomPreview{}
	for _, room := range Core.GetRooms() {
		response = append(response, room.GetPreview())
	}

	c.JSON(200, response)
}
