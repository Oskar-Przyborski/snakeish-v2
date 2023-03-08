package battleroyale

import "snakeish/pkg/core/room"

func Configure(base room.Base) *Room {
	switch base.ModeName {
	default:
		return &Room{
			Base:           base,
			MaxPlayers:     8,
			MinPlayers:     2, //TODO change it to 4
			GameStatus:     "waiting-for-players",
			FrameTime:      250,
			ApplesQuantity: 6,
			GridSize:       20,
			ShrinkSize:     0,
		}
	}
}
