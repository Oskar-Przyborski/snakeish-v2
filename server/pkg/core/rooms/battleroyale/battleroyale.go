package battleroyale

import (
	"snakeish/pkg/core/utils"
	"snakeish/pkg/notifier"
	"time"
)

type Mode struct {
	ModeName           string
	Players            []*Player
	MaxPlayers         int
	MinPlayers         int
	Apples             []utils.Vector2
	AppleSpawnCounter  int
	MaxApplesQuantity  int
	FrameTime          int
	GridSize           int
	ShrinkSize         int
	ShrinkFrameCounter int
	OnUpdate           notifier.Notifier[*Mode]
	GameStatus         string  // "waiting-for-players", "playing", "finished"
	Winner             *Player // The winner of the game. Nil if game not ended.
	StartUnix          int64   // Unix timestamp of game start moment. -1 if null.
	StartTimer         *time.Timer
	Freezed            bool  // After start of the game, game will be freezed for a couple seconds.
	UnfreezeUnix       int64 // Unix timestamp of the moment of unfreeze.
	ZoneKillTime       int
	ZoneShrinkTime     int
	AppleSpawnTime     int
}

func (mode Mode) GetMaxPlayers() int {
	return mode.MaxPlayers
}

func (mode Mode) GetTagName() string {
	return "battle-royale"
}
func (mode Mode) GetModeName() string {
	return mode.ModeName
}

func (mode Mode) GetPlayersCount() int {
	return len(mode.Players)
}

func (mode Mode) GetFrameTime() int {
	return mode.FrameTime
}
