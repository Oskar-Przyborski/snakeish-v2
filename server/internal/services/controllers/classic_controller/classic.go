package classic_controller

import (
	"snakeish/internal/services/clients"
	"snakeish/pkg/core"
	"snakeish/pkg/core/rooms"
	"snakeish/pkg/core/rooms/classic"
	"snakeish/pkg/core/utils"
	"snakeish/pkg/sockets"
	"time"
)

func SubscribeUpdates(room *rooms.Room) {
	room.OnUpdate.AddListener(onUpdate)
}

func onUpdate(room *rooms.Room) {
	response := generateGameUpdateResponse(*room)
	clients := clients.GetClientsFromRoom(room.Id)

	for _, client := range clients {
		client.WebSocket.Send("game-update", response)
	}
}

func generateGameUpdateResponse(room rooms.Room) GameUpdateResponse {
	mode := room.Mode.(*classic.Mode)
	response := GameUpdateResponse{
		FrameTime: mode.GetFrameTime(),
		GridSize:  mode.GridSize,
		Apples:    mode.Apples,
		Players:   []Player{},
	}

	for _, player := range mode.Players {
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

func ConnectWebsocket(room *rooms.Room, socket *sockets.SocketClient) {
	core.StopAfkForRoom(room.Id)

	client := clients.CreateClient(socket, *room)
	client.OnDisconnect.AddListener(onClientDisconnect)

	client.OnEvent("request-join", onJoinRequest)
	client.OnEvent("request-leave", onLeaveRequest)
	client.OnEvent("change-direction", onChangeDirection)

	go socket.ListenForMessages()
}

func onClientDisconnect(client *clients.Client) {
	client.Room.RemovePlayer(client.PlayerId)

	if client.Room.GetPlayersCount() == 0 {
		core.StartAfkForRoom(client.Room.Id, 30*time.Second)
	}
	clients.RemoveClient(client.WebSocket.Id)
}

func onJoinRequest(client *clients.Client, c sockets.MessageContext) {
	data := JoinRequestType{}
	if err := c.BindJSON(&data); err != nil {
		return
	}

	mode := client.Room.Mode.(*classic.Mode)
	player, err := mode.AddPlayer(data.Name, data.Color, data.Pin)
	if err != nil {
		client.WebSocket.Send("join-error", map[string]any{
			"error": err.Error(),
		})
		return
	}

	client.IsPlayer = true
	client.PlayerId = player.Id

	client.WebSocket.Send("join-success", JoinSuccesType{
		PlayerId: player.Id,
		Name:     player.Name,
		Color:    player.Color,
	})

	println("JOIN_CLASSIC_ROOM:", client.Room.Name, "with name:", player.Name)
}

func onLeaveRequest(client *clients.Client, c sockets.MessageContext) {
	client.Room.RemovePlayer(client.PlayerId)
	client.IsPlayer = false
	client.PlayerId = ""
}

func onChangeDirection(client *clients.Client, c sockets.MessageContext) {
	mode := client.Room.Mode.(*classic.Mode)
	player := mode.GetPlayerById(client.PlayerId)

	payload := ChangeDirectionRequest{}
	if err := c.BindJSON(&payload); err != nil {
		println("Error while parsing json:", err.Error())
		return
	}

	direction, err := utils.DirectionToVector(payload.Direction)
	if err != nil {
		return
	}

	player.ChangeDirection(direction)
}
