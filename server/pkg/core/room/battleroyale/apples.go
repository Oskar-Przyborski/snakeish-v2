package battleroyale

import "snakeish/pkg/core/utils"

func (room *Room) SpawnMissingApples() {
	applesToSpawn := room.ApplesQuantity - len(room.Apples)

	for i := 0; i < applesToSpawn; i++ {
		room.SpawnApple()
	}
}

func (room *Room) SpawnApple() {
	position, err := room.GetRandomFreePosition(0)
	if err != nil {
		return
	}

	room.Apples = append(room.Apples, position)
}

func (room Room) IsAppleAtPosition(at utils.Vector2) bool {
	for _, apple := range room.Apples {
		if apple.IsEqual(at) {
			return true
		}
	}
	return false
}

func (room *Room) EatAppleAt(at utils.Vector2) bool {
	for idx, apple := range room.Apples {
		if apple.IsEqual(at) {
			room.Apples = append(room.Apples[:idx], room.Apples[idx+1:]...)
			return true
		}
	}

	return false
}
