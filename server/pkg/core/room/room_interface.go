package room

type IRoom interface {
	GetName() string
	GetId() string
	GetModeTag() string
	GetModeName() string
	GetPlayersCount() int
	GetMaxPlayers() int
	GetPreview() RoomPreview
	OnUpdate(func())
	StartRoom()
	CheckPin([4]int) bool
}
