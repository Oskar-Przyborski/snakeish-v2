package main

import (
	"errors"
	"snakeish/golang/gosockets"
)

type User struct {
	Websocket *gosockets.GosocketClient
	IsPlayer  bool
	player    Player
}

type Player struct {
	Color           string    `json:"color"`
	Name            string    `json:"name"`
	SnakeTail       []Vector2 `json:"snakeTail"`
	Direction       Vector2
	TargetDirection Vector2
	IsAlive         bool
}

type Vector2 struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (vector Vector2) IsEqual(other Vector2) bool {
	return vector.X == other.X && vector.Y == other.Y
}

func directionToVector(direction string) (Vector2, error) {
	switch direction {
	case "up":
		return Vector2{X: 0, Y: -1}, nil
	case "down":
		return Vector2{X: 0, Y: 1}, nil
	case "right":
		return Vector2{X: 1, Y: 0}, nil
	case "left":
		return Vector2{X: -1, Y: 0}, nil
	default:
		return Vector2{}, errors.New("can not convert direction to vector2")
	}
}

func (player *Player) Kill() {
	player.IsAlive = false
	player.SnakeTail = make([]Vector2, 0)
}
