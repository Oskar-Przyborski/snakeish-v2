package core

import (
	"errors"

	"snakeish/core/room"
	classic_room "snakeish/core/room/classic"

	"github.com/google/uuid"
)

type CoreInstance struct {
	rooms []room.IRoom
}

func CreateCore() CoreInstance {
	return CoreInstance{
		rooms: []room.IRoom{},
	}
}

func (core CoreInstance) GetRooms() []room.IRoom {
	return core.rooms
}

func (core CoreInstance) GetRoomByName(name string) (room.IRoom, bool) {
	for _, room := range core.GetRooms() {
		if room.GetName() == name {
			return room, true
		}
	}
	return nil, false
}
func (core CoreInstance) GetRoomById(id string) (room.IRoom, bool) {
	for _, room := range core.GetRooms() {
		if room.GetId() == id {
			return room, true
		}
	}
	return nil, false
}

func (core *CoreInstance) CreateRoom(name string, configName string) (room.IRoom, error) {
	if _, foundDuplicate := core.GetRoomByName(name); foundDuplicate {
		return nil, errors.New("name already used by other room")
	}
	base := room.RoomBase{
		Name: name,
		Id:   uuid.NewString(),
	}

	room := CreateRoomByConfigName(configName, base)

	core.rooms = append(core.rooms, room)

	go room.StartRoom()

	return room, nil
}

func CreateRoomByConfigName(configName string, base room.RoomBase) room.IRoom {
	switch configName {
	default:
		return classic_room.CreateClassicRoom(base, "Casual")
	case "classic-huuge":
		return classic_room.CreateClassicRoom(base, "Huuge")
	}
}
