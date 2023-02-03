package main

type AppManager struct {
	Rooms map[string]*Room
}

func CreateAppManager() AppManager {
	return AppManager{
		Rooms: make(map[string]*Room),
	}
}

func (gameManager *AppManager) GetRoomById(id string) *Room {
	return gameManager.Rooms[id]
}
