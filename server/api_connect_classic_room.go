package main

import (
	classic_room "snakeish/core/room/classic"
	"snakeish/core/utils"
	"snakeish/gosockets"

	"github.com/gin-gonic/gin"
)

func ConnectClassicRoomEndpoint(c *gin.Context) {
	roomId := c.Params.ByName("id")

	room, found := Core.GetRoomById(roomId)
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

	websocket, err := gosockets.CreateClient(c.Writer, c.Request)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "WEBSOCKET_ERROR",
			"message": "errro while creating websocket client",
		})
		return
	}

	connectedClient := ClientsManager.CreateConnectedClient(websocket, room)
	connectedClient.WebSocket.OnDisconnect = append(connectedClient.WebSocket.OnDisconnect, func() {
		classicRoom.RemovePlayer(connectedClient.PlayerId)
		ClientsManager.RemoveConnectedClient(websocket.Id)

		println("Disconnected client with id: " + websocket.Id)
	})

	connectedClient.WebSocket.AddListener("request-join", func(c gosockets.MessageContext) {
		type requestType struct {
			Color string `json:"color"`
			Name  string `json:"name"`
		}
		data := requestType{}
		if err := c.BindJSON(&data); err != nil {
			return
		}

		player := classicRoom.AddPlayer(data.Name, data.Color)

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

	connectedClient.WebSocket.AddListener("request-leave", func(c gosockets.MessageContext) {
		classicRoom.RemovePlayer(connectedClient.PlayerId)
		connectedClient.IsPlayer = false
		connectedClient.PlayerId = ""
	})

	connectedClient.WebSocket.AddListener("change-direction", func(c gosockets.MessageContext) {
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
