package main

import (
	"snakeish/golang/gosockets"
)

type AppManager struct {
	Rooms            map[string]*Room
	WebsocketManager gosockets.GosocketManager
}

func CreateAppManager() AppManager {
	return AppManager{
		Rooms:            make(map[string]*Room),
		WebsocketManager: gosockets.GosocketManager{},
	}
}

func (gameManager *AppManager) GetRoomById(id string) *Room {
	return gameManager.Rooms[id]
}
