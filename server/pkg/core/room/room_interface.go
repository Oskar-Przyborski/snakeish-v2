package room

type IRoom interface {
	GetName() string
	GetId() string
	GetModeTag() string
	GetModeName() string
	GetPlayersCount() int
	GetMaxPlayers() int
	GetPreview() RoomPreview
	StartRoom()
	CheckPin([4]int) bool
	RemovePlayer(string)
}
