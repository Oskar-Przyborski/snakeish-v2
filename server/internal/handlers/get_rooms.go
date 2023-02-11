package handlers

import (
	"snakeish/pkg/core"
	"snakeish/pkg/core/room"

	"github.com/gin-gonic/gin"
)

func GetRoomsEndpoint(c *gin.Context) {
	response := []room.RoomPreview{}
	for _, room := range core.Instance.GetRooms() {
		response = append(response, room.GetPreview())
	}

	c.JSON(200, response)
}
