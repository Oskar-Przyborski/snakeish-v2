package battleroyale

import "time"

func (room *Room) StartRoom() {
	for {
		time.Sleep(time.Duration(room.FrameTime) * time.Millisecond)
		room.Update()
	}
}

func (room *Room) Update() {
	room.OnUpdate.Notify(room)
}

func (room Room) RemovePlayer(id string) {

}
