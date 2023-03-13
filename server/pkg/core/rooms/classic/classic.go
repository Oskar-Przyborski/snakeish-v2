package classic

import (
	"errors"
	"snakeish/pkg/core/player"
	"snakeish/pkg/core/utils"

	"github.com/google/uuid"
	"github.com/thoas/go-funk"
)

type Mode struct {
	Apples         []utils.Vector2
	Players        []*Player
	MaxPlayers     int
	FrameTime      int
	GridSize       int
	ApplesQuantity int
	ModeName       string
}

func (mode Mode) GetFrameTime() int {
	return mode.FrameTime
}

func (mode Mode) GetTagName() string {
	return "classic"
}
func (mode Mode) GetModeName() string {
	return mode.ModeName
}

func (mode Mode) GetMaxPlayers() int {
	return mode.MaxPlayers
}
func (mode Mode) GetPlayersCount() int {
	return len(mode.Players)
}

func (mode *Mode) Init() {
	mode.SpawnMissingApples()
}

func (mode *Mode) AddPlayer(name string, color string, pin [4]int) (*Player, error) {
	// if !mode.CheckPin(pin) {
	// 	return nil, errors.New("incorrect-pin")
	// }

	if len(name) < 3 || len(name) > 10 {
		return nil, errors.New("player-name-incorrect-length")
	}

	if !mode.IsPlayerNameAvailable(name) {
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

	mode.Players = append(mode.Players, &player)
	return &player, nil
}

func (mode *Mode) RemovePlayer(id string) {
	for idx, player := range mode.Players {
		if player.Id == id {
			mode.Players = append(mode.Players[:idx], mode.Players[idx+1:]...)
			break
		}
	}
}

func (mode *Mode) GetPlayerById(id string) *Player {
	for _, player := range mode.Players {
		if player.Id == id {
			return player
		}
	}
	return nil
}

func (mode *Mode) IsPlayerNameAvailable(name string) bool {
	for _, player := range mode.Players {
		if player.Name == name {
			return false
		}
	}

	return true
}
