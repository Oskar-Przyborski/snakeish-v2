package classic

import "snakeish/pkg/core/utils"

type Player struct {
	Id        string          `json:"id"`
	Name      string          `json:"name"`
	Color     string          `json:"color"`
	SnakeTail []utils.Vector2 `json:"snakeTail"`
	Score     int             `json:"score"`
	Direction utils.Vector2   `json:"direction"`
}

type GameUpdateResponse struct {
	FrameTime int             `json:"frameTime"`
	GridSize  int             `json:"gridSize"`
	Apples    []utils.Vector2 `json:"apples"`
	Players   []Player        `json:"players"`
}

type JoinRequestType struct {
	Color string `json:"color"`
	Name  string `json:"name"`
	Pin   [4]int `json:"pin"`
}

type JoinSuccesType struct {
	PlayerId string `json:"playerId"`
	Name     string `json:"name"`
	Color    string `json:"color"`
}

type ChangeDirectionRequest struct {
	Direction string `json:"direction"`
}
