package gopherhelmet

import (
	"image/color"
	"machine"
	"tinygo.org/x/drivers/ws2812"
)

// AntennaDevice controls the LED of the Gopherbot Helmet's antenna.
type AntennaDevice struct {
	ws2812.Device
	LED []color.RGBA
}

// Antenna returns a the Antenna to control the Gopherbot Antenna LED.
func Antenna() *AntennaDevice {
	led := machine.A6
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	v := ws2812.New(led)

	return &AntennaDevice{
		Device: v,
		LED:    make([]color.RGBA, TopLEDCount),
	}
}

func (b *AntennaDevice) Show() {
	b.WriteColors(b.LED)
}

// Off turns off all the LEDs.
func (b *AntennaDevice) Off() {
	b.Clear()
}

// Clear clears the Antenna LEDs.
func (b *AntennaDevice) Clear() {
	b.SetColor(color.RGBA{R: 0x00, G: 0x00, B: 0x00})
}

// SetColor sets the Antenna LEDs to a single color.
func (b *AntennaDevice) SetColor(color color.RGBA) {
	for i := uint8(0); i < TopLEDCount; i++ {
		b.LED[i] = color
	}

	b.Show()
}

// Red turns all of the Antenna LEDs red.
func (b *AntennaDevice) Red() {
	b.SetColor(color.RGBA{R: 0xff, G: 0x00, B: 0x00})
}

// Green turns all of the Antenna LEDs green.
func (b *AntennaDevice) Green() {
	b.SetColor(color.RGBA{R: 0x00, G: 0xff, B: 0x00})
}

// Blue turns all of the Antenna LEDs blue.
func (b *AntennaDevice) Blue() {
	b.SetColor(color.RGBA{R: 0x00, G: 0x00, B: 0xff})
}

func (b *AntennaDevice) Rainbow(n uint8) {
	for i := uint8(0); i < TopLEDCount; i++ {
		b.LED[i] = getRainbowRGB(n + i)
	}

	b.Show()
}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 85 {
		return color.RGBA{i * 3, 255 - i*3, 0, 255}
	} else if i < 170 {
		i -= 85
		return color.RGBA{255 - i*3, 0, i * 3, 255}
	}
	i -= 170
	return color.RGBA{0, i * 3, 255 - i*3, 255}
}
