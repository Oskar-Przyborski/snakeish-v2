package room

type RoomBase struct {
	Id   string
	Name string
}

func (room RoomBase) GetName() string {
	return room.Name
}
func (room RoomBase) GetId() string {
	return room.Id
}

type RoomPreview struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ModeTag    string `json:"modeTag"`
	ModeName   string `json:"modeName"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"maxPlayers"`
}
