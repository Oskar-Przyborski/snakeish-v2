package main

import (
	"encoding/json"
	"errors"
	"math/rand"
	"snakeish/golang/gosockets"
	"time"
)

type Room struct {
	Id             string
	RoomName       string
	FrameTime      int
	GridSize       int
	ApplesQuantity int
	Users          []*User
	Apples         []Vector2
}

func (room *Room) StartRoom() {
	room.SpawnMissingApples()
	for {
		time.Sleep(time.Duration(room.FrameTime) * time.Millisecond)
		room.Update()
	}
}

// Removes user from room
func (room *Room) RemoveUserById(id string) {
	for idx, user := range room.Users {
		if user.Websocket.Id == id {
			println("Removing user from rooms")
			room.Users = append(room.Users[:idx], room.Users[idx+1:]...)
			break
		}
	}
}

func (room *Room) AddUser(ws *gosockets.GosocketClient) {
	ws.OnDisconnect = append(ws.OnDisconnect, func() {
		room.RemoveUserById(ws.Id)
	})

	user := User{
		Websocket: ws,
		IsPlayer:  false,
	}

	type JoinPlayerPayload struct {
		Color string `json:"color"`
		Name  string `json:"name"`
	}

	ws.AddListener("join-player", func(s string) {
		payload := JoinPlayerPayload{}
		if err := json.Unmarshal([]byte(s), &payload); err != nil {
			println("Error while parsing json:", err.Error())
			return
		}

		user.IsPlayer = true

		user.player = Player{
			SnakeTail: []Vector2{},
			Color:     payload.Color,
			Name:      payload.Name,
			IsAlive:   false,
			TargetDirection: Vector2{
				X: 0,
				Y: 1,
			},
		}
	})

	ws.AddListener("leave-game", func(s string) {
		user.IsPlayer = false
		user.player = Player{}
	})

	type ChangeDirectionPayload struct {
		Direction string `json:"direction"`
	}

	ws.AddListener("change-direction", func(s string) {
		payload := ChangeDirectionPayload{}
		if err := json.Unmarshal([]byte(s), &payload); err != nil {
			println("Error while parsing json:", err.Error())
			return
		}

		dirVector, err := directionToVector(payload.Direction)
		if err != nil {
			return
		}

		if user.player.Direction.X+dirVector.X == 0 && user.player.Direction.Y+dirVector.Y == 0 {
			return
		}

		user.player.TargetDirection = dirVector
	})

	room.Users = append(room.Users, &user)
}

func (room *Room) getUsersPlayers() []*User {
	players := []*User{}
	for _, user := range room.Users {
		if user.IsPlayer {
			players = append(players, user)
		}
	}
	return players
}

func (room *Room) getPlayersCount() int {
	count := 0
	for _, user := range room.Users {
		if user.IsPlayer {
			count++
		}
	}
	return count
}

type GameUpdatePayload struct {
	Players  []Player  `json:"players"`
	Apples   []Vector2 `json:"apples"`
	GridSize int       `json:"gridSize"`
}

func (room *Room) BroadcastGameUpdate() {
	message := GameUpdatePayload{
		Players:  []Player{},
		Apples:   room.Apples,
		GridSize: room.GridSize,
	}

	for _, user := range room.getUsersPlayers() {
		message.Players = append(message.Players, user.player)
	}

	for _, user := range room.Users {
		user.Websocket.Send("game-update", message)
	}
}

func (room *Room) Update() {
	usersToKill := []*User{}

	for _, user := range room.getUsersPlayers() {
		if !user.player.IsAlive {
			continue
		}
		room.MovePlayer(&user.player)
	}

	for _, user := range room.getUsersPlayers() {
		if !user.player.IsAlive {
			continue
		}

		if user.player.IsOutOfBounds(room.GridSize) || user.player.IsCollidingSelf() || room.IsPlayerCollidingWithAnyOther(*user) {
			usersToKill = append(usersToKill, user)
			continue
		}
	}

	for _, user := range usersToKill {
		user.player.Kill()
	}

	for _, user := range room.getUsersPlayers() {
		if user.player.IsAlive {
			continue
		}
		room.RespawnPlayer(&user.player)
	}

	room.SpawnMissingApples()
	room.BroadcastGameUpdate()
}

func (room *Room) SpawnMissingApples() {
	applesToSpawn := room.ApplesQuantity - len(room.Apples)

	for i := 0; i < applesToSpawn; i++ {
		room.SpawnApple()
	}
}

func (room *Room) SpawnApple() {
	position, err := room.GetRandomFreePosition(0)
	if err != nil {
		return
	}

	room.Apples = append(room.Apples, position)
}

func (room *Room) IsPositionEmpty(at Vector2) bool {
	if room.IsAnySnakeAtPosition(at) || room.IsAppleAtPosition(at) {
		return false
	}
	return true
}

func (room *Room) IsAnySnakeAtPosition(at Vector2) bool {
	for _, user := range room.getUsersPlayers() {
		for _, snakeElement := range user.player.SnakeTail {
			if snakeElement.IsEqual(at) {
				return true
			}
		}
	}
	return false
}

func (room *Room) IsAppleAtPosition(at Vector2) bool {
	for _, apple := range room.Apples {
		if apple.IsEqual(at) {
			return true
		}
	}
	return false
}

func (room *Room) EatAppleAt(at Vector2) bool {
	for idx, apple := range room.Apples {
		if apple.IsEqual(at) {
			room.Apples = append(room.Apples[:idx], room.Apples[idx+1:]...)
			return true
		}
	}

	return false
}

func (room *Room) GetRandomFreePosition(distanceFromBounds int) (Vector2, error) {
	freePositions := []Vector2{}
	for y := distanceFromBounds; y < room.GridSize-distanceFromBounds; y++ {
		for x := distanceFromBounds; x < room.GridSize-distanceFromBounds; x++ {
			position := Vector2{Y: y, X: x}
			if room.IsPositionEmpty(position) {
				freePositions = append(freePositions, position)
			}
		}
	}
	if len(freePositions) == 0 {
		return Vector2{}, errors.New("there is no free position")
	}

	return freePositions[rand.Intn(len(freePositions))], nil
}

func (room *Room) MovePlayer(player *Player) {
	if len(player.SnakeTail) == 0 {
		return
	}

	player.Direction = player.TargetDirection

	newHeadPos := Vector2{
		X: player.SnakeTail[0].X + player.Direction.X,
		Y: player.SnakeTail[0].Y + player.Direction.Y,
	}

	if !room.EatAppleAt(newHeadPos) {
		player.SnakeTail = append([]Vector2{newHeadPos}, player.SnakeTail[:len(player.SnakeTail)-1]...)
	} else {
		player.SnakeTail = append([]Vector2{newHeadPos}, player.SnakeTail...)
	}
}

func (room *Room) RespawnPlayer(player *Player) {
	position, err := room.GetRandomFreePosition(1)
	if err != nil {
		return
	}
	player.SnakeTail = []Vector2{position}
	player.IsAlive = true
}

func (room *Room) IsPlayerCollidingWithAnyOther(user User) bool {
	for _, searchUser := range room.getUsersPlayers() {
		if searchUser.Websocket.Id == user.Websocket.Id {
			continue
		}
		if user.player.IsCollidingWith(searchUser.player) {
			return true
		}
	}

	return false
}
