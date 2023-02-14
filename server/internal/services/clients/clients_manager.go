package clients

import (
	"snakeish/pkg/core/room"
	"snakeish/pkg/notifier"
	"snakeish/pkg/sockets"
)

var instance ClientsManager

type ClientsManager struct {
	clients []*Client
}

func Init() {
	instance = ClientsManager{
		clients: []*Client{},
	}
}

func GetInstance() *ClientsManager {
	return &instance
}

func CreateClient(gosocket *sockets.SocketClient, room room.IRoom) *Client {
	client := &Client{
		WebSocket:    gosocket,
		Room:         room,
		IsPlayer:     false,
		PlayerId:     "",
		OnDisconnect: notifier.Create[*Client](),
	}

	client.WebSocket.OnDisconnect = append(client.WebSocket.OnDisconnect, func() {
		client.OnDisconnect.Notify(client)
	})

	instance.clients = append(instance.clients, client)
	return client
}

func RemoveClient(websocketID string) {
	for idx, client := range instance.clients {
		if client.WebSocket.Id == websocketID {
			instance.clients = append(instance.clients[:idx], instance.clients[idx+1:]...)
			break
		}
	}
}

func GetClientsFromRoom(roomId string) []*Client {
	foundClients := []*Client{}
	for _, client := range instance.clients {
		if client.Room.GetId() == roomId {
			foundClients = append(foundClients, client)
		}
	}

	return foundClients
}
