package battleroyale

func (room Room) GetMaxPlayers() int {
	return room.MaxPlayers
}

func (room Room) GetModeTag() string {
	return "battle royale"
}

func (room Room) GetPlayersCount() int {
	return len(room.Players)
}

func (room Room) RemovePlayer(id string) {

}

func (room Room) StartRoom() {

}
