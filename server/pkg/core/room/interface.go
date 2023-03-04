package room

type Room interface {
	GetName() string
	GetId() string
	GetModeTag() string
	GetModeName() string
	GetPlayersCount() int
	GetMaxPlayers() int
	StartRoom()
	Stop()
	CheckPin([4]int) bool
	IsPinEnabled() bool
	RemovePlayer(string)
}
