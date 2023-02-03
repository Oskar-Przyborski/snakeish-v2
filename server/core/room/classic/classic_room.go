package classic_room

import (
	"snakeish/core/room"
	"snakeish/core/utils"
)

type ClassicRoom struct {
	room.RoomBase
	Apples     []utils.Vector2
	Players    []ClassicPlayer
	MaxPlayers int
}

func (room ClassicRoom) GetModeTag() string {
	return "CLASSIC"
}
func (room ClassicRoom) GetModeName() string {
	return "Casual"
}
