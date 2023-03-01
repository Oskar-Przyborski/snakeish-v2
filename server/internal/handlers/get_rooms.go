package handlers

import (
	"regexp"
	"snakeish/pkg/core"
	r "snakeish/pkg/core/room"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

func GetRoomsEndpoint(c *gin.Context) {
	rooms := core.GetRooms()

	//filtering
	filterOnlyPublic(&rooms, c.Query("public") == "1")
	filterSearch(&rooms, c.Query("s"))
	filterModes(&rooms, c.Query("modes"))

	response := funk.Map(rooms, func(room r.Room) r.RoomPreview {
		return core.GetRoomPreview(room)
	})

	c.JSON(200, response)
}

func filterOnlyPublic(rooms *[]r.Room, value bool) {
	if value {
		(*rooms) = (funk.Filter(*rooms, func(room r.Room) bool {
			return !room.IsPinEnabled()
		}).([]r.Room))
	}
}

func filterSearch(rooms *[]r.Room, query string) {
	if query == "" {
		return
	}

	rgx, err := regexp.Compile("(?i).*" + query + ".*")
	if err != nil {
		return
	}

	(*rooms) = funk.Filter((*rooms), func(room r.Room) bool {
		return rgx.MatchString(room.GetName())
	}).([]r.Room)
}

func filterModes(rooms *[]r.Room, modesString string) {
	if modesString == "" {
		return
	}

	modes := strings.Split(modesString, ",")

	if len(modes) == 0 {
		return
	}

	(*rooms) = funk.Filter((*rooms), func(room r.Room) bool {
		for _, mode := range modes {
			if room.GetModeName() == mode {
				return true
			}
		}
		return false
	}).([]r.Room)
}
