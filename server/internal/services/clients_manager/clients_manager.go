package clients_manager

import (
	"snakeish/pkg/core/room"
	"snakeish/pkg/sockets"
)

var instance ConnectedClientsManager

type ConnectedClientsManager struct {
	clients []*ConnectedClient
}

type ConnectedClient struct {
	WebSocket *sockets.GosocketClient
	RoomId    string
	IsPlayer  bool
	PlayerId  string
}

func Init() {
	instance = ConnectedClientsManager{
		clients: []*ConnectedClient{},
	}
}
func GetInstance() *ConnectedClientsManager {
	return &instance
}

func CreateConnectedClient(gosocket *sockets.GosocketClient, room room.IRoom) *ConnectedClient {
	client := &ConnectedClient{
		WebSocket: gosocket,
		RoomId:    room.GetId(),
		IsPlayer:  false,
		PlayerId:  "",
	}

	instance.clients = append(instance.clients, client)
	return client
}

func RemoveConnectedClient(websocketID string) {
	for idx, client := range instance.clients {
		if client.WebSocket.Id == websocketID {
			instance.clients = append(instance.clients[:idx], instance.clients[idx+1:]...)
			break
		}
	}
}

func GetClientsFromRoom(roomId string) []*ConnectedClient {
	foundClients := []*ConnectedClient{}
	for _, client := range instance.clients {
		if client.RoomId == roomId {
			foundClients = append(foundClients, client)
		}
	}

	return foundClients
}
