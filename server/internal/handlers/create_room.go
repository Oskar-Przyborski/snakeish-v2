package handlers

import (
	"snakeish/internal/services/controllers/battleroyale_controller"
	"snakeish/internal/services/controllers/classic_controller"
	"snakeish/pkg/core"
	"snakeish/pkg/core/rooms"

	"github.com/gin-gonic/gin"
)

type CreateRoomEndpointData struct {
	RoomName string  `json:"roomName" binding:"required"`
	ModeTag  string  `json:"modeTag" binding:"required"`
	ModeName string  `json:"modeName" binding:"required"`
	Pin      *[4]int `json:"pin"`
}

func CreateRoomEndpoint(c *gin.Context) {
	var data CreateRoomEndpointData
	if err := c.BindJSON(&data); err != nil {
		return
	}

	for _, room := range core.GetRooms() {
		if room.Name == data.RoomName {
			c.JSON(400, gin.H{
				"code":    "NAME_EXISTS",
				"message": "room with given name already exists",
			})
			return
		}
	}

	room, err := core.CreateRoom(data.RoomName, data.ModeTag, data.ModeName, data.Pin)
	SubscribeToRoom(room)

	if err != nil {
		c.JSON(500, gin.H{
			"code":    "CREATION_ERROR",
			"message": "error while creating room",
		})
	}

	c.JSON(200, core.GetRoomPreview(*room))
	println("Created room. Id:", room.Id)
}

func SubscribeToRoom(room *rooms.Room) {
	switch room.GetTagName() {
	default:
		classic_controller.SubscribeUpdates(room)
	case "battle-royale":
		battleroyale_controller.SubscribeUpdates(room)
	}
}
