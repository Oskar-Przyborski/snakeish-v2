package clients

import (
	"snakeish/pkg/core/room"
	"snakeish/pkg/notifier"
	"snakeish/pkg/sockets"
)

type Client struct {
	WebSocket    *sockets.SocketClient
	Room         room.Room
	IsPlayer     bool
	PlayerId     string
	OnDisconnect notifier.Notifier[*Client]
}

func (client *Client) OnEvent(event string, callback func(*Client, sockets.MessageContext)) {
	client.WebSocket.AddListener(event, func(mc sockets.MessageContext) {
		callback(client, mc)
	})
}
