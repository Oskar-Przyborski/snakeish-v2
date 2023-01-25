package gosockets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	// "time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Decodes message from client. Message should look like 'event-name;payload'
func DecodeMessage(message string) (string, string, error) {
	var event, json string
	for idx, char := range strings.TrimSpace(message) {
		if char == ';' && idx > 0 {
			event = strings.TrimSpace(message[:idx])
			json = strings.TrimSpace(message[idx+1:])
			return event, json, nil
		}
	}
	return "", "", fmt.Errorf("couldn't parse the message")
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
func CreateClient(w http.ResponseWriter, r *http.Request) (*GosocketClient, error) {
	var id string = "ws-" + uuid.NewString()
	// if cookie, err := r.Cookie("id"); err == nil {
	// 	id = cookie.Value
	// } else {
	// 	id = "ws-" + uuid.NewString()
	// }

	headers := http.Header{}
	// headers.Add(
	// 	"Set-Cookie",
	// 	"id="+id+"; Expires="+time.Now().Add(time.Hour*24*360).Format(http.TimeFormat)+"; Path=/; HttpOnly; SameSite=None; Secure",
	// )

	ws, err := upgrader.Upgrade(w, r, headers)
	if err != nil {
		return &GosocketClient{}, fmt.Errorf("error while handshaking: %s", err.Error())
	}

	client := GosocketClient{
		WebSocket:    ws,
		Id:           id,
		Listeners:    map[string]func(string){},
		OnDisconnect: []func(){},
	}

	return &client, nil
}

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type GosocketClient struct {
	WebSocket    *websocket.Conn
	Id           string
	Listeners    map[string]func(string)
	OnDisconnect []func()
}

// Starts an infinite loop for listening to the events from client
func (client *GosocketClient) ListenForMessages() {
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
		event, json, err := DecodeMessage(message)
		if err != nil {
			continue
		}

		if callback := client.Listeners[event]; callback != nil {
			client.Listeners[event](json)
		}
	}

}

// Adds event listener that listen for events from clients
func (client *GosocketClient) AddListener(event string, callback func(string)) {
	client.Listeners[event] = callback
}

// Removes event listener
func (client *GosocketClient) RemoveListener(event string) {
	delete(client.Listeners, event)
}

// Sends event to the client
func (client *GosocketClient) Send(event string, payload interface{}) {
	client.WebSocket.WriteMessage(1, []byte(EncodeMessage(event, payload)))
}
