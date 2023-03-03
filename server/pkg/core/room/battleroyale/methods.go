package battleroyale

import (
	"errors"
	"snakeish/pkg/core/player"
	"snakeish/pkg/core/utils"
	"time"

	"github.com/google/uuid"
	"github.com/thoas/go-funk"
)

func (room *Room) StartRoom() {
	for {
		time.Sleep(time.Duration(room.FrameTime) * time.Millisecond)
		room.Update()
	}
}

func (room *Room) Update() {
	room.OnUpdate.Notify(room)
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

	if !funk.ContainsString(PlayerColors, color) {
		return nil, errors.New("color-not-allowed")
	}

	player := Player{
		Base: player.Base{
			Id: uuid.NewString(),
			SnakeTail: []utils.Vector2{
				{X: 0, Y: 0},
				{X: 0, Y: 1},
			},
			TargetDirection: utils.Vector2{X: 1, Y: 0},
		},
		Name:    name,
		Color:   color,
		IsAlive: false,
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
