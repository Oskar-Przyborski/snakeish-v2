package battleroyale

import (
	"snakeish/pkg/core/room"
)

func Configure(base room.Base) *Room {
	switch base.ModeName {
	default:
		return &Room{
			Base:              base,
			MaxPlayers:        8,
			MinPlayers:        3,
			GameStatus:        "waiting-for-players",
			FrameTime:         250,
			MaxApplesQuantity: 6,
			GridSize:          20,
			ShrinkSize:        0,
			ZoneShrinkTime:    15,
			ZoneKillTime:      3,
		}
	}
}
