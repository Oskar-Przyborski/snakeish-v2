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

func (player *Player) IsCollidingSelf() bool {
	if len(player.SnakeTail) <= 1 {
		return false
	}

	head := player.SnakeTail[0]
	for idx, tailElement := range player.SnakeTail {
		if idx == 0 {
			continue //do not check head
		}
		if head.IsEqual(tailElement) {
			return true
		}
	}

	return false
}

func (player *Player) IsOutOfBounds(gridSize int) bool {
	if len(player.SnakeTail) == 0 {
		return false
	}
	head := player.SnakeTail[0]

	if head.X < 0 || head.Y < 0 || head.X >= gridSize || head.Y >= gridSize {
		return true
	}
	return false
}

func (player Player) IsCollidingWith(other Player) bool {
	if len(player.SnakeTail) == 0 || len(other.SnakeTail) == 0 {
		return false
	}

	playerHead := player.SnakeTail[0]
	for _, tailElem := range other.SnakeTail {
		if playerHead.IsEqual(tailElem) {
			return true
		}
	}
	return false
}
