package rooms

import (
	"snakeish/pkg/core/rooms/battleroyale"
	"snakeish/pkg/core/rooms/classic"
)

func CreateMode(tag string, name string) Mode {
	switch tag {
	default: //classic
		switch name {
		default: //casual
			return &classic.Mode{
				ApplesQuantity: 3,
				GridSize:       8,
				MaxPlayers:     4,
				FrameTime:      250,
				ModeName:       "Casual",
			}
		case "Huuge":
			return &classic.Mode{
				ApplesQuantity: 8,
				GridSize:       16,
				MaxPlayers:     8,
				FrameTime:      250,
				ModeName:       "Huuge",
			}
		}

	case "battle-royale":
		switch name {
		default: //Eat&Win
			return &battleroyale.Mode{
				MaxPlayers:        8,
				MinPlayers:        3,
				GameStatus:        "waiting-for-players",
				FrameTime:         250,
				MaxApplesQuantity: 6,
				GridSize:          20,
				ShrinkSize:        0,
				ZoneShrinkTime:    15,
				ZoneKillTime:      3,
				ModeName:          "Eat&Win",
			}
		}
	}
}
