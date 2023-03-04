package battleroyale

import (
	"snakeish/pkg/core/room"
	"snakeish/pkg/core/utils"
	"snakeish/pkg/notifier"
)

type Room struct {
	room.Base
	Players        []*Player
	MaxPlayers     int
	MinPlayers     int
	Apples         []utils.Vector2
	ApplesQuantity int
	FrameTime      int
	GridSize       int
	ShrinkSize     int
	OnUpdate       notifier.Notifier[*Room]
	GameStatus     string  // "waiting-for-players", "starting", "playing", "finished"
	Winner         *Player // the winner of the game. Nil if game not ended.
	StartUnix      int64   // Unix timestamp of game start moment. -1 if null
}

func (room Room) GetMaxPlayers() int {
	return room.MaxPlayers
}

func (room Room) GetModeTag() string {
	return "battle-royale"
}

func (room Room) GetPlayersCount() int {
	return len(room.Players)
}
