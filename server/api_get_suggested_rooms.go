package main

import (
	"math"
	"snakeish/core/room"
	"sort"

	"github.com/gin-gonic/gin"
)

type roomEvaluation struct {
	value float64
	room  room.IRoom
}

type resposnseStruct struct {
	Rooms          []room.RoomPreview `json:"rooms"`
	RemainingRooms int                `json:"remainingRooms"`
}

func GetSuggestedRoomsEndpoint(c *gin.Context) {

	evaluations := []roomEvaluation{}

	for _, room := range Core.GetRooms() {
		evaluations = append(evaluations, evaluate(room))
	}

	sort.SliceStable(evaluations, func(i, j int) bool {
		return evaluations[i].value > evaluations[j].value
	})

	response := resposnseStruct{}
	for i := 0; i < len(evaluations); i++ {
		if i >= 4 {
			break
		}
		response.Rooms = append(response.Rooms, evaluations[i].room.GetPreview())
	}
	if rooms := len(evaluations); rooms > 4 {
		response.RemainingRooms = rooms - 4
	}

	c.JSON(200, response)
}

// TODO implement better algorithim for suggesting rooms
func evaluate(room room.IRoom) roomEvaluation {
	players := float64(room.GetPlayersCount())
	max := float64(room.GetMaxPlayers())
	var score float64 = 0

	halfMax := max / 2
	playersScore := (halfMax - math.Abs(players-halfMax)) / halfMax
	score += playersScore * 5

	return roomEvaluation{
		value: score,
		room:  room,
	}
}
