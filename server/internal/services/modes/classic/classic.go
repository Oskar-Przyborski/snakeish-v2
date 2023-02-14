package classic

import (
	"snakeish/internal/services/clients"
	"snakeish/pkg/core"
	classic_room "snakeish/pkg/core/room/classic"
	"snakeish/pkg/core/utils"
	"snakeish/pkg/sockets"
	"time"
)

func CreateRoom(roomName string, modeName string, pin *[4]int) (*classic_room.ClassicRoom, error) {
	room, err := core.CreateClassicRoom(roomName, modeName, pin)
	if err != nil {
		return &classic_room.ClassicRoom{}, err
	}

	room.OnUpdate.AddListener(onUpdate)

	return room, nil
}

func onUpdate(room *classic_room.ClassicRoom) {
	response := generateGameUpdateResponse(*room)
	clients := clients.GetClientsFromRoom(room.Id)

	for _, client := range clients {
		client.WebSocket.Send("game-update", response)
	}
}

func generateGameUpdateResponse(room classic_room.ClassicRoom) GameUpdateResponse {
	response := GameUpdateResponse{
		FrameTime: room.FrameTime,
		GridSize:  room.GridSize,
		Apples:    room.Apples,
		Players:   []Player{},
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

func ConnectWebsocket(room *classic_room.ClassicRoom, socket *sockets.SocketClient) {
	core.StopAfkForRoom(room.Id)

	client := clients.CreateClient(socket, room)
	client.OnDisconnect.AddListener(onClientDisconnect)

	client.OnEvent("request-join", onJoinRequest)
	client.OnEvent("request-leave", onLeaveRequest)
	client.OnEvent("change-direction", onChangeDirection)

	go socket.ListenForMessages()

	println("Connected websocket. Client id:", socket.Id)
}

func onClientDisconnect(client *clients.Client) {
	client.Room.RemovePlayer(client.PlayerId)

	if client.Room.GetPlayersCount() == 0 {
		core.StartAfkForRoom(client.Room.GetId(), 30*time.Second)
	}
	clients.RemoveClient(client.WebSocket.Id)

	println("Disconnected client with id: " + client.WebSocket.Id)
}

func onJoinRequest(client *clients.Client, c sockets.MessageContext) {
	data := JoinRequestType{}
	if err := c.BindJSON(&data); err != nil {
		return
	}

	player, err := client.Room.(*classic_room.ClassicRoom).AddPlayer(data.Name, data.Color, data.Pin)
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
}

func onLeaveRequest(client *clients.Client, c sockets.MessageContext) {
	client.Room.RemovePlayer(client.PlayerId)
	client.IsPlayer = false
	client.PlayerId = ""
}

func onChangeDirection(client *clients.Client, c sockets.MessageContext) {
	player := client.Room.(*classic_room.ClassicRoom).GetPlayerById(client.PlayerId)

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
