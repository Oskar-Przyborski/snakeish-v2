package classic

import "snakeish/pkg/core/room"

func ConfigureClassicRoom(base room.Base) *Room {
	switch base.ModeName {
	default:
		return &Room{
			Base:           base,
			ApplesQuantity: 3,
			GridSize:       8,
			MaxPlayers:     4,
			FrameTime:      250,
			ModeName:       "Casual",
		}
	case "Huuge":
		return &Room{
			Base:           base,
			ApplesQuantity: 8,
			GridSize:       16,
			MaxPlayers:     8,
			FrameTime:      250,
			ModeName:       "Huuge",
		}

	}
}
