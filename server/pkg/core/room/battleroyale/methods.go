package battleroyale

import (
	"time"
)

func (room *Room) StartRoom() {
	for {
		if !room.IsRunning {
			break
		}

		time.Sleep(time.Duration(room.FrameTime) * time.Millisecond)

		switch room.GameStatus {
		case "waiting-for-players":
			room.TryStart()
		case "playing":
			room.Update()
		}

		room.OnUpdate.Notify(room)
	}
}

func (room *Room) TryStart() {
	if room.GameStatus != "waiting-for-players" {
		return
	}

	if len(room.Players) < room.MinPlayers {
		return
	}

	time.AfterFunc(5*time.Second, room.StartGame)
	room.StartUnix = time.Now().Add(5 * time.Second).UnixMilli()
	room.GameStatus = "starting"
}

func (room *Room) StartGame() {
	room.StartUnix = -1
	room.GameStatus = "playing"
	room.SpawnAllPlayers()
}

func (room *Room) Update() {
	playersToKill := []*Player{}

	alivePlayers := room.GetAlivePlayers()

	for _, player := range alivePlayers {
		room.MovePlayer(player)
	}

	for _, player := range alivePlayers {
		if player.IsOutOfBounds(room.GridSize) || player.IsCollidingSelf() || room.IsPlayerCollidingWithAnyOther(*player) {
			playersToKill = append(playersToKill, player)
		}
	}

	for _, player := range playersToKill {
		player.Kill()
	}

	room.SpawnMissingApples()
	room.CheckFinished()
}

func (room *Room) CheckFinished() {
	playersLeft := room.GetAlivePlayers()
	playersLeftCount := len(playersLeft)

	if playersLeftCount > 1 {
		return
	}

	if playersLeftCount == 1 {
		room.Finish(playersLeft[0])
	} else {
		room.Finish(nil)
	}
}

func (room *Room) Finish(winner *Player) {
	room.GameStatus = "finished"
	room.Winner = winner

	time.AfterFunc(time.Second*5, func() {
		room.Winner = nil
		room.GameStatus = "waiting-for-players"
	})
}
