package classic

import (
	"errors"
	"snakeish/pkg/core/room"
	"snakeish/pkg/core/utils"
	"snakeish/pkg/notifier"
	"time"

	"github.com/google/uuid"
)

type Room struct {
	room.Base
	Apples         []utils.Vector2
	Players        []*Player
	MaxPlayers     int
	FrameTime      int
	GridSize       int
	ApplesQuantity int
	ModeName       string
	OnUpdate       notifier.Notifier[*Room]
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

func (room Room) GetModeTag() string {
	return "classic"
}

func (room Room) GetMaxPlayers() int {
	return room.MaxPlayers
}
func (room Room) GetPlayersCount() int {
	return len(room.Players)
}

func (croom *Room) GetPreview() room.RoomPreview {
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

func (room *Room) StartRoom() {
	room.SpawnMissingApples()
	for {
		time.Sleep(time.Duration(room.FrameTime) * time.Millisecond)
		room.Update()
	}
}

func ConfigureClassicRoom(base room.Base) *Room {
	switch base.ModeName {
	default:
		return &Room{
			Base:           base,
			ApplesQuantity: 3,
			GridSize:       8,
			MaxPlayers:     4,
			FrameTime:      250,
			ModeName:       "Casual",
		}
	case "Huuge":
		return &Room{
			Base:           base,
			ApplesQuantity: 8,
			GridSize:       16,
			MaxPlayers:     8,
			FrameTime:      250,
			ModeName:       "Huuge",
		}

	}
}

func (room *Room) AddPlayer(name string, color string, pin [4]int) (*Player, error) {
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

	player := Player{
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

func (room *Room) RemovePlayer(id string) {
	for idx, player := range room.Players {
		if player.Id == id {
			room.Players = append(room.Players[:idx], room.Players[idx+1:]...)
			break
		}
	}
}

func (room *Room) GetPlayerById(id string) *Player {
	for _, player := range room.Players {
		if player.Id == id {
			return player
		}
	}
	return nil
}

func (room *Room) IsPlayerNameAvailable(name string) bool {
	for _, player := range room.Players {
		if player.Name == name {
			return false
		}
	}

	return true
}
