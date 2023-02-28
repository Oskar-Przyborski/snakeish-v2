package battleroyale

import r "snakeish/pkg/core/room"

func Configure(base r.Base) *Room {
	switch base.ModeName {
	default:
		return &Room{
			Base: base,
			// ApplesQuantity: 3,
			// GridSize:       8,
			// MaxPlayers:     4,
			// FrameTime:      250,
		}
	}
}

func (room Room) GetMaxPlayers() int {
	return room.MaxPlayers
}

func (room Room) GetModeTag() string {
	return "battle royale"
}

func (room Room) GetPlayersCount() int {
	return len(room.Players)
}

func (room Room) GetPreview() r.RoomPreview {
	return r.RoomPreview{
		Id:         room.GetId(),
		Name:       room.GetName(),
		ModeTag:    room.GetModeTag(),
		ModeName:   room.GetModeName(),
		Players:    room.GetPlayersCount(),
		MaxPlayers: room.MaxPlayers,
		PinEnbled:  room.IsPinEnabled(),
	}
}

func (room Room) RemovePlayer(id string) {

}

func (room Room) StartRoom() {

}
