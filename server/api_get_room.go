package main

import (
	"github.com/gin-gonic/gin"
)

type RoomPreviewStruct struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Players      int    `json:"players"`
	MaxPlayers   int    `json:"maxPlayers"`
	GameModeName string `json:"gameModeName"`
	GameModeTag  string `json:"gameModeTag"`
}

func GetRoomEndpoint(c *gin.Context) {
	id, found := c.GetQuery("id")
	if !found {
		c.JSON(400, gin.H{
			"code":    "MISSING_PARAMETER",
			"message": "query should contain 'id' parameter",
		})
		return
	}

	room, _ := Core.GetRoomById(id)
	if room == nil {
		c.JSON(404, gin.H{
			"code":    "ROOM_NOT_FOUND",
			"message": "Couldn't find room with given id",
		})
		return
	}

	c.JSON(200, room.GetPreview())
}
