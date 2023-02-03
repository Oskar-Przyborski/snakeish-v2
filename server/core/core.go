package core

import (
	"errors"

	"snakeish/core/room"
	classic_room "snakeish/core/room/classic"
	"snakeish/core/utils"

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

func (core *CoreInstance) CreateRoom(name string) (room.IRoom, error) {
	if _, foundDuplicate := core.GetRoomByName(name); foundDuplicate {
		return nil, errors.New("name already used by other room")
	}
	base := room.RoomBase{
		Name: name,
		Id:   uuid.NewString(),
	}

	newRoom := classic_room.ClassicRoom{
		RoomBase:   base,
		Apples:     []utils.Vector2{},
		MaxPlayers: 5,
	}

	core.rooms = append(core.rooms, &newRoom)

	return newRoom, nil
}
