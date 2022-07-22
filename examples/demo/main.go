// Connects to an WS2812 RGB LED strip with 10 LEDS.
//
// See either the others.go or digispark.go files in this directory
// for the neopixels pin assignments.
package main

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"tinygo.org/x/drivers/scd4x"

	"machine"

	gopherhelmet "github.com/conejoninja/gopher-helmet"
)

type Msg struct {
	text string
	c    color.RGBA
}

const (
	UseCO2Sensor = true

	Bootup uint8 = iota
	Demo
	Eyes
	CO2
	Axis
	Message

	Idle uint8 = iota
	Swipe

	RGB = iota
	Fading
	Beeping
)

var (
	accel    *gopherhelmet.AccelerometerDevice
	visor    *gopherhelmet.VisorDevice
	speaker  *gopherhelmet.SpeakerDevice
	antenna  *gopherhelmet.AntennaDevice
	backpack *gopherhelmet.BackpackDevice
	ears     *gopherhelmet.EarsDevice

	left  *gopherhelmet.ButtonDevice
	right *gopherhelmet.ButtonDevice

	co2sensor *scd4x.Device

	visorMode    uint8 = Demo
	earsMode     uint8 = Idle
	antennaMode  uint8 = RGB
	antennaColor color.RGBA

	msgs = [5]Msg{
		{"HACK SESSION TINYGO FRIDAY 9:00", gopherhelmet.Red},
		{"ASK ME ABOUT TINYGO", gopherhelmet.Blue},
		{"@_CONEJO - TECHNOLOGIST FOR HIRE", gopherhelmet.Magenta},
		{"TALK - TINYGO: GETTING THE UPPER HEN BY DONIA CHAIEHLOUDJ", gopherhelmet.Yellow},
		{"FREE PINS AND STICKERS", gopherhelmet.Green},
	}
)

func main() {

	time.Sleep(2 * time.Second)
	accel = gopherhelmet.Accelerometer()
	visor = gopherhelmet.Visor()
	speaker = gopherhelmet.Speaker()

	antenna = gopherhelmet.Antenna()
	backpack = gopherhelmet.Backpack()
	ears = gopherhelmet.Ears()

	left = gopherhelmet.LeftButton()
	right = gopherhelmet.RightButton()

	if UseCO2Sensor {
		machine.I2C0.Configure(machine.I2CConfig{SCL: machine.SCL_PIN, SDA: machine.SDA_PIN})
		co2sensor = scd4x.New(machine.I2C0)
		co2sensor.Configure()

		if err := co2sensor.StartPeriodicMeasurement(); err != nil {
			println(err)
		}
	}

	gopherhelmet.InitBLE()

	//visor.BootUp()
	go earsLoop()
	go antennaLoop()

	visorLoop()
}

func visorLoop() {
	var step uint8 = 2
	for {
		switch visorMode {
		case Demo:
			switch step {
			case 0:
				visor.Marquee(msgs[0].text, msgs[0].c)
				break
			case 1:
				co2Marquee()
				break
			case 2:
				visor.Marquee(msgs[1].text, msgs[1].c)
				break
			case 3:
				lookingSides(gopherhelmet.White)
				break
			case 4:
				visor.Marquee(msgs[2].text, msgs[2].c)
				break
			case 5:
				break
			case 6:
				visor.Marquee(msgs[3].text, msgs[3].c)
				break
			case 7:
				demoAxis()
				break
			case 8:
				visor.Marquee(msgs[4].text, msgs[4].c)
				break
			case 9:
				break
			}
			step++

			if step > 9 {
				step = 0
			}
			break
		case Eyes:
			visorMode = Demo
			break
		case CO2:
			break
		case Axis:
			break
		case Message:
			visorMode = Demo
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func earsLoop() {
	angle := int16(90)
	forward := true
	for {
		switch earsMode {
		case Idle:
			if rand.Int31n(1000) == 1 {
				angle = 90
				forward = true
				//earsMode = Swipe
			}
			break
		case Swipe:
			if forward {
				println("FORWARD", angle)
				angle -= 10
				if angle <= 0 {
					angle = 0
					forward = false
				}
				ears.Set(2, angle)
			} else {
				angle += 10
				println("BACK", angle)
				if angle >= 90 {
					angle = 90
					forward = true
					earsMode = Idle
				}
				ears.Set(2, angle)
			}
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func antennaLoop() {
	var i uint8
	var f uint8 = 2
	for {
		switch antennaMode {
		case RGB:
			antenna.Rainbow(i)
			i++
			break
		case Fading:
			i += f
			if i >= 255 {
				i = 255
				f = -f
			}
			if i <= 0 {
				i = 0
				f = -f
			}
			antenna.SetColor(color.RGBA{i, 0, 0, 255})
			break
		case Beeping:
			i++
			if i >= 20 {
				i = 0
			}
			if i < 10 {
				antenna.SetColor(antennaColor)
			} else {
				antenna.SetColor(gopherhelmet.Black)
			}
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func demoAxis() {
	val := make([]int32, 3)
	c := gopherhelmet.Red
	for i := 0; i < 400; i++ {
		val[0], val[1], val[2], _ = accel.ReadAcceleration()
		visor.Clear()
		visor.SetPixel(8, 0, gopherhelmet.Red)
		visor.SetPixel(8, 2, gopherhelmet.Green)
		visor.SetPixel(8, 4, gopherhelmet.Blue)
		for k := int16(0); k < 3; k++ {
			val[k] = val[k] / 125000
			if k == 0 {
				c = gopherhelmet.Red
			} else if k == 1 {
				c = gopherhelmet.Green
			} else {
				c = gopherhelmet.Blue
			}
			if val[k] < 0 {
				for i := int16(8 + val[k]); i < 8; i++ {
					visor.SetPixel(i, k*2, c)
				}
			} else {
				for i := int16(8 + val[k]); i > 8; i-- {
					visor.SetPixel(i, k*2, c)
				}
			}
		}

		visor.Display()
		time.Sleep(50 * time.Millisecond)
	}
}

func co2Marquee() {
	var co2 int32
	var err error
	for i := 0; i < 5; i++ {
		co2, err = co2sensor.ReadCO2()
		value, _ := co2sensor.ReadTemperature()
		println("TEMP", value)
		value, _ = co2sensor.ReadHumidity()
		println("HUM", value)
		if err != nil {
			println(err)
		}
		println(co2)
		if co2 != 0 {
			break
		} else {
			time.Sleep(200 * time.Millisecond)
		}
	}
	switch {
	case co2 < 800:
		antennaColor = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
	case co2 < 1500:
		antennaColor = color.RGBA{R: 0xff, G: 0xff, B: 0x00}
	default:
		antennaColor = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
	}
	antennaMode = Beeping
	visor.Marquee("CO2 LEVELS: "+strconv.Itoa(int(co2))+" PPM", antennaColor)
	antennaMode = RGB
}
