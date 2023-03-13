package core

import (
	"testing"
	"time"
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
	modeTag := "classic"
	modeName := "Casual"

	room, err := CreateRoom(roomName, modeTag, modeName, nil)
	if err != nil {
		t.Fatalf("CoreInstance.CreateRoom() shouldn't return error")
	}
	if room.Name != roomName {
		t.Fatalf("IRoom.GetName() should return %s, but returned %s", roomName, room.Name)
	}
	if room.PinRequirer.PinEnabled {
		t.Fatalf("Room's pin shouldn't be enabled")
	}

	rooms := GetRooms()

	if len(rooms) != 1 {
		t.Fatalf("CoreInstance.GetRooms() should return only 1 object, but returned %d objects", len(rooms))
	}

	if rooms[0].Name != roomName {
		t.Fatalf("IRoom.GetName() should return %s, but returned %s", roomName, rooms[0].Name)
	}
}

func TestFindRoomByName(t *testing.T) {
	InitCore()
	room1, _ := CreateRoom("1", "classic", "Casual", nil)
	room2, _ := CreateRoom("2", "classic", "Casual", nil)

	roomFound1, found1 := GetRoomByName("1")
	roomFound2, found2 := GetRoomByName("2")
	_, found3 := GetRoomByName("3")

	if found1 == false {
		t.Fatalf("CoreInstance.GetRoomByName() should return true, but returned %t", found1)
	}
	if room1.Id != roomFound1.Id {
		t.Fatalf("CoreInstance.GetRoomByName() should return room named %s, but returned room named %s", room1.Name, roomFound1.Name)
	}
	if found2 == false {
		t.Fatalf("CoreInstance.GetRoomByName() should return true, but returned %t", found2)
	}
	if room2.Id != roomFound2.Id {
		t.Fatalf("CoreInstance.GetRoomByName() should return room named %s, but returned room named %s", room2.Name, roomFound2.Name)
	}

	if found3 == true {
		t.Fatalf("CoreInstance.GetRoomByName() should return false, but returned %t", found3)
	}
}

func TestFindRoomById(t *testing.T) {
	InitCore()
	room1, _ := CreateRoom("1", "classic", "Casual", nil)
	room2, _ := CreateRoom("2", "classic", "Casual", nil)

	roomFound1, found1 := GetRoomById(room1.Id)
	roomFound2, found2 := GetRoomById(room2.Id)
	_, found3 := GetRoomById("jfhaksjdhfklhas")

	if found1 == false {
		t.Fatalf("CoreInstance.GetRoomById() should return true, but returned %t", found1)
	}
	if room1.Id != roomFound1.Id {
		t.Fatalf("CoreInstance.GetRoomById() should return room with id %s, but returned room with id %s", room1.Id, roomFound1.Id)
	}
	if found2 == false {
		t.Fatalf("CoreInstance.GetRoomById() should return true, but returned %t", found2)
	}
	if room2.Id != roomFound2.Id {
		t.Fatalf("CoreInstance.GetRoomById() should return room with id %s, but returned room with id %s", room2.Id, roomFound2.Id)
	}

	if found3 == true {
		t.Fatalf("CoreInstance.GetRoomById() should return false, but returned %t", found3)
	}
}

func TestRoomNamesDuplication(t *testing.T) {
	InitCore()

	_, err1 := CreateRoom("room-name", "classic", "Casual", nil)
	if err1 != nil {
		t.Fatal("CoreInstance.CreateRoom() shouldn't return error")
	}

	_, err2 := CreateRoom("room-name", "classic", "Casual", nil)
	if err2 == nil {
		t.Fatal("CoreInstance.CreateRoom() should return error")
	}

	_, err3 := CreateRoom("room-name-other", "classic", "Casual", nil)
	if err3 != nil {
		t.Fatal("CoreInstance.CreateRoom() shouldn't return error")
	}

	rooms := GetRooms()
	if len(rooms) != 2 {
		t.Fatalf("CoreInstance.GetRooms() should return 2 objects, but returned %d objects", len(rooms))
	}
}

func TestRoomRemoving(t *testing.T) {
	InitCore()
	r1, _ := CreateRoom("r1", "classic", "Casual", nil)
	r2, _ := CreateRoom("r2", "classic", "Casual", nil)
	r3, _ := CreateRoom("r3", "classic", "Huuge", nil)

	RemoveRoom(r1.Id)
	if len(GetRooms()) != 2 {
		t.Fatalf("Rooms qty should equal 2 but equals: %v", len(GetRooms()))
	}
	for _, room := range GetRooms() {
		if room.Id == r1.Id {
			t.Fatalf("Found room1 after removing by it's Id")
		}
	}

	RemoveRoom(r2.Id)
	if len(GetRooms()) != 1 {
		t.Fatalf("Rooms qty should equal 1 but equals: %v", len(GetRooms()))
	}
	for _, room := range GetRooms() {
		if room.Id == r2.Id {
			t.Fatalf("Found room2 after removing by it's Id")
		}
	}

	RemoveRoom(r3.Id)
	if len(GetRooms()) != 0 {
		t.Fatalf("Rooms qty should equal 0 but equals: %v", len(GetRooms()))
	}
	for _, room := range GetRooms() {
		if room.Id == r3.Id {
			t.Fatalf("Found room3 after removing by it's Id")
		}
	}
}

func TestRoomsAfk(t *testing.T) {
	InitCore()

	room, _ := CreateRoom("room", "classic", "Casual", nil)
	StartAfkForRoom(room.Id, time.Second*3)

	if _, found := GetRoomById(room.Id); !found {
		t.Fatalf("Room should exist, but it doesn't")
	}

	time.Sleep(1 * time.Second)
	if _, found := GetRoomById(room.Id); !found {
		t.Fatalf("Room should exist, but it doesn't")
	}

	time.Sleep(1 * time.Second)
	if _, found := GetRoomById(room.Id); !found {
		t.Fatalf("Room should exist, but it doesn't")
	}

	time.Sleep(2 * time.Second)
	if _, found := GetRoomById(room.Id); found {
		t.Fatalf("Room shouldn't exist, but it does")
	}
}

func TestAfkStop(t *testing.T) {
	InitCore()

	room, _ := CreateRoom("room", "classic", "Casual", nil)
	StartAfkForRoom(room.Id, time.Second*3)

	if _, found := GetRoomById(room.Id); !found {
		t.Fatalf("Room should exist, but it doesn't")
	}

	time.Sleep(1 * time.Second)
	if _, found := GetRoomById(room.Id); !found {
		t.Fatalf("Room should exist, but it doesn't")
	}

	time.Sleep(1 * time.Second)
	if _, found := GetRoomById(room.Id); !found {
		t.Fatalf("Room should exist, but it doesn't")
	}

	StopAfkForRoom(room.Id)
	time.Sleep(2 * time.Second)
	if _, found := GetRoomById(room.Id); !found {
		t.Fatalf("Room should exist, but it doesn't")
	}
}
