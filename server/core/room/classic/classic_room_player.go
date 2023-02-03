package classic_room

import "snakeish/core/utils"

type ClassicPlayer struct {
	Id        string
	Name      string
	Color     string
	SnakeTail []utils.Vector2
}
