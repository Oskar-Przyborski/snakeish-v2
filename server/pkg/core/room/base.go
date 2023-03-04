package room

type Base struct {
	Id        string
	Name      string
	ModeName  string
	IsRunning bool
	PinRequirer
}

func (room Base) GetName() string {
	return room.Name
}

func (room Base) GetModeName() string {
	return room.ModeName
}

func (room Base) GetId() string {
	return room.Id
}

func (room *Base) Stop() {
	room.IsRunning = false
}

type RoomPreview struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ModeTag    string `json:"modeTag"`
	ModeName   string `json:"modeName"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"maxPlayers"`
	PinEnbled  bool   `json:"pinEnabled"`
}

type PinRequirer struct {
	PinEnabled bool
	Pin        [4]int
}

func (manager PinRequirer) CheckPin(value [4]int) bool {
	if manager.PinEnabled {
		return manager.Pin == value
	} else {
		return true
	}
}

func (manager PinRequirer) IsPinEnabled() bool {
	return manager.PinEnabled
}

func CreatePinRequirer(pin *[4]int) PinRequirer {
	pinRequirer := PinRequirer{}
	if pin == nil {
		pinRequirer.PinEnabled = false
	} else {
		pinRequirer.PinEnabled = true
		pinRequirer.Pin = *pin
	}
	return pinRequirer
}
