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
				MaxPlayers:     8,
				MinPlayers:     3,
				FrameTime:      250,
				GridSize:       20,
				ShrinkSize:     0,
				ZoneShrinkTime: 15,
				ZoneKillTime:   3,
				AppleSpawnTime: 5,
				ModeName:       "Eat&Win",
			}
		case "Zone Shark":
			return &battleroyale.Mode{
				MaxPlayers:     8,
				MinPlayers:     3,
				FrameTime:      200,
				GridSize:       25,
				ShrinkSize:     0,
				ZoneShrinkTime: 10,
				ZoneKillTime:   0,
				AppleSpawnTime: 4,
				ModeName:       "Zone Shark",
			}
		}
	}
}
