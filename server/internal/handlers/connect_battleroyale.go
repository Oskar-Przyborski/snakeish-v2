package handlers

import (
	battleroyaleMode "snakeish/internal/services/modes/battleroyale"
	"snakeish/pkg/core"
	"snakeish/pkg/core/room/battleroyale"
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

	if room.GetModeTag() != "battle-royale" {
		c.JSON(403, gin.H{
			"code":    "ROOM_WRONG_MODE_TAG",
			"message": "found room, but it's mode tag is wrong",
		})
		return
	}

	battleroyaleRoom := room.(*battleroyale.Room)

	websocket, err := sockets.CreateClient(c.Writer, c.Request)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "WEBSOCKET_ERROR",
			"message": "errro while creating websocket client",
		})
		return
	}

	battleroyaleMode.ConnectWebsocket(battleroyaleRoom, websocket)
}
