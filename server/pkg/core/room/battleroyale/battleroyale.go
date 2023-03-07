package battleroyale

import (
	"snakeish/pkg/core/room"
	"snakeish/pkg/core/utils"
	"snakeish/pkg/notifier"
	"time"
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
	GameStatus     string  // "waiting-for-players", "playing", "finished"
	Winner         *Player // The winner of the game. Nil if game not ended.
	StartUnix      int64   // Unix timestamp of game start moment. -1 if null.
	StartTimer     *time.Timer
	Freezed        bool  // After start of the game, game will be freezed for a couple seconds.
	UnfreezeUnix   int64 // Unix timestamp of the moment of unfreeze.
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
