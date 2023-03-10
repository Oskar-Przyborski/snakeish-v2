package core

import (
	"errors"
	"snakeish/pkg/core/room"
	"snakeish/pkg/core/room/battleroyale"
	"snakeish/pkg/core/room/classic"
	"time"

	"github.com/google/uuid"
)

var instance CoreInstance

type CoreInstance struct {
	rooms          []room.Room
	roomsAfkTimers map[string]*time.Timer
}

func InitCore() {
	instance = CoreInstance{
		rooms:          []room.Room{},
		roomsAfkTimers: make(map[string]*time.Timer),
	}
}

func GetRooms() []room.Room {
	return instance.rooms
}

func GetRoomByName(name string) (room.Room, bool) {
	for _, room := range GetRooms() {
		if room.GetName() == name {
			return room, true
		}
	}
	return nil, false
}
func GetRoomById(id string) (room.Room, bool) {
	for _, room := range GetRooms() {
		if room.GetId() == id {
			return room, true
		}
	}
	return nil, false
}

func CreateClassicRoom(roomName string, modeName string, pin *[4]int) (*classic.Room, error) {
	if _, foundDuplicate := GetRoomByName(roomName); foundDuplicate {
		return nil, errors.New("name already used by other room")
	}

	base := room.Base{
		PinRequirer: room.CreatePinRequirer(pin),
		Name:        roomName,
		Id:          uuid.NewString(),
		ModeName:    modeName,
		IsRunning:   true,
	}

	room := classic.Configure(base)
	StartAfkForRoom(room.GetId(), 30*time.Second)

	instance.rooms = append(instance.rooms, room)
	go room.StartRoom()

	return room, nil
}

func CreateBattleRoyaleRoom(roomName string, modeName string, pin *[4]int) (*battleroyale.Room, error) {
	if _, foundDuplicate := GetRoomByName(roomName); foundDuplicate {
		return nil, errors.New("name already used by other room")
	}
	base := room.Base{
		PinRequirer: room.CreatePinRequirer(pin),
		Name:        roomName,
		Id:          uuid.NewString(),
		ModeName:    modeName,
		IsRunning:   true,
	}

	room := battleroyale.Configure(base)
	StartAfkForRoom(room.GetId(), 30*time.Second)

	instance.rooms = append(instance.rooms, room)
	go room.StartRoom()

	return room, nil
}

func RemoveRoom(id string) {
	for idx, room := range GetRooms() {
		if room.GetId() == id {
			room.Stop()
			instance.rooms = append(instance.rooms[:idx], instance.rooms[idx+1:]...)
			return
		}
	}
}

func StartAfkForRoom(roomId string, duration time.Duration) {
	if timer := instance.roomsAfkTimers[roomId]; timer != nil {
		timer.Reset(duration)
	} else {
		instance.roomsAfkTimers[roomId] =
			time.AfterFunc(duration, func() {
				RemoveRoom(roomId)
			})
	}
}
func StopAfkForRoom(roomId string) {
	if timer := instance.roomsAfkTimers[roomId]; timer != nil {
		timer.Stop()
		delete(instance.roomsAfkTimers, roomId)
	}
}

func GetRoomPreview(r room.Room) room.RoomPreview {
	return room.RoomPreview{
		Id:         r.GetId(),
		Name:       r.GetName(),
		ModeTag:    r.GetModeTag(),
		ModeName:   r.GetModeName(),
		Players:    r.GetPlayersCount(),
		MaxPlayers: r.GetMaxPlayers(),
		PinEnbled:  r.IsPinEnabled(),
	}
}
