package classic_room

import (
	r "snakeish/core/room"
	"snakeish/core/utils"
	"time"

	"github.com/google/uuid"
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
	return "classic"
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

// TODO add name validation
func (room *ClassicRoom) AddPlayer(name string, color string) *ClassicPlayer {
	player := ClassicPlayer{
		Id:              uuid.NewString(),
		Name:            name,
		Color:           color,
		SnakeTail:       []utils.Vector2{},
		IsAlive:         false,
		TargetDirection: utils.Vector2{X: 1, Y: 0},
	}

	room.Players = append(room.Players, &player)
	return &player
}

func (room *ClassicRoom) RemovePlayer(id string) {
	for idx, player := range room.Players {
		if player.Id == id {
			room.Players = append(room.Players[:idx], room.Players[idx+1:]...)
			break
		}
	}
}

func (room *ClassicRoom) GetPlayerById(id string) *ClassicPlayer {
	for _, player := range room.Players {
		if player.Id == id {
			return player
		}
	}
	return nil
}
