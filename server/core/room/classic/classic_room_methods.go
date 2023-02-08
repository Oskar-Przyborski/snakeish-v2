package classic_room

import (
	"errors"
	"math/rand"
	"snakeish/core/utils"
)

func (room *ClassicRoom) Update() {
	playersToKill := []*ClassicPlayer{}

	for _, player := range room.Players {
		if !player.IsAlive {
			continue
		}
		room.MovePlayer(player)
	}

	for _, player := range room.Players {
		if !player.IsAlive {
			continue
		}

		if player.IsOutOfBounds(room.GridSize) || player.IsCollidingSelf() || room.IsPlayerCollidingWithAnyOther(*player) {
			playersToKill = append(playersToKill, player)
			continue
		}
	}

	for _, player := range playersToKill {
		player.Kill()
	}

	for _, player := range room.Players {
		if player.IsAlive {
			continue
		}
		room.RespawnPlayer(player)
	}

	room.SpawnMissingApples()
	if room.Listener != nil {
		room.Listener()
	}
}

func (room *ClassicRoom) SpawnMissingApples() {
	applesToSpawn := room.ApplesQuantity - len(room.Apples)

	for i := 0; i < applesToSpawn; i++ {
		room.SpawnApple()
	}
}

func (room *ClassicRoom) SpawnApple() {
	position, err := room.GetRandomFreePosition(0)
	if err != nil {
		return
	}

	room.Apples = append(room.Apples, position)
}

func (room *ClassicRoom) GetRandomFreePosition(distanceFromBounds int) (utils.Vector2, error) {
	freePositions := []utils.Vector2{}
	for y := distanceFromBounds; y < room.GridSize-distanceFromBounds; y++ {
		for x := distanceFromBounds; x < room.GridSize-distanceFromBounds; x++ {
			position := utils.Vector2{Y: y, X: x}
			if room.IsPositionEmpty(position) {
				freePositions = append(freePositions, position)
			}
		}
	}
	if len(freePositions) == 0 {
		return utils.Vector2{}, errors.New("there is no free position")
	}

	return freePositions[rand.Intn(len(freePositions))], nil
}

func (room *ClassicRoom) IsPositionEmpty(at utils.Vector2) bool {
	if room.IsAnySnakeAtPosition(at) || room.IsAppleAtPosition(at) {
		return false
	}
	return true
}

func (room *ClassicRoom) IsAnySnakeAtPosition(at utils.Vector2) bool {
	for _, player := range room.Players {
		for _, snakeElement := range player.SnakeTail {
			if snakeElement.IsEqual(at) {
				return true
			}
		}
	}
	return false
}

func (room ClassicRoom) IsAppleAtPosition(at utils.Vector2) bool {
	for _, apple := range room.Apples {
		if apple.IsEqual(at) {
			return true
		}
	}
	return false
}

func (room *ClassicRoom) MovePlayer(player *ClassicPlayer) {
	if len(player.SnakeTail) == 0 {
		return
	}

	player.Direction = player.TargetDirection

	newHeadPos := utils.Vector2{
		X: player.SnakeTail[0].X + player.Direction.X,
		Y: player.SnakeTail[0].Y + player.Direction.Y,
	}

	if !room.EatAppleAt(newHeadPos) {
		player.SnakeTail = append([]utils.Vector2{newHeadPos}, player.SnakeTail[:len(player.SnakeTail)-1]...)
	} else {
		player.SnakeTail = append([]utils.Vector2{newHeadPos}, player.SnakeTail...)
	}
}

func (room *ClassicRoom) EatAppleAt(at utils.Vector2) bool {
	for idx, apple := range room.Apples {
		if apple.IsEqual(at) {
			room.Apples = append(room.Apples[:idx], room.Apples[idx+1:]...)
			return true
		}
	}

	return false
}

func (room *ClassicRoom) IsPlayerCollidingWithAnyOther(player ClassicPlayer) bool {
	for _, searchUser := range room.Players {
		if searchUser.Id == player.Id {
			continue
		}
		if player.IsCollidingWith(*searchUser) {
			return true
		}
	}

	return false
}

func (room ClassicRoom) RespawnPlayer(player *ClassicPlayer) {
	position, err := room.GetRandomFreePosition(1)
	if err != nil {
		return
	}
	player.SnakeTail = []utils.Vector2{position}
	player.IsAlive = true
}
