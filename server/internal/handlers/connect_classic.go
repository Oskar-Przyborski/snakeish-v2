package handlers

import (
	classicMode "snakeish/internal/services/modes/classic"
	"snakeish/pkg/core"
	"snakeish/pkg/core/room/classic"
	"snakeish/pkg/sockets"

	"github.com/gin-gonic/gin"
)

func ConnectClassicRoomEndpoint(c *gin.Context) {
	roomId := c.Params.ByName("id")

	room, found := core.GetRoomById(roomId)
	if !found {
		c.JSON(404, gin.H{
			"code":    "ROOM_NOT_FOUND",
			"message": "couldn't find room with given id",
		})
		return
	}

	if room.GetModeTag() != "classic" {
		c.JSON(403, gin.H{
			"code":    "ROOM_WRONG_MODE_TAG",
			"message": "found room, but it's mode tag is wrong",
		})
		return
	}

	classicRoom := room.(*classic.Room)

	websocket, err := sockets.CreateClient(c.Writer, c.Request)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "WEBSOCKET_ERROR",
			"message": "errro while creating websocket client",
		})
		return
	}

	classicMode.ConnectWebsocket(classicRoom, websocket)
}
