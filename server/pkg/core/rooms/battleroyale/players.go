package battleroyale

import (
	"errors"
	"snakeish/pkg/core/player"
	"snakeish/pkg/core/utils"

	"github.com/google/uuid"
	"github.com/thoas/go-funk"
)

func (room *Mode) AddPlayer(name string, color string, pin [4]int) (*Player, error) {
	// if !room.CheckPin(pin) {
	// 	return nil, errors.New("incorrect-pin")
	// }

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
			Id:              uuid.NewString(),
			SnakeTail:       []utils.Vector2{},
			TargetDirection: utils.Vector2{X: 1, Y: 0},
		},
		Name:    name,
		Color:   color,
		IsAlive: false,
	}

	room.Players = append(room.Players, &player)

	return &player, nil
}

func (room *Mode) MovePlayer(player *Player) {
	if len(player.SnakeTail) == 0 {
		return
	}

	player.Direction = player.TargetDirection

	newHeadPos := utils.Vector2{
		X: player.SnakeTail[0].X + player.Direction.X,
		Y: player.SnakeTail[0].Y + player.Direction.Y,
	}

	if !room.EatAppleAt(newHeadPos) {
		player.SnakeTail = append([]utils.Vector2{newHeadPos}, player.SnakeTail[:len(player.SnakeTail)-1]...)
	} else {
		player.SnakeTail = append([]utils.Vector2{newHeadPos}, player.SnakeTail...)
	}
}

func (room *Mode) RemovePlayer(id string) {
	for idx, player := range room.Players {
		if player.Id == id {
			room.Players = append(room.Players[:idx], room.Players[idx+1:]...)
			break
		}
	}
}

func (room *Mode) IsAnySnakeAtPosition(at utils.Vector2) bool {
	for _, player := range room.Players {
		for _, snakeElement := range player.SnakeTail {
			if snakeElement.IsEqual(at) {
				return true
			}
		}
	}
	return false
}

func (room *Mode) IsPlayerCollidingWithAnyOther(player Player) bool {
	for _, searchUser := range room.Players {
		if searchUser.Id == player.Id {
			continue
		}
		if player.IsCollidingWith(searchUser.Base) {
			return true
		}
	}

	return false
}

func (room Mode) SpawnAllPlayers() {
	for _, player := range room.Players {
		room.SpawnPlayer(player)
	}
}

func (room Mode) SpawnPlayer(player *Player) {
	position, err := room.GetRandomFreePosition(1)
	if err != nil {
		return
	}
	player.SnakeTail = []utils.Vector2{position}
	player.IsAlive = true
}

func (room *Mode) GetAlivePlayers() []*Player {
	return funk.Filter(room.Players, func(player *Player) bool {
		return player.IsAlive
	}).([]*Player)
}
