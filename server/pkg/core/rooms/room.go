package rooms

import (
	"snakeish/pkg/notifier"
	"time"

	"github.com/google/uuid"
)

type Room struct {
	Id   string
	Name string
	PinRequirer
	Running bool
	Mode
	OnUpdate notifier.Notifier[*Room]
}

type RoomPreview struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ModeTag    string `json:"modeTag"`
	ModeName   string `json:"modeName"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"maxPlayers"`
	PinEnbled  bool   `json:"pinEnabled"`
}

func (room *Room) Start() {
	room.Running = true
	room.Mode.Init()

	for room.Running {
		time.Sleep(time.Duration(room.GetFrameTime()) * time.Millisecond)

		room.Mode.Update()
		room.OnUpdate.Notify(room)
	}
}

func (room *Room) Stop() {
	room.Running = false
}

func Create(name string, pin *[4]int, mode Mode) *Room {
	return &Room{
		Id:          uuid.NewString(),
		Name:        name,
		PinRequirer: CreatePinRequirer(pin),
		Running:     false,
		Mode:        mode,
	}
}
