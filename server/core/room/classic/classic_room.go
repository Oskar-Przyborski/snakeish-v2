package classic_room

import (
	r "snakeish/core/room"
	"snakeish/core/utils"
)

type ClassicRoom struct {
	r.RoomBase
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
func (room ClassicRoom) GetMaxPlayers() int {
	return room.MaxPlayers
}
func (room ClassicRoom) GetPlayersCount() int {
	return len(room.Players)
}

func (room ClassicRoom) GetPreview() r.RoomPreview {
	return r.RoomPreview{
		Id:         room.Id,
		Name:       room.Name,
		ModeTag:    room.GetModeTag(),
		ModeName:   room.GetModeName(),
		Players:    len(room.Players),
		MaxPlayers: room.MaxPlayers,
	}
}
