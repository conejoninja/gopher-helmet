package gopherhelmet

import (
	"image/color"
	"machine"
	"time"

	"github.com/conejoninja/gopher-helmet/fonts"
	"github.com/conejoninja/gopher-helmet/ledstripdisplay"
	"tinygo.org/x/tinyfont"

	"tinygo.org/x/drivers/ws2812"
)

// VisorDevice controls the Gopherbot Visor Neopixel LED.
type VisorDevice struct {
	*ledstripdisplay.Device
}

// Visor returns a new VisorDevice to control Gopherbot Visor.
func Visor() *VisorDevice {
	neo := machine.A2
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	v := ws2812.New(neo)

	display := ledstripdisplay.New(&v, VisorWidth, VisorHeight, ledstripdisplay.LAYOUT_Z)
	display.Configure(ledstripdisplay.Config{Rotation: ledstripdisplay.ROTATION_180})

	display.ClearDisplay()
	display.Display()

	return &VisorDevice{
		Device: &display,
	}
}

// Show sets the visor to display the current LED array state.
func (v *VisorDevice) Show() {
	v.Display()
}

// Off turns off all the LEDs.
func (v *VisorDevice) Off() {
	v.Clear()
}

// Clear clears the visor.
func (v *VisorDevice) Clear() {
	v.ClearDisplay()
}

func (v *VisorDevice) BootUp() {
	// 3 beeps
	for i := 0; i < 3; i++ {
		time.Sleep(800 * time.Millisecond)
		v.ClearDisplay()
		v.Display()
		time.Sleep(800 * time.Millisecond)
		v.SetPixel(8, 2, Red)
		v.Display()
	}
	time.Sleep(400 * time.Millisecond)

	// line
	for i := int16(0); i < 8; i++ {
		v.SetPixel(8-1-i, 2, Red)
		v.SetPixel(8+1+i, 2, Red)
		v.Display()
		time.Sleep(100 * time.Millisecond)
	}

	// grow vertically
	for i := int16(0); i < 17; i++ {
		v.SetPixel(i, 1, Red)
		v.SetPixel(i, 3, Red)
	}
	v.Display()
	time.Sleep(100 * time.Millisecond)

	for i := int16(0); i < 17; i++ {
		v.SetPixel(i, 0, Red)
		v.SetPixel(i, 4, Red)
	}
	v.Display()
	time.Sleep(100 * time.Millisecond)
	v.Clear()
}

func (v *VisorDevice) Marquee(text string, c color.RGBA) {
	w32, _ := tinyfont.LineWidth(&fonts.TomThumb, text)
	for i := int16(17); i > int16(-w32); i-- {
		v.Clear()
		tinyfont.WriteLine(v.Device, &fonts.TomThumb, i, 5, text, c)
		v.Display()
		time.Sleep(200 * time.Millisecond)
	}
}
