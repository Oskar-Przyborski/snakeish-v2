package core

import (
	"testing"
)

func TestCoreInitialization(t *testing.T) {
	InitCore()

	if len(GetRooms()) != 0 {
		t.Fatal("New core instance should have 0 rooms")
	}

	if len(instance.roomsAfkTimers) != 0 {
		t.Fatalf("New core instance should have 0 afk timers")
	}
}

func TestRoomCreation(t *testing.T) {
	InitCore()
	roomName := "Hello123"
	modeName := "Casual"

	room, err := CreateClassicRoom(roomName, modeName, nil)
	if err != nil {
		t.Fatalf("CoreInstance.CreateRoom() shouldn't return error")
	}
	if room.GetName() != roomName {
		t.Fatalf("IRoom.GetName() should return %s, but returned %s", roomName, room.GetName())
	}
	if room.PinRequirer.PinEnabled {
		t.Fatalf("Room's pin shouldn't be enabled")
	}

	rooms := GetRooms()

	if len(rooms) != 1 {
		t.Fatalf("CoreInstance.GetRooms() should return only 1 object, but returned %d objects", len(rooms))
	}

	if rooms[0].GetName() != roomName {
		t.Fatalf("IRoom.GetName() should return %s, but returned %s", roomName, rooms[0].GetName())
	}
}

func TestFindRoomByName(t *testing.T) {
	InitCore()
	room1, _ := CreateClassicRoom("1", "Casual", nil)
	room2, _ := CreateClassicRoom("2", "Casual", nil)

	roomFound1, found1 := GetRoomByName("1")
	roomFound2, found2 := GetRoomByName("2")
	_, found3 := GetRoomByName("3")

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
	InitCore()
	room1, _ := CreateClassicRoom("1", "Casual", nil)
	room2, _ := CreateClassicRoom("2", "Casual", nil)

	roomFound1, found1 := GetRoomById(room1.GetId())
	roomFound2, found2 := GetRoomById(room2.GetId())
	_, found3 := GetRoomById("jfhaksjdhfklhas")

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
	InitCore()

	_, err1 := CreateClassicRoom("room-name", "Casual", nil)
	if err1 != nil {
		t.Fatal("CoreInstance.CreateRoom() shouldn't return error")
	}

	_, err2 := CreateClassicRoom("room-name", "Casual", nil)
	if err2 == nil {
		t.Fatal("CoreInstance.CreateRoom() should return error")
	}

	_, err3 := CreateClassicRoom("room-name-other", "Casual", nil)
	if err3 != nil {
		t.Fatal("CoreInstance.CreateRoom() shouldn't return error")
	}

	rooms := GetRooms()
	if len(rooms) != 2 {
		t.Fatalf("CoreInstance.GetRooms() should return 2 objects, but returned %d objects", len(rooms))
	}
}
