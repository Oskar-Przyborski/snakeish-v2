package battleroyaleMode

import "snakeish/pkg/core/utils"

type Player struct {
	Id        string          `json:"id"`
	Name      string          `json:"name"`
	Color     string          `json:"color"`
	SnakeTail []utils.Vector2 `json:"snakeTail"`
	Score     int             `json:"score"`
	Direction utils.Vector2   `json:"direction"`
	Alive     bool            `json:"alive"`
}

type GameUpdateResponse struct {
	FrameTime    int             `json:"frameTime"`
	GridSize     int             `json:"gridSize"`
	ShrinkSize   int             `json:"shrinkSize"`
	Apples       []utils.Vector2 `json:"apples"`
	Players      []Player        `json:"players"`
	GameStatus   string          `json:"gameStatus"`
	Winner       *Player         `json:"winner"`
	StartUnix    int64           `json:"startUnix"`
	UnfreezeUnix int64           `json:"unfreezeUnix"`
	MinPlayers   int             `json:"minPlayers"`
	MaxPlayers   int             `json:"maxPlayers"`
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
