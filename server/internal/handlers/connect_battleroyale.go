package handlers

import (
	"snakeish/internal/services/controllers/battleroyale_controller"
	"snakeish/pkg/core"
	"snakeish/pkg/sockets"

	"github.com/gin-gonic/gin"
)

func ConnectBattleroyaleEndpoint(c *gin.Context) {
	roomId := c.Params.ByName("id")

	room, found := core.GetRoomById(roomId)
	if !found {
		c.JSON(404, gin.H{
			"code":    "ROOM_NOT_FOUND",
			"message": "couldn't find room with given id",
		})
		return
	}

	if room.GetTagName() != "battle-royale" {
		c.JSON(403, gin.H{
			"code":    "ROOM_WRONG_MODE_TAG",
			"message": "found room, but it's mode tag is wrong",
		})
		return
	}

	websocket, err := sockets.CreateClient(c.Writer, c.Request)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "WEBSOCKET_ERROR",
			"message": "errro while creating websocket client",
		})
		return
	}

	battleroyale_controller.ConnectWebsocket(room, websocket)
}
