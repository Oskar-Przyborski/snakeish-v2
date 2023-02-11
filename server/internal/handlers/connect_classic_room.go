package handlers

import (
	"snakeish/internal/services"
	"snakeish/pkg/core"
	classic_room "snakeish/pkg/core/room/classic"
	"snakeish/pkg/core/utils"
	"snakeish/pkg/sockets"
	"time"

	"github.com/gin-gonic/gin"
)

func ConnectClassicRoomEndpoint(c *gin.Context) {
	roomId := c.Params.ByName("id")

	room, found := core.Instance.GetRoomById(roomId)
	if !found {
		c.JSON(404, gin.H{
			"code":    "ROOM_NOT_FOUND",
			"message": "couldn't find room with given id",
		})
		return
	}

	if room.GetModeTag() != "classic" {
		c.JSON(403, gin.H{
			"code":    "ROOM_WRONG_MODE_TAG",
			"message": "found room, but it's mode tag is wrong",
		})
		return
	}

	classicRoom := room.(*classic_room.ClassicRoom)

	websocket, err := sockets.CreateClient(c.Writer, c.Request)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "WEBSOCKET_ERROR",
			"message": "errro while creating websocket client",
		})
		return
	}

	core.Instance.StopAfkForRoom(roomId)
	connectedClient := services.ClientsManager.CreateConnectedClient(websocket, room)
	connectedClient.WebSocket.OnDisconnect = append(connectedClient.WebSocket.OnDisconnect, func() {
		classicRoom.RemovePlayer(connectedClient.PlayerId)
		services.ClientsManager.RemoveConnectedClient(websocket.Id)

		if classicRoom.GetPlayersCount() == 0 {
			core.Instance.StartAfkForRoom(roomId, 30*time.Second)
		}
		println("Disconnected client with id: " + websocket.Id)
	})

	connectedClient.WebSocket.AddListener("request-join", func(c sockets.MessageContext) {
		type requestType struct {
			Color string `json:"color"`
			Name  string `json:"name"`
			Pin   [4]int `json:"pin"`
		}
		data := requestType{}
		if err := c.BindJSON(&data); err != nil {
			return
		}

		player, err := classicRoom.AddPlayer(data.Name, data.Color, data.Pin)
		if err != nil {
			connectedClient.WebSocket.Send("join-error", map[string]any{
				"error": err.Error(),
			})
			return
		}

		connectedClient.IsPlayer = true
		connectedClient.PlayerId = player.Id

		type responseType struct {
			PlayerId string `json:"playerId"`
			Name     string `json:"name"`
			Color    string `json:"color"`
		}
		connectedClient.WebSocket.Send("join-success", responseType{
			PlayerId: player.Id,
			Name:     player.Name,
			Color:    player.Color,
		})
	})

	connectedClient.WebSocket.AddListener("request-leave", func(c sockets.MessageContext) {
		classicRoom.RemovePlayer(connectedClient.PlayerId)
		connectedClient.IsPlayer = false
		connectedClient.PlayerId = ""
	})

	connectedClient.WebSocket.AddListener("change-direction", func(c sockets.MessageContext) {
		player := classicRoom.GetPlayerById(connectedClient.PlayerId)

		type ChangeDirectionPayload struct {
			Direction string `json:"direction"`
		}
		payload := ChangeDirectionPayload{}
		if err := c.BindJSON(&payload); err != nil {
			println("Error while parsing json:", err.Error())
			return
		}

		direction, err := utils.DirectionToVector(payload.Direction)
		if err != nil {
			return
		}

		player.ChangeDirection(direction)
	})

	go websocket.ListenForMessages()

	println("Connected websocket. Client id:", websocket.Id)
}
