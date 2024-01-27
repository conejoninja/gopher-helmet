//go:build circuitplay_express
// +build circuitplay_express

package gopherhelmet

import (
	"machine"

	"tinygo.org/x/drivers/servo"
)

var (
	servo1 servo.PWM = machine.TCC0
	servo2 servo.PWM = machine.TCC1

	earLeftPin  = machine.A1
	earRightPin = machine.A3
	antennaPin  = machine.A6

	neopixelsPin     = machine.NEOPIXELS
	btnAPin          = machine.BUTTONA
	btnBPin          = machine.BUTTONB
	sclPin           = machine.SCL1_PIN
	sdaPin           = machine.SDA1_PIN
	ledPin           = machine.LED
	tempPin          = machine.TEMPSENSOR
	lightPin         = machine.LIGHTSENSOR
	sliderPin        = machine.SLIDER
	visorPin         = machine.A2
	speakerEnablePin = machine.D11
	speakerPin       = machine.D12
)

func InitBLE(device *BackpackDevice, fn BLECallback) {
	// Do nothing, this board doesn't have BLE
}
