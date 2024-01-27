package main

import (
	"image/color"
	"math/rand"
	"time"

	gopherhelmet "github.com/conejoninja/gopher-helmet"
)

const (
	Idle uint8 = iota
	Swipe

	RGB = iota
	Fading
	Beeping
)

var (
	antenna *gopherhelmet.AntennaDevice
	ears    *gopherhelmet.EarsDevice

	earsMode     uint8 = Idle
	antennaMode  uint8 = RGB
	antennaColor color.RGBA
)

func main() {

	antenna = gopherhelmet.Antenna()
	ears = gopherhelmet.Ears()
	ears.Set(2, 90)
	time.Sleep(1000 * time.Second)
	var i uint8
	angle := int16(90)
	forward := true
	for {
		switch earsMode {
		case Idle:
			ears.Off()
			if rand.Int31n(800) == 1 {
				angle = 90
				forward = true
				earsMode = Swipe
			}
			break
		case Swipe:
			if forward {
				angle -= 10
				if angle <= 0 {
					angle = 0
					forward = false
				}
				ears.Set(2, angle)
			} else {
				angle += 10
				if angle >= 90 {
					angle = 90
					forward = true
					earsMode = Idle
				}
				ears.Set(2, angle)
			}
			break
		}
		antenna.Rainbow(i)
		i++
		time.Sleep(50 * time.Millisecond)
	}

}
