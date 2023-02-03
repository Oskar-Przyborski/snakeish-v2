package main

import (
	"encoding/json"
	"math"
	"net/http"
	"snakeish/http_utils"
	"sort"
)

type roomEvaluation struct {
	value float64
	room  *Room
}

type resposnseStruct struct {
	Rooms          []RoomPreviewStruct `json:"rooms"`
	RemainingRooms int                 `json:"remainingRooms"`
}

func GetSuggestedRoomsEndpoint(w http.ResponseWriter, r *http.Request) {
	if ended := http_utils.CheckCors(&w, r); ended {
		return
	}

	evaluations := []roomEvaluation{}

	for _, room := range Manager.Rooms {
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

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}

// TODO implement better algorithim for suggesting rooms
func evaluate(room *Room) roomEvaluation {
	players := float64(room.getPlayersCount())
	max := float64(room.MaxPlayers)
	var score float64 = 0

	halfMax := max / 2
	playersScore := (halfMax - math.Abs(players-halfMax)) / halfMax
	score += playersScore * 5

	return roomEvaluation{
		value: score,
		room:  room,
	}
}
