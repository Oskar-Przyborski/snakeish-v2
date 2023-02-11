package classic

import "snakeish/pkg/core/utils"

type Player struct {
	Id        string          `json:"id"`
	Name      string          `json:"name"`
	Color     string          `json:"color"`
	SnakeTail []utils.Vector2 `json:"snakeTail"`
	Score     int             `json:"score"`
}

type GameUpdateResponse struct {
	GridSize int             `json:"gridSize"`
	Apples   []utils.Vector2 `json:"apples"`
	Players  []Player        `json:"players"`
}
