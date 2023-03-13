package core

import (
	"errors"
	"snakeish/pkg/core/rooms"
	"time"
)

var instance CoreInstance

type CoreInstance struct {
	rooms          []*rooms.Room
	roomsAfkTimers map[string]*time.Timer
}

func InitCore() {
	instance = CoreInstance{
		rooms:          []*rooms.Room{},
		roomsAfkTimers: make(map[string]*time.Timer),
	}
}

func GetRooms() []*rooms.Room {
	return instance.rooms
}

func GetRoomByName(name string) (*rooms.Room, bool) {
	for _, room := range GetRooms() {
		if room.Name == name {
			return room, true
		}
	}
	return nil, false
}

func GetRoomById(id string) (*rooms.Room, bool) {
	for _, room := range GetRooms() {
		if room.Id == id {
			return room, true
		}
	}
	return nil, false
}

func CreateRoom(name string, tagName string, modeName string, pin *[4]int) (*rooms.Room, error) {
	if _, foundDuplicate := GetRoomByName(name); foundDuplicate {
		return nil, errors.New("name already used by other room")
	}

	mode := rooms.CreateMode(tagName, modeName)
	room := rooms.Create(name, pin, mode)

	StartAfkForRoom(room.Id, 30*time.Second)
	instance.rooms = append(instance.rooms, room)
	go room.Start()

	return room, nil
}

func RemoveRoom(id string) {
	for idx, room := range GetRooms() {
		if room.Id == id {
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

func GetRoomPreview(r rooms.Room) rooms.RoomPreview {
	return rooms.RoomPreview{
		Id:         r.Id,
		Name:       r.GetModeName(),
		ModeTag:    r.GetTagName(),
		ModeName:   r.GetModeName(),
		Players:    r.GetPlayersCount(),
		MaxPlayers: r.GetMaxPlayers(),
		PinEnbled:  r.IsPinEnabled(),
	}
}
