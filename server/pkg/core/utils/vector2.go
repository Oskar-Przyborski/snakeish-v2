package utils

import "errors"

type Vector2 struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (vector Vector2) IsEqual(other Vector2) bool {
	return vector.X == other.X && vector.Y == other.Y
}

func DirectionToVector(direction string) (Vector2, error) {
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
