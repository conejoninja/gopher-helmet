package gopherhelmet

import (
	"machine"

	"tinygo.org/x/drivers/servo"
)

const (
	ServoMin    = 600
	ServoCenter = 1500
	ServoMax    = 2400
)

type EarsDevice struct {
	Left  servo.Servo
	Right servo.Servo
}

func Ears() *EarsDevice {
	left, err := servo.New(servo1, machine.A1)
	if err != nil {
		println("could not configure servo")
	}
	right, err := servo.New(servo2, machine.A3)
	if err != nil {
		println("could not configure servo R")
	}
	return &EarsDevice{
		Left:  left,
		Right: right,
	}
}

func (b *EarsDevice) Center() {
	b.Set(2, 90)
}

func (b *EarsDevice) Back() {
	b.Set(2, 180)
}

func (b *EarsDevice) Front() {
	b.Set(2, 0)
}

func (b *EarsDevice) Off() {
	b.Left.SetMicroseconds(0)
	b.Right.SetMicroseconds(0)
}

func (b *EarsDevice) Set(ear uint8, angle int16) {
	x := 10 * int16(angle)
	if ear == 0 || ear == 2 {
		b.Left.SetMicroseconds(ServoMax - x)
	}
	if ear == 1 || ear == 2 {
		b.Right.SetMicroseconds(ServoMin + x)
	}
}
