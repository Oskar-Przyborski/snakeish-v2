package player

import "snakeish/pkg/core/utils"

type Base struct {
	Id              string
	Direction       utils.Vector2
	TargetDirection utils.Vector2
	SnakeTail       []utils.Vector2
}

func (player *Base) ChangeDirection(direcion utils.Vector2) {
	if player.Direction.X+direcion.X == 0 && player.Direction.Y+direcion.Y == 0 {
		return
	}

	player.TargetDirection = direcion
}

func (player Base) IsCollidingSelf() bool {
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

func (player Base) IsOutOfBounds(gridSize int) bool {
	if len(player.SnakeTail) == 0 {
		return false
	}

	head := player.SnakeTail[0]

	if head.X < 0 || head.Y < 0 || head.X >= gridSize || head.Y >= gridSize {
		return true
	}
	return false
}

func (player Base) IsCollidingWith(other Base) bool {
	if player.Id == other.Id {
		return false
	}

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
