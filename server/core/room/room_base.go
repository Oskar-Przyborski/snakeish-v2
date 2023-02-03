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
