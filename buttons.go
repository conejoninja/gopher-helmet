package gopherhelmet

import (
	"machine"
)

// ButtonDevice is a button.
type ButtonDevice struct {
	machine.Pin
}

// Pushed checks to see if the button is being pushed.
func (b *ButtonDevice) Pushed() bool {
	return b.Get()
}

// LeftButton returns the left ButtonDevice.
func LeftButton() *ButtonDevice {
	left := btnAPin
	left.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	return &ButtonDevice{left}
}

// RightButton returns the right ButtonDriver.
func RightButton() *ButtonDevice {
	right := btnBPin
	right.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	return &ButtonDevice{right}
}
