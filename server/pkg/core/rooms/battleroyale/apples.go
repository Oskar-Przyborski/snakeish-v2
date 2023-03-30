package battleroyale

import "snakeish/pkg/core/utils"

func (room *Mode) HandleAppleSpawn() {
	room.AppleSpawnCounter++

	if room.AppleSpawnCounter > room.AppleSpawnTime {
		room.SpawnApple()
		room.AppleSpawnCounter = 0
	}
}

func (room *Mode) SpawnApple() {
	position, err := room.GetRandomFreePosition(room.ShrinkSize)
	if err != nil {
		return
	}

	room.Apples = append(room.Apples, position)
}

func (room Mode) IsAppleAtPosition(at utils.Vector2) bool {
	for _, apple := range room.Apples {
		if apple.IsEqual(at) {
			return true
		}
	}
	return false
}

func (room *Mode) EatAppleAt(at utils.Vector2) bool {
	for idx, apple := range room.Apples {
		if apple.IsEqual(at) {
			room.Apples = append(room.Apples[:idx], room.Apples[idx+1:]...)
			return true
		}
	}

	return false
}
