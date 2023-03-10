package classic

import (
	"snakeish/pkg/core/player"
	"snakeish/pkg/core/utils"
)

type Player struct {
	player.Base
	Name    string
	Color   string
	IsAlive bool
}

func (player *Player) Kill() {
	player.IsAlive = false
	player.SnakeTail = make([]utils.Vector2, 0)
}
