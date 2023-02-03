package core

import (
	"errors"

	"github.com/google/uuid"
)

type CoreInstance struct {
	rooms []IRoom
}

func CreateCore() CoreInstance {
	return CoreInstance{
		rooms: []IRoom{},
	}
}

func (core CoreInstance) GetRooms() []IRoom {
	return core.rooms
}

func (core CoreInstance) GetRoomByName(name string) (IRoom, bool) {
	for _, room := range core.GetRooms() {
		if room.GetName() == name {
			return room, true
		}
	}
	return RoomBase{}, false
}

func (core *CoreInstance) CreateRoom(name string) (IRoom, error) {
	if _, foundDuplicate := core.GetRoomByName(name); foundDuplicate {
		return RoomBase{}, errors.New("name already used by other room")
	}

	newRoom := RoomBase{
		Name: name,
		Id:   uuid.NewString(),
	}

	core.rooms = append(core.rooms, &newRoom)

	return newRoom, nil
}
