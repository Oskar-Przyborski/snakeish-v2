package gosockets

import (
	"github.com/google/uuid"
)

type GosocketManager struct {
	Clients []*GosocketClient
	Rooms   []*GosocketRoom
}

// Adds client to the manager
func (manager *GosocketManager) AddClient(client *GosocketClient) {
	client.OnDisconnect = append(client.OnDisconnect, func() { manager.RemoveClientById(client.Id) })

	manager.Clients = append(manager.Clients, client)
}

// Broadcasts event to all connected clients
func (manager *GosocketManager) Broadcast(event string, payload interface{}) {
	for _, client := range manager.Clients {
		client.Send(event, payload)
	}
}

// Broadcasts event to all connected clients, except the passed sender
func (manager *GosocketManager) BroadcastFrom(sender GosocketClient, event string, payload interface{}) {
	for _, client := range manager.Clients {
		if client.Id != sender.Id {
			client.Send(event, payload)
		}
	}
}

func (manager *GosocketManager) CreateRoom(max int) *GosocketRoom {
	room := GosocketRoom{
		Id:         uuid.NewString(),
		Clients:    []*GosocketClient{},
		MaxClients: max,
	}
	manager.Rooms = append(manager.Rooms, &room)
	println("Created wsroom. Id:", room.Id)
	return &room
}

func (manager *GosocketManager) RemoveClientById(id string) {
	for idx, client := range manager.Clients {
		if client.Id == id {
			println("Removing client from clients")
			manager.Clients = append(manager.Clients[:idx], manager.Clients[idx+1:]...)
			break
		}
	}
	for _, room := range manager.Rooms {
		room.RemoveClientById(id)
	}
}
