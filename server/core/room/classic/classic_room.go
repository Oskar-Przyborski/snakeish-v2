package classic_room

import (
	r "snakeish/core/room"
	"snakeish/core/utils"
	"time"
)

type ClassicRoom struct {
	r.RoomBase
	Apples         []utils.Vector2
	Players        []*ClassicPlayer
	MaxPlayers     int
	FrameTime      int
	GridSize       int
	ApplesQuantity int
	ModeName       string
}

func (room ClassicRoom) GetModeTag() string {
	return "CLASSIC"
}
func (room ClassicRoom) GetModeName() string {
	return room.ModeName
}
func (room ClassicRoom) GetMaxPlayers() int {
	return room.MaxPlayers
}
func (room ClassicRoom) GetPlayersCount() int {
	return len(room.Players)
}

func (room *ClassicRoom) GetPreview() r.RoomPreview {
	return r.RoomPreview{
		Id:         room.Id,
		Name:       room.Name,
		ModeTag:    room.GetModeTag(),
		ModeName:   room.GetModeName(),
		Players:    len(room.Players),
		MaxPlayers: room.MaxPlayers,
	}
}

func (room *ClassicRoom) StartRoom() {
	room.SpawnMissingApples()
	for {
		time.Sleep(time.Duration(room.FrameTime) * time.Millisecond)
		room.Update()
	}
}

func CreateClassicRoom(base r.RoomBase, mode string) *ClassicRoom {
	switch mode {
	default:
		return &ClassicRoom{
			RoomBase:       base,
			ApplesQuantity: 3,
			GridSize:       8,
			MaxPlayers:     4,
			FrameTime:      250,
			ModeName:       "Casual",
		}
	case "Huuge":
		return &ClassicRoom{
			RoomBase:       base,
			ApplesQuantity: 8,
			GridSize:       16,
			MaxPlayers:     8,
			FrameTime:      250,
			ModeName:       "Huuge",
		}

	}
}
