package classic_room

import "snakeish/core/utils"

type ClassicPlayer struct {
	Id              string
	Name            string
	Color           string
	SnakeTail       []utils.Vector2
	IsAlive         bool
	Direction       utils.Vector2
	TargetDirection utils.Vector2
}

func (player *ClassicPlayer) ChangeDirection(direcion utils.Vector2) {
	if player.Direction.X+direcion.X == 0 && player.Direction.Y+direcion.Y == 0 {
		return
	}

	player.TargetDirection = direcion
}

func (player *ClassicPlayer) Kill() {
	player.IsAlive = false
	player.SnakeTail = make([]utils.Vector2, 0)
}

func (player ClassicPlayer) IsCollidingSelf() bool {
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

func (player ClassicPlayer) IsOutOfBounds(gridSize int) bool {
	if len(player.SnakeTail) == 0 {
		return false
	}
	head := player.SnakeTail[0]

	if head.X < 0 || head.Y < 0 || head.X >= gridSize || head.Y >= gridSize {
		return true
	}
	return false
}

func (player ClassicPlayer) IsCollidingWith(other ClassicPlayer) bool {
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
