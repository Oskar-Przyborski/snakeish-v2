package battleroyale

import (
	"errors"
	"math/rand"
	"snakeish/pkg/core/utils"
)

func (room *Room) GetPlayerById(id string) *Player {
	for _, player := range room.Players {
		if player.Id == id {
			return player
		}
	}
	return nil
}

func (room *Room) IsPlayerNameAvailable(name string) bool {
	for _, player := range room.Players {
		if player.Name == name {
			return false
		}
	}

	return true
}

func (room *Room) IsPositionEmpty(at utils.Vector2) bool {
	if room.IsAnySnakeAtPosition(at) || room.IsAppleAtPosition(at) {
		return false
	}
	return true
}

func (room Room) IsInsideZone(at utils.Vector2) bool {
	if at.X < room.ShrinkSize || at.X > room.GridSize-room.ShrinkSize {
		return true
	}

	if at.Y < room.ShrinkSize || at.Y > room.GridSize-room.ShrinkSize {
		return true
	}

	return false
}

func (room *Room) GetRandomFreePosition(distanceFromBounds int) (utils.Vector2, error) {
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
