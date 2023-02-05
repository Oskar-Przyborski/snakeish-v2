package main

import (
	"github.com/gin-gonic/gin"
)

func RoomNameAvailableEndpoint(c *gin.Context) {
	name, found := c.GetQuery("name")
	if !found {
		c.JSON(400, gin.H{
			"code":    "MISSING_PARAMETER",
			"message": "query should contain 'name' parameter",
		})
		return
	}

	roomNameAvailable := true
	for _, room := range Core.GetRooms() {
		if room.GetName() == name {
			roomNameAvailable = false
			break
		}
	}

	c.JSON(200, gin.H{"available": roomNameAvailable})
}
