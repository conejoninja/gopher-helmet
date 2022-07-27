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
)

func InitBLE(device *BackpackDevice, fn BLECallback) {
	// Do nothing, this board doesn't have BLE
}
