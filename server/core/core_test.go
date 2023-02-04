package core

import (
	"testing"
)

func TestCoreInitialization(t *testing.T) {
	core := CreateCore()

	if len(core.GetRooms()) != 0 {
		t.Fatal("New core instance should have 0 rooms")
	}
}

func TestRoomCreation(t *testing.T) {
	core := CreateCore()
	roomName := "Hello123"
	modeName := "Casual"

	room, err := core.CreateClassicRoom(roomName, modeName)
	if err != nil {
		t.Fatalf("CoreInstance.CreateRoom() shouldn't return error")
	}
	if room.GetName() != roomName {
		t.Fatalf("IRoom.GetName() should return %s, but returned %s", roomName, room.GetName())
	}

	rooms := core.GetRooms()

	if len(rooms) != 1 {
		t.Fatalf("CoreInstance.GetRooms() should return only 1 object, but returned %d objects", len(rooms))
	}

	if rooms[0].GetName() != roomName {
		t.Fatalf("IRoom.GetName() should return %s, but returned %s", roomName, rooms[0].GetName())
	}
}

func TestFindRoomByName(t *testing.T) {
	core := CreateCore()
	room1, _ := core.CreateClassicRoom("1", "Casual")
	room2, _ := core.CreateClassicRoom("2", "Casual")

	roomFound1, found1 := core.GetRoomByName("1")
	roomFound2, found2 := core.GetRoomByName("2")
	_, found3 := core.GetRoomByName("3")

	if found1 == false {
		t.Fatalf("CoreInstance.GetRoomByName() should return true, but returned %t", found1)
	}
	if room1.GetId() != roomFound1.GetId() {
		t.Fatalf("CoreInstance.GetRoomByName() should return room named %s, but returned room named %s", room1.GetName(), roomFound1.GetName())
	}
	if found2 == false {
		t.Fatalf("CoreInstance.GetRoomByName() should return true, but returned %t", found2)
	}
	if room2.GetId() != roomFound2.GetId() {
		t.Fatalf("CoreInstance.GetRoomByName() should return room named %s, but returned room named %s", room2.GetName(), roomFound2.GetName())
	}

	if found3 == true {
		t.Fatalf("CoreInstance.GetRoomByName() should return false, but returned %t", found3)
	}
}
func TestFindRoomById(t *testing.T) {
	core := CreateCore()
	room1, _ := core.CreateClassicRoom("1", "Casual")
	room2, _ := core.CreateClassicRoom("2", "Casual")

	roomFound1, found1 := core.GetRoomById(room1.GetId())
	roomFound2, found2 := core.GetRoomById(room2.GetId())
	_, found3 := core.GetRoomById("jfhaksjdhfklhas")

	if found1 == false {
		t.Fatalf("CoreInstance.GetRoomById() should return true, but returned %t", found1)
	}
	if room1.GetId() != roomFound1.GetId() {
		t.Fatalf("CoreInstance.GetRoomById() should return room with id %s, but returned room with id %s", room1.GetId(), roomFound1.GetId())
	}
	if found2 == false {
		t.Fatalf("CoreInstance.GetRoomById() should return true, but returned %t", found2)
	}
	if room2.GetId() != roomFound2.GetId() {
		t.Fatalf("CoreInstance.GetRoomById() should return room with id %s, but returned room with id %s", room2.GetId(), roomFound2.GetId())
	}

	if found3 == true {
		t.Fatalf("CoreInstance.GetRoomById() should return false, but returned %t", found3)
	}
}

func TestRoomNamesDuplication(t *testing.T) {
	core := CreateCore()

	_, err1 := core.CreateClassicRoom("room-name", "Casual")
	if err1 != nil {
		t.Fatal("CoreInstance.CreateRoom() shouldn't return error")
	}

	_, err2 := core.CreateClassicRoom("room-name", "Casual")
	if err2 == nil {
		t.Fatal("CoreInstance.CreateRoom() should return error")
	}

	_, err3 := core.CreateClassicRoom("room-name-other", "Casual")
	if err3 != nil {
		t.Fatal("CoreInstance.CreateRoom() shouldn't return error")
	}

	rooms := core.GetRooms()
	if len(rooms) != 2 {
		t.Fatalf("CoreInstance.GetRooms() should return 2 objects, but returned %d objects", len(rooms))
	}
}
