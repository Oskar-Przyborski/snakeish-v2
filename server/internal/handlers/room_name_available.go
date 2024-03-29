package handlers

import (
	"snakeish/pkg/core"

	"github.com/gin-gonic/gin"
)

func RoomNameAvailableEndpoint(c *gin.Context) {
	name := c.Params.ByName("name")

	roomNameAvailable := true
	for _, room := range core.GetRooms() {
		if room.Name == name {
			roomNameAvailable = false
			break
		}
	}

	c.JSON(200, gin.H{"available": roomNameAvailable})
}
