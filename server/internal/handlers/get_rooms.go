package handlers

import (
	"snakeish/pkg/core"
	"snakeish/pkg/core/room"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

func GetRoomsEndpoint(c *gin.Context) {
	rooms := core.GetRooms()

	//filtering
	filterOnlyPublic(&rooms, c.Query("public") == "1")

	response := funk.Map(rooms, func(room room.IRoom) room.RoomPreview {
		return room.GetPreview()
	})

	c.JSON(200, response)
}

func filterOnlyPublic(rooms *[]room.IRoom, value bool) {
	if value {
		(*rooms) = (funk.Filter(*rooms, func(room room.IRoom) bool {
			return !room.IsPinEnabled()
		}).([]room.IRoom))
	}
}
