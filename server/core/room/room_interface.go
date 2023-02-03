package room

type IRoom interface {
	GetName() string
	GetId() string
	GetModeTag() string
	GetModeName() string
}
