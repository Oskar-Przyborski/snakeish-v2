package battleroyaleMode

import (
	"snakeish/internal/services/clients"
	"snakeish/pkg/core"
	"snakeish/pkg/core/room/battleroyale"
	"snakeish/pkg/sockets"
)

func CreateRoom(roomName string, modeName string, pin *[4]int) (*battleroyale.Room, error) {
	print("creating room")
	room, err := core.CreateBattleRoyaleRoom(roomName, modeName, pin)
	if err != nil {
		return nil, err
	}

	print("adding listener")
	room.OnUpdate.AddListener(onUpdate)

	return room, nil
}

func onUpdate(room *battleroyale.Room) {
	response := generateGameUpdateResponse(*room)
	clients := clients.GetClientsFromRoom(room.Id)

	for _, client := range clients {
		client.WebSocket.Send("game-update", response)
	}
}

func generateGameUpdateResponse(room battleroyale.Room) GameUpdateResponse {
	response := GameUpdateResponse{
		FrameTime: room.FrameTime,
		GridSize:  5,
		// Apples:    room.Apples,
		Players: []Player{},
	}

	for _, player := range room.Players {
		response.Players = append(response.Players, Player{
			Id:        player.Id,
			Name:      player.Name,
			Color:     player.Color,
			SnakeTail: player.SnakeTail,
			Score:     len(player.SnakeTail),
			Direction: player.Direction,
		})
	}

	return response
}

func ConnectWebsocket(room *battleroyale.Room, socket *sockets.SocketClient) {
	core.StopAfkForRoom(room.Id)

	clients.CreateClient(socket, room)
	// client.OnDisconnect.AddListener(onClientDisconnect)

	// client.OnEvent("request-join", onJoinRequest)
	// client.OnEvent("request-leave", onLeaveRequest)
	// client.OnEvent("change-direction", onChangeDirection)

	go socket.ListenForMessages()

	println("Connected websocket. Client id:", socket.Id)
}
