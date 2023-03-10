package battleroyale

import (
	"snakeish/pkg/core/utils"
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
		if room.StartTimer != nil {
			room.StartTimer.Stop()
			room.StartTimer = nil
			room.StartUnix = -1
		}
		return
	}

	if room.StartTimer == nil {
		room.StartTimer = time.AfterFunc(5*time.Second, room.StartGame)
		room.StartUnix = time.Now().Add(5 * time.Second).UnixMilli()
	}
}

func (room *Room) StartGame() {
	room.Reset()
	room.StartUnix = -1
	room.StartTimer = nil
	room.GameStatus = "playing"
	room.HandleAppleSpawn()
	room.SpawnAllPlayers()

	room.Freezed = true
	room.UnfreezeUnix = time.Now().Add(3 * time.Second).UnixMilli()
	time.AfterFunc(3*time.Second, func() {
		room.UnfreezeUnix = -1
		room.Freezed = false
	})
}

func (room *Room) Update() {
	if room.Freezed {
		return
	}

	room.HandleZoneShrink()

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

	for _, player := range alivePlayers {
		for _, snakeElement := range player.SnakeTail {
			if room.IsInsideZone(snakeElement) {
				player.ZoneKillCounter++

				if player.ZoneKillCounter >= room.ZoneKillTime {
					player.SnakeTail = player.SnakeTail[:len(player.SnakeTail)-1]
					player.ZoneKillCounter = 0

					if len(player.SnakeTail) == 0 {
						playersToKill = append(playersToKill, player)
					}
				}
				break
			}
		}
	}

	for _, player := range playersToKill {
		player.Kill()
	}

	room.HandleAppleSpawn()
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

	time.AfterFunc(time.Second*5, room.Reset)
}

func (room *Room) Reset() {
	room.Winner = nil
	room.GameStatus = "waiting-for-players"
	room.Apples = []utils.Vector2{}
	room.Freezed = false
	room.UnfreezeUnix = -1
	room.ShrinkFrameCounter = 0
	room.ShrinkSize = 0

	for _, p := range room.Players {
		p.Direction = utils.Vector2{}
		p.SnakeTail = []utils.Vector2{}
		p.Kill()
	}
}

func (room *Room) HandleZoneShrink() {
	if room.ShrinkSize >= room.GridSize/2 {
		return
	}

	room.ShrinkFrameCounter++

	if room.ShrinkFrameCounter > room.ZoneShrinkTime {
		room.ShrinkSize++
		room.ShrinkFrameCounter = 0
	}
}
