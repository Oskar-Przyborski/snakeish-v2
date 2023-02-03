package main

import (
	"math"
	"net/http"
	"snakeish/core/room"
	"snakeish/http_utils"
	"sort"
)

type roomEvaluation struct {
	value float64
	room  room.IRoom
}

type resposnseStruct struct {
	Rooms          []room.RoomPreview `json:"rooms"`
	RemainingRooms int                `json:"remainingRooms"`
}

func GetSuggestedRoomsEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

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

	http_utils.WriteJSON(w, response)
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
