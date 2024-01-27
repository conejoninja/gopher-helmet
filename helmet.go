package gopherhelmet

import (
	"image/color"
	"machine"
)

type BLECallback func(string)

const (
	// BackpackLEDCount is how many WS2812 LEDs are on the backpack, aka
	// the Circuit Playground Express.
	BackpackLEDCount = 10

	TopLEDCount = 10

	VisorWidth  = 17
	VisorHeight = 5
)

var (
	adcInitComplete = false
	i2cInitComplete = false

	Red     = color.RGBA{255, 0, 0, 255}
	Green   = color.RGBA{0, 255, 0, 255}
	Blue    = color.RGBA{0, 0, 255, 255}
	Yellow  = color.RGBA{255, 255, 0, 255}
	Magenta = color.RGBA{255, 0, 255, 255}
	Cyan    = color.RGBA{0, 255, 255, 255}
	White   = color.RGBA{255, 255, 255, 255}
	Black   = color.RGBA{0, 0, 0, 255}
)

// EnsureADCInit makes sure that the Gopherbot ADC has been initialized, but
// is only initialized once.
func EnsureADCInit() {
	if !adcInitComplete {
		machine.InitADC()
		adcInitComplete = true
	}
}

// EnsureI2CInit makes sure that the Gopherbot I2C has been initialized, but
// is only initialized once.
func EnsureI2CInit() {
	if !i2cInitComplete {
		machine.I2C1.Configure(machine.I2CConfig{SCL: sclPin, SDA: sdaPin})
		i2cInitComplete = true
	}
}
