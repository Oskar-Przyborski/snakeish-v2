package sockets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Decodes message from client. Message should look like 'event-name;payload'
func DecodeMessage(message string, event *string, payload *string) error {
	for idx, char := range strings.TrimSpace(message) {
		if char == ';' && idx > 0 {
			*event = strings.TrimSpace(message[:idx])
			*payload = strings.TrimSpace(message[idx+1:])
			return nil
		}
	}
	return fmt.Errorf("couldn't parse the message")
}

// Encodes the event and payload so it can be send to the client. Output looks like 'event-name;payload'
func EncodeMessage(event string, payload interface{}) string {
	json, err := json.Marshal(payload)
	if err != nil {
		println("Error while converting to json:" + err.Error())
	}
	return event + ";" + string(json)
}

// Creates client from request (upgrades the request to WebSocket)
func CreateClient(w http.ResponseWriter, r *http.Request) (*SocketClient, error) {
	var id string = "ws-" + uuid.NewString()
	headers := http.Header{}

	ws, err := upgrader.Upgrade(w, r, headers)
	if err != nil {
		return &SocketClient{}, fmt.Errorf("error while handshaking: %s", err.Error())
	}

	client := SocketClient{
		WebSocket:    ws,
		Id:           id,
		Listeners:    map[string]func(MessageContext){},
		OnDisconnect: []func(){},
	}

	return &client, nil
}

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type SocketClient struct {
	WebSocket    *websocket.Conn
	Id           string
	Listeners    map[string]func(MessageContext)
	OnDisconnect []func()
}

// Starts an infinite loop for listening to the events from client
func (client *SocketClient) ListenForMessages() {
	for {

		msgType, bytes, err := client.WebSocket.ReadMessage()
		if err != nil {
			for _, f := range client.OnDisconnect {
				f()
			}
			break
		}
		if msgType != 1 {
			continue
		}
		message := string(bytes)
		var event, payload string

		if err := DecodeMessage(message, &event, &payload); err != nil {
			continue
		}

		if callback := client.Listeners[event]; callback != nil {
			client.Listeners[event](MessageContext{payload: payload})
		}
	}

}

// Adds event listener that listen for events from clients
func (client *SocketClient) AddListener(event string, callback func(MessageContext)) {
	client.Listeners[event] = callback
}

// Removes event listener
func (client *SocketClient) RemoveListener(event string) {
	delete(client.Listeners, event)
}

// Sends event to the client
func (client *SocketClient) Send(event string, payload interface{}) {
	client.WebSocket.WriteMessage(1, []byte(EncodeMessage(event, payload)))
}

type MessageContext struct {
	payload string
}

func (context MessageContext) BindJSON(object any) error {
	return json.Unmarshal([]byte(context.payload), object)
}
