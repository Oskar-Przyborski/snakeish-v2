package main

import (
	"snakeish/core/room"
	"time"

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

	for _, room := range Core.GetRooms() {
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
	Core.StartAfkForRoom(room.GetId(), 30*time.Second)

	c.JSON(200, room.GetPreview())
	println("Created room. Id:", room.GetId())
}

func CreateRoomByModeTag(name string, modeTag string, modeName string, pin *[4]int) (room.IRoom, error) {
	switch modeTag {
	default: //classic
		return CreateClassicRoom(name, modeName, pin)
	}
}
