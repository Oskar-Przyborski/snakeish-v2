package main

import (
	"snakeish/core/room"
	"snakeish/gosockets"
)

type ConnectedClientsManager struct {
	clients []*ConnectedClient
}

type ConnectedClient struct {
	WebSocket *gosockets.GosocketClient
	RoomId    string
	IsPlayer  bool
	PlayerId  string
}

func CreateConnectedClientsManager() ConnectedClientsManager {
	return ConnectedClientsManager{
		clients: []*ConnectedClient{},
	}
}

func (manager *ConnectedClientsManager) CreateConnectedClient(gosocket *gosockets.GosocketClient, room room.IRoom) *ConnectedClient {
	client := &ConnectedClient{
		WebSocket: gosocket,
		RoomId:    room.GetId(),
		IsPlayer:  false,
		PlayerId:  "",
	}

	manager.clients = append(manager.clients, client)
	return client
}

func (manager *ConnectedClientsManager) RemoveConnectedClient(websocketID string) {
	for idx, client := range manager.clients {
		if client.WebSocket.Id == websocketID {
			manager.clients = append(manager.clients[:idx], manager.clients[idx+1:]...)
			break
		}
	}
}

func (manager *ConnectedClientsManager) GetClientsFromRoom(roomId string) []*ConnectedClient {
	foundClients := []*ConnectedClient{}
	for _, client := range manager.clients {
		if client.RoomId == roomId {
			foundClients = append(foundClients, client)
		}
	}

	return foundClients
}
