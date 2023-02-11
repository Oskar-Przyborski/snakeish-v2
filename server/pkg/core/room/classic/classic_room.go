package classic_room

import (
	"errors"
	"snakeish/pkg/core/room"
	"snakeish/pkg/core/utils"
	"time"

	"github.com/google/uuid"
)

type ClassicRoom struct {
	room.RoomBase
	Apples         []utils.Vector2
	Players        []*ClassicPlayer
	MaxPlayers     int
	FrameTime      int
	GridSize       int
	ApplesQuantity int
	ModeName       string
}

var allowedPlayersColors = []string{
	"#deb135",
	"#80e356",
	"#e37e56",
	"#56e37e",
	"#56dee3",
	"#5672e3",
	"#8a56e3",
	"#e356bd",
	"#e3566b",
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

func (croom *ClassicRoom) GetPreview() room.RoomPreview {
	return room.RoomPreview{
		Id:         croom.Id,
		Name:       croom.Name,
		ModeTag:    croom.GetModeTag(),
		ModeName:   croom.GetModeName(),
		Players:    len(croom.Players),
		MaxPlayers: croom.MaxPlayers,
		PinEnbled:  croom.PinEnabled,
	}
}

func (room *ClassicRoom) StartRoom() {
	room.SpawnMissingApples()
	for {
		time.Sleep(time.Duration(room.FrameTime) * time.Millisecond)
		room.Update()
	}
}

func ConfigureClassicRoom(base room.RoomBase, mode string) *ClassicRoom {
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

func (room *ClassicRoom) AddPlayer(name string, color string, pin [4]int) (*ClassicPlayer, error) {
	if !room.CheckPin(pin) {
		return nil, errors.New("incorrect-pin")
	}

	if len(name) < 3 || len(name) > 10 {
		return nil, errors.New("player-name-incorrect-length")
	}

	if !room.IsPlayerNameAvailable(name) {
		return nil, errors.New("player-name-already-taken")
	}
	if !utils.ArrayIncludes(allowedPlayersColors, color) {
		return nil, errors.New("color-not-allowed")
	}

	player := ClassicPlayer{
		Id:              uuid.NewString(),
		Name:            name,
		Color:           color,
		SnakeTail:       []utils.Vector2{},
		IsAlive:         false,
		TargetDirection: utils.Vector2{X: 1, Y: 0},
	}

	room.Players = append(room.Players, &player)
	return &player, nil
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

func (room *ClassicRoom) IsPlayerNameAvailable(name string) bool {
	for _, player := range room.Players {
		if player.Name == name {
			return false
		}
	}

	return true
}
