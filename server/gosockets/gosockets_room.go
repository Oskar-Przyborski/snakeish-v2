package gosockets

import (
	"fmt"
)

type GosocketRoom struct {
	Id      string
	Clients []*GosocketClient
	// The limit of clients that can be in a room at once. If 0, then it's unlimited
	MaxClients int
}

// Broadcasts event to all clients in the room
func (room *GosocketRoom) Broadcast(event string, payload interface{}) {
	for _, client := range room.Clients {
		client.Send(event, payload)
	}
}

// Broadcasts event to all clients in the room, except the passed sender
func (room *GosocketRoom) BroadcastFrom(sender GosocketClient, event string, payload interface{}) {
	for _, client := range room.Clients {
		if client.Id != sender.Id {
			client.Send(event, payload)
		}
	}
}

// Removes client from room
func (room *GosocketRoom) RemoveClientById(id string) {
	for idx, searchClient := range room.Clients {
		if searchClient.Id == id {
			println("Removing client from rooms")
			room.Clients = append(room.Clients[:idx], room.Clients[idx+1:]...)
			break
		}
	}
}

// Checks if room has client with passed id
func (room *GosocketRoom) HasClientWithId(id string) bool {
	for _, client := range room.Clients {
		if client.Id == id {
			return true
		}
	}
	return false
}

// Joins passed client to the room
func (room *GosocketRoom) JoinClient(client *GosocketClient) error {
	println("joining client to the wsroom id:", room.Id)
	if room.MaxClients != 0 && len(room.Clients) >= room.MaxClients {
		return fmt.Errorf("the room is full")
	}

	if room.HasClientWithId(client.Id) {
		return fmt.Errorf("the client is already in this room")
	}

	room.Clients = append(room.Clients, client)
	println("joined client to the wsroom id:", room.Id)
	return nil
}
