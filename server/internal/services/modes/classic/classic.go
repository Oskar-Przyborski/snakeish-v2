package classic

import (
	"snakeish/internal/services/clients_manager"
	"snakeish/pkg/core"
	classic_room "snakeish/pkg/core/room/classic"
)

func CreateRoom(roomName string, modeName string, pin *[4]int) (*classic_room.ClassicRoom, error) {
	room, err := core.Instance.CreateClassicRoom(roomName, modeName, pin)
	if err != nil {
		return &classic_room.ClassicRoom{}, err
	}

	room.OnUpdate(func() {
		response := generateGameUpdateResponse(*room)
		clients := clients_manager.GetClientsFromRoom(room.Id)

		for _, client := range clients {
			client.WebSocket.Send("game-update", response)
		}
	})

	return room, nil
}

func generateGameUpdateResponse(room classic_room.ClassicRoom) GameUpdateResponse {
	response := GameUpdateResponse{
		GridSize: room.GridSize,
		Apples:   room.Apples,
		Players:  []Player{},
	}

	for _, player := range room.Players {
		response.Players = append(response.Players, Player{
			Id:        player.Id,
			Name:      player.Name,
			Color:     player.Color,
			SnakeTail: player.SnakeTail,
			Score:     len(player.SnakeTail),
		})
	}

	return response
}
