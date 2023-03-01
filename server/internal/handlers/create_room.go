package handlers

import (
	classicMode "snakeish/internal/services/modes/classic"
	"snakeish/pkg/core"
	"snakeish/pkg/core/room"

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
		if room.GetName() == data.RoomName {
			c.JSON(400, gin.H{
				"code":    "NAME_EXISTS",
				"message": "room with given name already exists",
			})
			return
		}
	}

	room, err := CreateRoomByModeTag(data.RoomName, data.ModeTag, data.ModeName, data.Pin)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "CREATION_ERROR",
			"message": "error while creating room",
		})
	}

	c.JSON(200, core.GetRoomPreview(room))
	println("Created room. Id:", room.GetId())
}

func CreateRoomByModeTag(name string, modeTag string, modeName string, pin *[4]int) (room.Room, error) {
	switch modeTag {
	default: //classic
		return classicMode.CreateRoom(name, modeName, pin)
	case "battle royale":
		return core.CreateBattleRoyaleRoom(name, modeName, pin)
	}
}
