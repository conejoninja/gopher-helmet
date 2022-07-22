// Connects to an WS2812 RGB LED strip with 10 LEDS.
//
// See either the others.go or digispark.go files in this directory
// for the neopixels pin assignments.
package main

import (
	"image/color"
	"time"

	gopherhelmet "github.com/conejoninja/gopher-helmet"
)

var (
	accel   *gopherhelmet.AccelerometerDevice
	visor   *gopherhelmet.VisorDevice
	speaker *gopherhelmet.SpeakerDevice

	antenna  *gopherhelmet.AntennaDevice
	backpack *gopherhelmet.BackpackDevice
	ears     *gopherhelmet.EarsDevice
)

func main() {

	time.Sleep(2 * time.Second)
	accel = gopherhelmet.Accelerometer()
	visor = gopherhelmet.Visor()
	speaker = gopherhelmet.Speaker()

	antenna = gopherhelmet.Antenna()
	backpack = gopherhelmet.Backpack()
	ears = gopherhelmet.Ears()

	//left := gopherhelmet.LeftButton()
	//right := gopherhelmet.RightButton()

	red := color.RGBA{R: 0xff}
	for {
		visor.Clear()
		eye(0, 0, red)
		eye(12, 0, red)
		visor.Display()
		time.Sleep(2000 * time.Millisecond)
		visor.Clear()
		eye(0, -1, red)
		eye(12, -1, red)
		visor.Display()
		time.Sleep(600 * time.Millisecond)
		visor.Clear()
		eye(0, 1, red)
		eye(12, 1, red)
		visor.Display()
		time.Sleep(600 * time.Millisecond)
		visor.Clear()
		eye(0, -1, red)
		eye(12, -1, red)
		visor.Display()
		time.Sleep(600 * time.Millisecond)
		visor.Clear()
		eye(0, 1, red)
		eye(12, 1, red)
		visor.Display()
		time.Sleep(600 * time.Millisecond)
		visor.Clear()
		eye(0, 0, red)
		eye(12, 0, red)
		visor.Display()
		time.Sleep(2000 * time.Millisecond)
		visor.Clear()
		eye_close(0, 0, red)
		eye_close(12, 0, red)
		visor.Display()
		time.Sleep(100 * time.Millisecond)
		visor.Clear()
		eye_fully_close(0, 0, red)
		eye_fully_close(12, 0, red)
		visor.Display()
		time.Sleep(100 * time.Millisecond)
		visor.Clear()
		eye_close(0, 0, red)
		eye_close(12, 0, red)
		visor.Display()
		time.Sleep(100 * time.Millisecond)
	}
}

func eye(x, z int16, c color.RGBA) {
	visor.SetPixel(x, 1, c)
	visor.SetPixel(x, 2, c)
	visor.SetPixel(x, 3, c)

	visor.SetPixel(x+1, 0, c)
	visor.SetPixel(x+2, 0, c)
	visor.SetPixel(x+3, 0, c)

	visor.SetPixel(x+4, 1, c)
	visor.SetPixel(x+4, 2, c)
	visor.SetPixel(x+4, 3, c)

	visor.SetPixel(x+1, 4, c)
	visor.SetPixel(x+2, 4, c)
	visor.SetPixel(x+3, 4, c)

	visor.SetPixel(x+2+z, 2, c)
}

func eye_close(x, z int16, c color.RGBA) {
	visor.SetPixel(x, 2, c)

	visor.SetPixel(x+1, 1, c)
	visor.SetPixel(x+2, 1, c)
	visor.SetPixel(x+3, 1, c)

	visor.SetPixel(x+4, 2, c)

	visor.SetPixel(x+1, 3, c)
	visor.SetPixel(x+2, 3, c)
	visor.SetPixel(x+3, 3, c)
}

func eye_fully_close(x, z int16, c color.RGBA) {
	visor.SetPixel(x, 2, c)
	visor.SetPixel(x+1, 2, c)
	visor.SetPixel(x+2, 2, c)
	visor.SetPixel(x+3, 2, c)
	visor.SetPixel(x+4, 2, c)
}
