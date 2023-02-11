package core

import (
	"errors"
	"snakeish/pkg/core/room"
	classic_room "snakeish/pkg/core/room/classic"
	"time"

	"github.com/google/uuid"
)

var instance CoreInstance

type CoreInstance struct {
	rooms          []room.IRoom
	roomsAfkTimers map[string]*time.Timer
}

func InitCore() {
	instance = CoreInstance{
		rooms:          []room.IRoom{},
		roomsAfkTimers: make(map[string]*time.Timer),
	}
}

func GetRooms() []room.IRoom {
	return instance.rooms
}

func GetRoomByName(name string) (room.IRoom, bool) {
	for _, room := range GetRooms() {
		if room.GetName() == name {
			return room, true
		}
	}
	return nil, false
}
func GetRoomById(id string) (room.IRoom, bool) {
	for _, room := range GetRooms() {
		if room.GetId() == id {
			return room, true
		}
	}
	return nil, false
}

func CreateClassicRoom(roomName string, modeName string, pin *[4]int) (*classic_room.ClassicRoom, error) {
	if _, foundDuplicate := GetRoomByName(roomName); foundDuplicate {
		return &classic_room.ClassicRoom{}, errors.New("name already used by other room")
	}

	base := room.RoomBase{
		PinRequirer: room.CreatePinRequirer(pin),
		Name:        roomName,
		Id:          uuid.NewString(),
	}

	room := classic_room.ConfigureClassicRoom(base, modeName)
	instance.rooms = append(instance.rooms, room)
	go room.StartRoom()

	return room, nil
}

func RemoveRoom(id string) {
	for idx, room := range GetRooms() {
		if room.GetId() == id {
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
