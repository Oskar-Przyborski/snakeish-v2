package battleroyale

import "snakeish/pkg/core/room"

type Room struct {
	room.Base
	Players    []*Player
	MaxPlayers int
	GameStatus string
	Winner     *Player
}
