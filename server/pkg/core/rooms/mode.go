package rooms

type Mode interface {
	Init()
	Update()
	RemovePlayer(string)

	GetFrameTime() int
	GetTagName() string
	GetModeName() string
	GetPlayersCount() int
	GetMaxPlayers() int
}
