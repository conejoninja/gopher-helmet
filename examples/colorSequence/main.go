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

type Msg struct {
	text string
	c    color.RGBA
}

var (
	visor *gopherhelmet.VisorDevice

	msg = gopherhelmet.TextColorSequence{
		{Text: "TINY GO", Color: gopherhelmet.Red},
		{Text: " ", Color: gopherhelmet.Black},
		{Text: "IS", Color: gopherhelmet.Magenta},
		{Text: " ", Color: gopherhelmet.Black},
		{Text: "AWESOME", Color: gopherhelmet.Green},
	}
)

func main() {

	time.Sleep(2 * time.Second)
	visor = gopherhelmet.Visor()

	visor.BootUp()

	visorLoop()
}

func visorLoop() {
	for {
		visor.MarqueeColored(msg)
		time.Sleep(50 * time.Millisecond)
	}
}
