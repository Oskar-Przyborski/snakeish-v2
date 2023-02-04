package main

import (
	classic_room "snakeish/core/room/classic"
	"snakeish/core/utils"
)

func CreateClassicRoom(roomName string, modeName string) (*classic_room.ClassicRoom, error) {
	room, err := Core.CreateClassicRoom(roomName, modeName)
	if err != nil {
		return &classic_room.ClassicRoom{}, err
	}

	room.OnUpdate(func() {
		response := generateGameUpdateResponse(*room)
		clients := ClientsManager.GetClientsFromRoom(room.Id)

		for _, client := range clients {
			client.WebSocket.Send("game-update", response)
		}
	})

	return room, nil
}

type ClassicRoomPlayerResponse struct {
	Id        string          `json:"id"`
	Name      string          `json:"name"`
	Color     string          `json:"color"`
	SnakeTail []utils.Vector2 `json:"snakeTail"`
	Score     int             `json:"score"`
}

type ClassicRoomGameUpdateResponse struct {
	GridSize int                         `json:"gridSize"`
	Apples   []utils.Vector2             `json:"apples"`
	Players  []ClassicRoomPlayerResponse `json:"players"`
}

func generateGameUpdateResponse(room classic_room.ClassicRoom) ClassicRoomGameUpdateResponse {
	response := ClassicRoomGameUpdateResponse{
		GridSize: room.GridSize,
		Apples:   room.Apples,
		Players:  []ClassicRoomPlayerResponse{},
	}

	for _, player := range room.Players {
		response.Players = append(response.Players, ClassicRoomPlayerResponse{
			Id:        player.Id,
			Name:      player.Name,
			Color:     player.Color,
			SnakeTail: player.SnakeTail,
			Score:     len(player.SnakeTail),
		})
	}

	return response
}
