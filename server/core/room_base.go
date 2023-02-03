package core

type RoomBase struct {
	Id       string
	Name     string
	ModeTag  string
	ModeName string
}

func (room RoomBase) GetName() string {
	return room.Name
}
func (room RoomBase) GetId() string {
	return room.Id
}
