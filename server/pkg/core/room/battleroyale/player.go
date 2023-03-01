package battleroyale

import (
	"snakeish/pkg/core/player"
)

type Player struct {
	player.Base
	Name    string
	Color   string
	IsAlive bool
}
