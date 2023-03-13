package rooms

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
