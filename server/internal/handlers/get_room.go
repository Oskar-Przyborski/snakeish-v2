package handlers

import (
	"snakeish/pkg/core"

	"github.com/gin-gonic/gin"
)

func GetRoomEndpoint(c *gin.Context) {
	id := c.Params.ByName("id")

	room, _ := core.GetRoomById(id)
	if room == nil {
		c.JSON(404, gin.H{
			"code":    "ROOM_NOT_FOUND",
			"message": "Couldn't find room with given id",
		})
		return
	}

	c.JSON(200, core.GetRoomPreview(room))
}
