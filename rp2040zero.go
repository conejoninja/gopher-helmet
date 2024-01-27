//go:build rp2040_zero
// +build rp2040_zero

package gopherhelmet

import (
	"machine"

	"tinygo.org/x/drivers/servo"
)

var (
	servo1 servo.PWM = machine.PWM5
	servo2 servo.PWM = machine.PWM5

	earLeftPin  = machine.A0
	earRightPin = machine.A1
	antennaPin  = machine.D5

	neopixelsPin     = machine.NoPin
	btnAPin          = machine.NoPin
	btnBPin          = machine.NoPin
	sclPin           = machine.NoPin
	sdaPin           = machine.NoPin
	ledPin           = machine.NoPin
	tempPin          = machine.NoPin
	lightPin         = machine.NoPin
	sliderPin        = machine.NoPin
	visorPin         = machine.NoPin
	speakerEnablePin = machine.NoPin
	speakerPin       = machine.NoPin
)

func InitBLE(device *BackpackDevice, fn BLECallback) {
	// Do nothing, this board doesn't have BLE
}
