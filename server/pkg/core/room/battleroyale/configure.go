package battleroyale

import "snakeish/pkg/core/room"

func Configure(base room.Base) *Room {
	switch base.ModeName {
	default:
		return &Room{
			Base:       base,
			MaxPlayers: 8,
			MinPlayers: 4,
			GameStatus: "starting",
			FrameTime:  250,
			GridSize:   8,
			ShrinkSize: 0,
		}
	}
}
