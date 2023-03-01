package battleroyale

import (
	"snakeish/pkg/core/room"
	"snakeish/pkg/notifier"
)

type Room struct {
	room.Base
	Players    []*Player
	MaxPlayers int
	MinPlayers int
	GameStatus string
	Winner     *Player
	OnUpdate   notifier.Notifier[*Room]
	FrameTime  int
}

func (room Room) GetMaxPlayers() int {
	return room.MaxPlayers
}

func (room Room) GetModeTag() string {
	return "battle royale"
}

func (room Room) GetPlayersCount() int {
	return len(room.Players)
}
