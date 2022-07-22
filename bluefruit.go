//go:build circuitplay_bluefruit
// +build circuitplay_bluefruit

package gopherhelmet

import (
	"machine"

	"tinygo.org/x/drivers/servo"
)

var (
	servo1 servo.PWM = machine.PWM1
	servo2 servo.PWM = machine.PWM2
)

func InitBLE() {
	// TODO
}
